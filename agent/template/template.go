// Package template ...
package template

import (
	"bytes"
	"fmt"
	"os"
	text_template "text/template"

	"my.domain/lb/agent/common"
)

type Template struct {
	tmpl *text_template.Template
}
type L4TemplateExecuter struct {
	template *Template
	data     *common.L4LbConfig
}

type L7TemplateExecuter struct {
	template *Template
	data     *common.L7LbConfig
}

func NewL4TemplateExecuter(t *Template, data *common.L4LbConfig) *L4TemplateExecuter {
	return &L4TemplateExecuter{
		template: t,
		data:     data,
	}
}

func NewL7TemplateExecuter(t *Template, data *common.L7LbConfig) *L7TemplateExecuter {
	return &L7TemplateExecuter{
		template: t,
		data:     data,
	}
}

func ExecuteL4(t *L4TemplateExecuter) (string, error) {
	w := new(bytes.Buffer)
	if err := t.template.tmpl.Execute(w, t.data); err != nil {
		return "", err
	}
	return w.String(), nil
}

func ExecuteL7(t *L7TemplateExecuter) (string, error) {
	w := new(bytes.Buffer)
	if err := t.template.tmpl.Execute(w, t.data); err != nil {
		return "", err
	}
	return w.String(), nil
}

// NewL7Template ...
func NewL7Template(file string) *Template {
	tmpl, err := NewTemplate(file, l7FuncMap)
	if err != nil {
		return nil
	}
	return tmpl
}

// NewL4Template ...
func NewL4Template(file string) *Template {
	tmpl, err := NewTemplate(file, l4FuncMap)
	if err != nil {
		return nil
	}
	return tmpl
}

// NewTemplate returns a new Template instance or an
// error if the specified template file contains errors
func NewTemplate(file string, funcMap text_template.FuncMap) (*Template, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("unexpected error reading template %s: %w", file, err)
	}
	tmpl, err := text_template.New("").Funcs(funcMap).Parse(string(data))
	if err != nil {
		return nil, err
	}

	return &Template{
		tmpl: tmpl,
	}, nil
}

var (
	l7FuncMap = text_template.FuncMap{}
)
var (
	l4FuncMap = text_template.FuncMap{}
)
