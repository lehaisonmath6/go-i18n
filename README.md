# This is lib and tool for autogenerate i18n


- Build binary tool in [i18gen](./cmd/i18ngen/)
- To test gen data file use in [test](./test/)

# Structure of lang folder 
- All language file place in lang folder 
- Language file must be json file with format [lang].json
- In each lang file have multi messsage id with lang 
```
{
    "$message_id_1" : {
        "description" : "descrition of message id",
        "other" : "default message id content",
        "one" :"this is content for singular object ex: {{.Name}} comment on your post",
        "two" :"this is for double object, ex : {.Name}} and {{.Name2}} comment to your post",
        "many" : "{{.Name}}, {{.Name2}} and {{.total}} other comment on your post",
    },
     "$message_id_2" : {
        "description" : "descrition of message id",
        "other" : "default message id content",
        "one" :"this is content for singular object ex: {{.Name}} comment on your post",
        "two" :"this is for double object, ex : {.Name}} and {{.Name2}} comment to your post",
        "many" : "{{.Name}}, {{.Name2}} and {{.total}} other comment on your post",
    }
    .....
}
```
- You can omit fields  "one","two" and "many" but field "other" must required


# About tool gen i18n
- Generate and validate all MesssageID and all language in lang folder
- Generate test file to check valid message result

# How it works ?

1. Build i18ngen tool in [i18gen](./cmd/i18ngen/) : ```$ go build . ```
2. Create folder lang and put all language file in this lang folder (ex vi.json,en.json,zh.json)
3. Run tool gen: ``` $i18ngen -dir ./lang ```
4. Tool'll generate 2 file message_id_gen.go and messasage_gen_test.go
5. Example main code 
```
package main

import (
	"log/slog"

	"github.com/lehaisonmath6/go-i18n"
)

func main() {
	var message string
	var err error
	loc, err := i18n.NewLocalizer([]string{"/home/lehaisonmath6/go/src/github.com/lehaisonmath6/go-i18n/test/lang/en.json", "/home/lehaisonmath6/go/src/github.com/lehaisonmath6/go-i18n/test/lang/vi.json"})
	if err != nil {
		slog.Error("New localizer err", err)
		return
	}
	templData := map[string]string{
		".Name": "abc", ".Name2": "abc", ".num": "abc",
	}
	message, err = loc.GetMessageLang(EN, USER_BEGIN_LIVESTREAM, templData, i18n.Other)
	if err != nil || message == "" {
		slog.Error("GenMessageLang in", EN, "msg id", USER_BEGIN_LIVESTREAM, err)
	} else {
		slog.Info("lang", EN, "msgID", USER_BEGIN_LIVESTREAM, ":", message)
	}
}
```