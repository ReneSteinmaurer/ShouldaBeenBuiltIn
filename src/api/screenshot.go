package api

import (
	"bytes"
	"github.com/kbinani/screenshot"
	"image/png"
	"win_tools/src/utils"
)

func (a *App) CaptureArea() ([]string, error) {
	n := screenshot.NumActiveDisplays()
	var base64Images []string

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			return nil, err
		}

		buffer := new(bytes.Buffer)
		err = png.Encode(buffer, img)
		if err != nil {
			return nil, err
		}

		base64Image := utils.EncodeToBase64Image(buffer.Bytes())
		base64Images = append(base64Images, base64Image)
	}

	return base64Images, nil
}
