package i18n

import (
	"encoding/json"
	"io/ioutil"
)

type MessageFile struct {
	Path     string
	Lang     string
	Format   string
	Messages []*Message
}

// Eacth path data for each language
func NewLocalizer(listPath []string) (Localizer, error) {
	loc := &localizer{
		mapDataLang: make(map[string]map[string]map[int]string),
	}
	for _, path := range listPath {
		messFile, err := LoadMessageFile(path)
		if err != nil {
			return nil, err
		}
		loc.mapDataLang[messFile.Lang] = make(map[string]map[int]string)
		for _, mess := range messFile.Messages {
			loc.mapDataLang[messFile.Lang][mess.ID] = make(map[int]string)
			loc.mapDataLang[messFile.Lang][mess.ID][0] = mess.Other
			loc.mapDataLang[messFile.Lang][mess.ID][1] = mess.One
			loc.mapDataLang[messFile.Lang][mess.ID][2] = mess.Two
			loc.mapDataLang[messFile.Lang][mess.ID][3] = mess.Many
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

	return messageFile, nil
}
func LoadMessageFile(path string) (*MessageFile, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseMessageFileBytes(buf, path)
}
