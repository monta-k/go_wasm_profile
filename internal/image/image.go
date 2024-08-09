package image

import (
	"encoding/base64"
	"syscall/js"
)

func convertToImageDataFromJSData(input interface{}) []byte {
	uint8Array := js.Global().Get("Uint8Array").New(input)
	length := uint8Array.Get("length").Int()
	imageData := make([]byte, length)
	js.CopyBytesToGo(imageData, uint8Array)
	return imageData
}

func generateBase64Image(imageData []byte) string {
	return base64.StdEncoding.EncodeToString(imageData)
}

func ConvertToBase64ImageFromJSData(input interface{}) string {
	imageData := convertToImageDataFromJSData(input)
	return generateBase64Image(imageData)
}
