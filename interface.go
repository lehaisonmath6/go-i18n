package i18n

type MessageID string

type Localizer interface {
	GetMessageLang(lang Lang, idmessage MessageID, templateData map[string]string, selectType SelectType) (string, error)
}
