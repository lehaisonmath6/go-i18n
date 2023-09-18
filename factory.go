package i18n

import (
	"encoding/json"
	"errors"
	"os"
	"regexp"
)

type MessageFile struct {
	Path     string
	Lang     string
	Format   string
	Messages []*Message

	ListVarName []string
}

var (
	reVarName = regexp.MustCompile(`\{\{([^}]+)\}\}`)
)

func getAllVarName(mess MessageFile) []string {
	setVar := make(map[string]bool)
	for _, mess := range mess.Messages {
		matches := reVarName.FindAllStringSubmatch(mess.One, -1)
		if len(matches) > 0 {
			for _, m := range matches {
				setVar[m[1]] = true
			}
		}
		matches = reVarName.FindAllStringSubmatch(mess.Two, -1)
		if len(matches) > 0 {
			for _, m := range matches {
				setVar[m[1]] = true
			}
		}
		matches = reVarName.FindAllStringSubmatch(mess.Many, -1)
		if len(matches) > 0 {
			for _, m := range matches {
				setVar[m[1]] = true
			}
		}

		matches = reVarName.FindAllStringSubmatch(mess.Other, -1)
		if len(matches) > 0 {
			for _, m := range matches {
				setVar[m[1]] = true
			}
		}
	}
	lsVar := []string{}
	for k, _ := range setVar {
		lsVar = append(lsVar, k)
	}
	return lsVar
}

// Eacth path data for each language
func NewLocalizer(listPath []string) (Localizer, error) {
	loc := &localizer{
		mapDataLang: make(map[Lang]map[MessageID]map[int]string),
	}
	totalMessageID := 0
	for index, path := range listPath {
		messFile, err := LoadMessageFile(path)
		if err != nil {
			return nil, err
		}
		if index == 0 {
			totalMessageID = len(messFile.Messages)
		} else if totalMessageID != len(messFile.Messages) {
			return nil, errors.New("all language file must same total message id")
		}
		loc.mapDataLang[Lang(messFile.Lang)] = make(map[MessageID]map[int]string)
		for _, mess := range messFile.Messages {
			if mess.Other == "" {
				return nil, errors.New("Message file " + path + " must have other field")
			}
			loc.mapDataLang[Lang(messFile.Lang)][MessageID(mess.ID)] = make(map[int]string)
			loc.mapDataLang[Lang(messFile.Lang)][MessageID(mess.ID)][0] = mess.Other
			loc.mapDataLang[Lang(messFile.Lang)][MessageID(mess.ID)][1] = mess.One
			loc.mapDataLang[Lang(messFile.Lang)][MessageID(mess.ID)][2] = mess.Two
			loc.mapDataLang[Lang(messFile.Lang)][MessageID(mess.ID)][3] = mess.Many

		}
	}
	return loc, nil
}

func ParseMessageFileBytes(buf []byte, path string) (*MessageFile, error) {
	lang, format := parsePath(path)
	// tag := language.Make(lang)
	messageFile := &MessageFile{
		Path:   path,
		Lang:   lang,
		Format: format,
	}
	if len(buf) == 0 {
		return messageFile, nil
	}
	// unmarshalFunc := unmarshalFuncs[messageFile.Format]
	// if unmarshalFunc == nil {
	// 	if messageFile.Format == "json" {
	// 		unmarshalFunc = json.Unmarshal
	// 	} else {
	// 		return nil, fmt.Errorf("no unmarshaler registered for %s", messageFile.Format)
	// 	}
	// }
	var err error
	var raw interface{}
	if err = json.Unmarshal(buf, &raw); err != nil {
		return nil, err
	}

	if messageFile.Messages, err = recGetMessages(raw, isMessage(raw), true); err != nil {
		return nil, err
	}
	messageFile.ListVarName = getAllVarName(*messageFile)
	return messageFile, nil
}
func LoadMessageFile(path string) (*MessageFile, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseMessageFileBytes(buf, path)
}
