package api

import (
	"bytes"
	"embed"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"win_tools/src/utils"
)

//go:embed all:frontend/dist
var assets embed.FS

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

func (a *App) StartScreenshotSelection() {
	a.createScreenshotSelectionWindow()
}

func (a *App) createScreenshotSelectionWindow() {
	n := screenshot.NumActiveDisplays()
	var virtualBounds image.Rectangle

	if n > 0 {
		virtualBounds = screenshot.GetDisplayBounds(0)
	}

	for i := 1; i < n; i++ {
		displayBounds := screenshot.GetDisplayBounds(i)
		virtualBounds = virtualBounds.Union(displayBounds)
	}
}
