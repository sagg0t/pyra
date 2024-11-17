package main

import (
	"html/template"
	"os"

	"pyra/internal/api"
)

func main() {
	drivers := template.New("drivers")

	drivers.Funcs(api.TemplateHelpers)
	template.Must(drivers.ParseGlob("view/layout/*.html.tmpl"))
	template.Must(drivers.ParseGlob("view/errors/*.html.tmpl"))

	signIn, err := drivers.Clone()
	if err != nil {
		panic(err)
	}

	template.Must(signIn.ParseFiles("view/auth/sign_in.html.tmpl"))

	err = signIn.ExecuteTemplate(os.Stdout, "sign-in", nil)
	if err != nil {
		panic(err)
	}
}
