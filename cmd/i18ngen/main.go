package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/lehaisonmath6/go-i18n"
)

type DataTemplate struct {
	ListLanguage  []string
	ListMessageID []string
	PackageName   string
	ListVarName   []string
	ListFileLang  []string
}

func CreateGenMessageIDFile(outFile string, dataTemp DataTemplate) {
	funMap := template.FuncMap{
		"toUpper": strings.ToUpper,
	}
	f, err := os.Create(outFile + ".go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	templ, err := template.New(outFile + ".go").Funcs(funMap).Parse(`package {{.PackageName}}

import "github.com/lehaisonmath6/go-i18n"

const ({{range .ListLanguage}}
	{{.| toUpper}} = i18n.Lang("{{.}}")
{{end}})
const ({{range $index,$element := .ListMessageID}}
	{{$element | toUpper}} = i18n.MessageID("{{$element}}")
{{end}})`)
	if err != nil {
		slog.Error("create and parse template", err)
		return
	}

	err = templ.Execute(f, dataTemp)
	if err != nil {
		slog.Error("template excute", err)
		return
	}
}

func CreateTestFile(outFile string, dataTemp DataTemplate) {
	funMap := template.FuncMap{
		"toUpper": strings.ToUpper,
		"join": func(elems []string, sep string) string {
			newElems := []string{}
			for _, e := range elems {
				newElems = append(newElems, "\""+e+"\"")
			}
			return strings.Join(newElems, sep)
		},
		"genMap": func(elems []string) string {
			newElems := []string{}
			for id, e := range elems {
				newElems = append(newElems, "\""+e+"\""+":"+"\""+fmt.Sprint("varname", id, "abc"), "\"")
			}
			return strings.Join(newElems, ",")
		},
	}
	f, err := os.Create(outFile + "_test.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	templ, err := template.New(outFile + "_test.go").Funcs(funMap).Parse(`package {{.PackageName}}

import (
	"testing"

	"github.com/lehaisonmath6/go-i18n"
)
{{$languages := .ListLanguage}}
{{$messageIDs := .ListMessageID}}
func TestCreateMessageIDGen(t *testing.T) {
	var message string
	var err error
	loc, err := i18n.NewLocalizer([]string{ {{join  .ListFileLang ","}} })
	if err != nil {
		t.Error("New localizer err", err)
		return
	}
	templData := map[string]string{
		{{range .ListVarName}} "{{.}}" : "abc" , {{end}}
	}{{range $lang := $languages}}{{range $messageID := $messageIDs}}
	message, err = loc.GetMessageLang({{$lang | toUpper}}, {{$messageID | toUpper}}, templData, i18n.Other)
	if err != nil || message == "" {
		t.Error("GenMessageLang in",  {{$lang | toUpper}}, "msg id", {{$messageID | toUpper}}, err)
	} else {
		t.Log("lang",{{$lang | toUpper}}, "msgID",{{$messageID | toUpper}}, ":", message)
	}
	{{end}}
{{end}}
}
	`)
	if err != nil {
		slog.Error("create and parse template", err)
		return
	}

	err = templ.Execute(f, dataTemp)
	if err != nil {
		slog.Error("template excute", err)
		return
	}

}

func main() {
	dirLang := flag.String("dir", "", "directory of list language i18n")
	outFile := flag.String("out", "message_id_gen", "generate file name")
	genTest := flag.Bool("genTest", true, "true is gen test file, false is disable gen test file")
	flag.Parse()
	if dirLang == nil || strings.TrimSpace(*dirLang) == "" {
		slog.Error("dir lang must not empty. Ex run command $autogenidmessage -dir ./lang")
		return
	}
	fmt.Println("Dir lang is", *dirLang)
	lsFiles, err := os.ReadDir(*dirLang)
	if err != nil {
		slog.Error("read dir", *dirLang, err)
		return
	}
	if len(lsFiles) == 0 {
		slog.Error("not have any lang file")
		return
	}
	listLang := []string{}
	firstLangFile := ""
	packageName := "models"
	currentDir, _ := os.Getwd()
	if currentDir != "" {
		log.Println("packageNAme", filepath.Base(currentDir))
		packageName = filepath.Base(currentDir)
	}
	listFileLang := []string{}
	for _, file := range lsFiles {
		if file.IsDir() {
			continue
		}
		fileName := file.Name()
		absFileName, _ := filepath.Abs(*dirLang + "/" + fileName)
		fileNameChunk := strings.Split(fileName, ".")
		fileNameBase := fileName
		if len(fileNameChunk) > 1 {
			fileNameBase = fileNameChunk[0]
		}
		listLang = append(listLang, (fileNameBase))
		listFileLang = append(listFileLang, absFileName)
		firstLangFile = absFileName
	}

	messFile, err := i18n.LoadMessageFile(firstLangFile)
	if err != nil {
		slog.Error("parse message file", err, "file path", firstLangFile)
		return
	}
	var listMessageID = []string{}
	for _, mess := range messFile.Messages {
		listMessageID = append(listMessageID, mess.ID)
	}
	dataTemplate := DataTemplate{
		ListLanguage:  listLang,
		ListMessageID: listMessageID,
		PackageName:   packageName,
		ListVarName:   messFile.ListVarName,
		ListFileLang:  listFileLang,
	}
	CreateGenMessageIDFile(*outFile, dataTemplate)
	if *genTest {
		CreateTestFile(*outFile, dataTemplate)
	}
	// fmt.Println("list Language", listLang)
	// fmt.Println("list MessageID", listMessageID)

}
