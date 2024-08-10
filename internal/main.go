package main

import (
	"context"
	"syscall/js"

	"github.com/a-h/templ"
	"github.com/monta-k/go_wasm_profile/internal/template"
)

func main() {
	js.Global().Set("serveHTML", js.FuncOf(serveHTML))
	select {}
}

func serveHTML(this js.Value, p []js.Value) interface{} {
	// base64Image := image.ConvertToBase64ImageFromJSData(p[0])

	templHTML := template.App()
	html, err := templ.ToGoHTML(context.Background(), templHTML)
	if err != nil {
		return js.ValueOf("")
	}

	return js.ValueOf(string(html))
}
