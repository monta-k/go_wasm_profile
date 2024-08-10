package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/monta-k/go_wasm_profile/internal/template"
)

func main() {
	component := template.App()

	http.Handle("/", templ.Handler(component))
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
