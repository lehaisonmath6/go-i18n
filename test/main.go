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
