package gosubst

import (
	"html/template"
	"os"
)

func (c *CLI) parse() error {
	tmpl, err := template.New("myTmpl").Parse("Hello, {{.}} ")
	if err != nil {
		return err
	}

	tmpl.Execute(os.Stdout, "World")

	// TODO: impl go/template
	input := c.input
	result := input
	c.result = result
	return nil
}
