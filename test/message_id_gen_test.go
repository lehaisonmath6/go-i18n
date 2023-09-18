package main

import (
	"testing"

	"github.com/lehaisonmath6/go-i18n"
)

func TestCreateMessageIDGen(t *testing.T) {
	var message string
	var err error
	loc, err := i18n.NewLocalizer([]string{"/home/lehaisonmath6/go/src/github.com/lehaisonmath6/go-i18n/test/lang/en.json", "/home/lehaisonmath6/go/src/github.com/lehaisonmath6/go-i18n/test/lang/vi.json"})
	if err != nil {
		t.Error("New localizer err", err)
		return
	}
	templData := map[string]string{
		".Name": "abc", ".Name2": "abc", ".num": "abc",
	}
	message, err = loc.GetMessageLang(EN, USER_BEGIN_LIVESTREAM, templData, i18n.Other)
	if err != nil || message == "" {
		t.Error("GenMessageLang in", EN, "msg id", USER_BEGIN_LIVESTREAM, err)
	} else {
		t.Log("lang", EN, "msgID", USER_BEGIN_LIVESTREAM, ":", message)
	}

	message, err = loc.GetMessageLang(EN, USER_COMMENT_IN_YOUR_POST, templData, i18n.Other)
	if err != nil || message == "" {
		t.Error("GenMessageLang in", EN, "msg id", USER_COMMENT_IN_YOUR_POST, err)
	} else {
		t.Log("lang", EN, "msgID", USER_COMMENT_IN_YOUR_POST, ":", message)
	}

	message, err = loc.GetMessageLang(VI, USER_BEGIN_LIVESTREAM, templData, i18n.Other)
	if err != nil || message == "" {
		t.Error("GenMessageLang in", VI, "msg id", USER_BEGIN_LIVESTREAM, err)
	} else {
		t.Log("lang", VI, "msgID", USER_BEGIN_LIVESTREAM, ":", message)
	}

	message, err = loc.GetMessageLang(VI, USER_COMMENT_IN_YOUR_POST, templData, i18n.Other)
	if err != nil || message == "" {
		t.Error("GenMessageLang in", VI, "msg id", USER_COMMENT_IN_YOUR_POST, err)
	} else {
		t.Log("lang", VI, "msgID", USER_COMMENT_IN_YOUR_POST, ":", message)
	}

}
