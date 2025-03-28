package main

import (
	"bytes"
	"context"
	b64 "encoding/base64"
	"github.com/kbinani/screenshot"
	"image/png"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) encodeToBase64Image(data []byte) string {
	base64 := "data:image/png;base64," + b64.StdEncoding.EncodeToString(data)
	return base64
}

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

		base64Image := a.encodeToBase64Image(buffer.Bytes())
		base64Images = append(base64Images, base64Image)
	}

	return base64Images, nil
}
