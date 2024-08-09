package template

import "github.com/a-h/templ"

func App(logoBase64ImageData string) templ.Component {
	return app(logoBase64ImageData)
}
