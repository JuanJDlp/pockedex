package main

import (
	"errors"
	"os"
	"text/template"
)

func inspect(args ...string) error {
	value, ok := pokedex[args[0]]
	if !ok {
		return errors.New("you dont have this pokemon yet")
	}
	const inspectTemplate = `
		Name: {{.Name}}
		Height: {{.Height}}
		Weight: {{.Weight}}
		Stats:
		{{- range .Stats }}
		- {{.Stat.Name}}: {{.BaseStat}}
		{{- end }}
		Types:
		{{- range .Types }}
		- {{.Type.Name}}
		{{- end }}
	`
	t, err := template.New("pokemonTemplate").Parse(inspectTemplate)
	if err != nil {
		return err
	}

	err = t.Execute(os.Stdout, value)
	if err != nil {
		return err
	}

	return nil
}
