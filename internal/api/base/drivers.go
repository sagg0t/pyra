package base

import (
	"html/template"
)

func Drivers(fms ...template.FuncMap) *template.Template {
	templateDriver := template.New("drivers")

	templateDriver.Funcs(TemplateHelpers)

	for _, fm := range fms {
		templateDriver.Funcs(fm)
	}

	template.Must(templateDriver.ParseGlob("view/layout/*.html"))
	template.Must(templateDriver.ParseGlob("view/errors/*.html"))
	template.Must(templateDriver.ParseGlob("view/components/*.html"))

	return templateDriver
}
