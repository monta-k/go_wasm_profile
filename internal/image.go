package internal

import (
	"encoding/base64"
	"fmt"
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

func GenerageLogoImgTag(base64ImageData string) string {
	return fmt.Sprintf(`<img width="600px" src="data:image/png;base64,%s" alt="Site Logo" style="padding-top: 50px; position: absolute; right: 0; transform: rotate(20deg);" />`, base64ImageData)
}
