package utils

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"image"
	"image/png"
)

func EncodeToBase64Image(data []byte) string {
	base64 := "data:image/png;base64," + b64.StdEncoding.EncodeToString(data)
	return base64
}

func ConvertToPng(image image.Image) (bytes.Buffer, error) {
	var buf bytes.Buffer
	err := png.Encode(&buf, image)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("error encoding image: %v", err)
	}
	return buf, nil
}
