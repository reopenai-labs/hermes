package i18n

import (
	"qiyibyte.com/hermes/pkg/builtin/collection"
)

type Translator struct {
	formats *collection.CowMap[string, *messageFormat]
}

func NewTranslator() *Translator {
	return &Translator{
		formats: collection.NewCowMap[string, *messageFormat](),
	}
}

func (Self *Translator) AddTemplates(templates map[string]string) error {
	elements := make(map[string]*messageFormat, len(templates))
	for code, template := range templates {
		format, err := newMessageFormat(template)
		if err != nil {
			return err
		}
		elements[code] = format
	}
	Self.formats.PutAll(elements)
	return nil
}

func (Self *Translator) AddTemplate(code, template string) error {
	format, err := newMessageFormat(template)
	if err != nil {
		return err
	}
	Self.formats.Put(code, format)
	return nil
}

func (Self *Translator) Translate(code string, args map[string]string) string {
	if format, ok := Self.formats.Get(code); ok {
		return format.format(args)
	}
	return code
}
