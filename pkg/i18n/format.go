package i18n

import (
	"bytes"
	gotemplate "text/template"
)

type messageFormat struct {
	message  string
	template *gotemplate.Template
}

func (Self *messageFormat) format(args map[string]string) string {
	var buffer bytes.Buffer
	err := Self.template.Execute(&buffer, args)
	if err != nil {
		return Self.message
	}
	return buffer.String()
}

func newMessageFormat(text string) (*messageFormat, error) {
	template, err := gotemplate.New("").Delims("${", "}").Parse(text)
	if err != nil {
		return nil, err
	}
	return &messageFormat{
		message:  text,
		template: template,
	}, nil
}
