package main

import (
	"log"

	"github.com/lehaisonmath6/go-i18n"
)

func main() {
	loc, err := i18n.NewLocalizer([]string{"/home/lehaisonmath6/go/src/github.com/lehaisonmath6/go-i18n/test/lang/vi.json"})
	if err != nil {
		log.Fatalln("Read err", err)
	}
	log.Println("loc data", loc)
	mess, err := loc.GetMessageLang("vi", "user_comment_in_your_post", map[string]string{
		"Name":  "Sơn",
		"Name2": "Long",
		"num":   "5",
	}, 7)
	log.Println("msg", mess, "err", err)
}
