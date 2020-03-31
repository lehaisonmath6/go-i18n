package i18n

import (
	"errors"
	"strings"
)

type localizer struct {
	mapDataLang map[string]map[string]map[int]string
}

func (m *localizer) GetMessageLang(lang string, idmessage string, templateData map[string]string, pluralcount int) (string, error) {
	if m.mapDataLang == nil || m.mapDataLang[lang] == nil || m.mapDataLang[lang][idmessage] == nil {
		return "", errors.New("Map data null")
	}
	var mess string
	var ok bool
	if pluralcount <= 2 {
		switch pluralcount {
		case 0, 1, 2:
			mess, ok = m.mapDataLang[lang][idmessage][pluralcount]
		default:
			mess, ok = m.mapDataLang[lang][idmessage][0]
		}
	} else {
		mess, ok = m.mapDataLang[lang][idmessage][3]
	}

	if !ok {
		mess, ok = m.mapDataLang[lang][idmessage][pluralcount]
	}
	if !ok {
		return "", errors.New("Not found message")
	}
	for k, v := range templateData {
		mess = strings.Replace(mess, "{{."+k+"}}", v, -1)
	}
	return mess, nil

}
