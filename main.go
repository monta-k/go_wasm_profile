package main

import (
	"fmt"
	"syscall/js"

	"github.com/monta-k/go_wasm_profile/internal"
)

func main() {
	js.Global().Set("serveHTML", js.FuncOf(serveHTML))
	select {}
}

func serveHTML(this js.Value, p []js.Value) interface{} {
	base64Image := internal.ConvertToBase64ImageFromJSData(p[0])
	imageTag := internal.GenerageLogoImgTag(base64Image)

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
		<head>
			<title>Go WebAssembly</title>
		</head>
		<body>
			%s
		</body>
	</html>
	`, imageTag)
	return js.ValueOf(html)
}
