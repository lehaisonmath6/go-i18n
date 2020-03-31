package i18n

type Localizer interface {
	GetMessageLang(lang string, idmessage string, templateData map[string]string, pluralcount int) (string, error)
}
