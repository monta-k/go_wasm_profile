package main

import "syscall/js"

func main() {
	js.Global().Set("serveHTML", js.FuncOf(serveHTML))
	select {}
}

func serveHTML(this js.Value, p []js.Value) interface{} {
	html := `
	<!DOCTYPE html>
	<html>
		<head>
			<title>Go WebAssembly</title>
		</head>
		<body>
			<h1>Hello, Go WebAssembly!!!</h1>
		</body>
	</html>
	`
	return js.ValueOf(html)
}
