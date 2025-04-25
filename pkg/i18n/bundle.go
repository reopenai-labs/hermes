package i18n

import (
	"context"
	"golang.org/x/text/language"
	"qiyibyte.com/hermes/pkg/builtin/collection"
)

type Bundle struct {
	defaultLanguage   string
	defaultTranslator *Translator
	translators       *collection.CowMap[string, *Translator]
}

func NewBundle(defaultLanguage language.Tag) *Bundle {
	return &Bundle{
		defaultTranslator: NewTranslator(),
		defaultLanguage:   defaultLanguage.String(),
		translators:       collection.NewCowMap[string, *Translator](),
	}
}

func (Self *Bundle) SetDefaultLanguage(lang language.Tag) {
	langValue := lang.String()
	if translator, ok := Self.translators.Get(langValue); ok {
		Self.defaultLanguage = langValue
		Self.defaultTranslator = translator
	} else {
		Self.defaultLanguage = langValue
		translator = NewTranslator()
	}
}

func (Self *Bundle) AddTranslator(lang string, translator *Translator) {
	Self.translators.Put(lang, translator)
	if Self.defaultLanguage == lang {
		Self.defaultTranslator = translator
	}
}

func (Self *Bundle) AddTemplates(lang string, templates map[string]string) error {
	translator := Self.translators.GetOrCreate(lang, func() *Translator { return NewTranslator() })
	return translator.AddTemplates(templates)
}

func (Self *Bundle) AddTemplate(lang, code, template string) error {
	translator := Self.translators.GetOrCreate(lang, func() *Translator { return NewTranslator() })
	return translator.AddTemplate(code, template)
}

func (Self *Bundle) GetTranslator(lang any) *Translator {
	translator := Self.defaultTranslator
	switch langType := lang.(type) {
	case string:
		if value, ok := Self.translators.Get(langType); ok {
			translator = value
		}
	case context.Context:
		locale := langType.Value("locale")
		if locale == nil {
			locale = langType.Value("language")
		}
		if locale == nil {
			locale = Self.defaultLanguage
		}
		translator = Self.GetTranslator(locale)
	case language.Tag:
		if value, ok := Self.translators.Get(langType.String()); ok {
			translator = value
		} else {
			base, _ := langType.Base()
			if value, ok = Self.translators.Get(base.String()); ok {
				translator = value
				Self.translators.Put(langType.String(), translator)
			}
		}
	}
	return translator
}
