package i18n

import (
	"golang.org/x/text/language"
)

var defaultLanguage = language.Chinese
var bundle = NewBundle(defaultLanguage)

func GetMessage(lang any, code string) string {
	return bundle.GetTranslator(lang).Translate(code, map[string]string{})
}

func GetMessageWithArgs(lang any, code string, args map[string]string) string {
	return bundle.GetTranslator(lang).Translate(code, args)
}

func GetDefaultLanguage() language.Tag {
	return defaultLanguage
}

func SetDefaultLanguage(language language.Tag) {
	defaultLanguage = language
	bundle.SetDefaultLanguage(language)
}
