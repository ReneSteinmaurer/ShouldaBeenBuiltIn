package utils

import b64 "encoding/base64"

func EncodeToBase64Image(data []byte) string {
	base64 := "data:image/png;base64," + b64.StdEncoding.EncodeToString(data)
	return base64
}
