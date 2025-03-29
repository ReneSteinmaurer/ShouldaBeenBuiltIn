package api

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/kbinani/screenshot"
	"image"
	"image/draw"
	"image/png"
	"win_tools/src/utils"
)

//go:embed all:frontend/dist
var assets embed.FS

type BoundsOfActiveDisplays struct {
	MinX, MinY, MaxX, MaxY int
	TotalWidth             int
	TotalHeight            int
	NumberOfDisplays       int
}

func NewBoundsOfActiveDisplays() *BoundsOfActiveDisplays {
	bounds := &BoundsOfActiveDisplays{}
	bounds.Calc()
	return bounds
}

func (b *BoundsOfActiveDisplays) Calc() {
	n := screenshot.NumActiveDisplays()
	var minX, minY, maxX, maxY int

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		if i == 0 || bounds.Min.X < minX {
			minX = bounds.Min.X
		}
		if i == 0 || bounds.Min.Y < minY {
			minY = bounds.Min.Y
		}
		if i == 0 || bounds.Max.X > maxX {
			maxX = bounds.Max.X
		}
		if i == 0 || bounds.Max.Y > maxY {
			maxY = bounds.Max.Y
		}
	}

	b.TotalWidth = maxX - minX
	b.TotalHeight = maxY - minY
	b.MinX = minX
	b.MaxX = maxX
	b.MinY = minY
	b.MaxY = maxY
	b.NumberOfDisplays = n
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

		base64Image := utils.EncodeToBase64Image(buffer.Bytes())
		base64Images = append(base64Images, base64Image)
	}

	return base64Images, nil
}

func (a *App) ScreenshotEveryDisplayAndMergeIt(numberOfDisplays, minX, minY, totalWidth, totalHeight int) (*image.RGBA, error) {
	var fullImage *image.RGBA
	fullImage = image.NewRGBA(image.Rect(0, 0, totalWidth, totalHeight))
	for i := 0; i < numberOfDisplays; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			return nil, fmt.Errorf("error capturing display %d: %v", i, err)
		}

		drawAt := image.Point{X: bounds.Min.X - minX, Y: bounds.Min.Y - minY}
		imgRect := image.Rectangle{Min: drawAt, Max: drawAt.Add(bounds.Size())}
		draw.Draw(fullImage, imgRect, img, image.Point{}, draw.Src)
	}
	return fullImage, nil
}

func (a *App) CropSectionFromImage(inputImg *image.RGBA, x, y, width, height float32, totalWidth, totalHeight int) *image.Image {
	relX := int(x)
	relY := int(y)

	if relX < 0 {
		relX = 0
	}
	if relY < 0 {
		relY = 0
	}
	if relX+int(width) > totalWidth {
		width = float32(totalWidth - relX)
	}
	if relY+int(height) > totalHeight {
		height = float32(totalHeight - relY)
	}

	croppedRect := image.Rect(relX, relY, relX+int(width), relY+int(height))
	croppedImage := inputImg.SubImage(croppedRect)
	return &croppedImage
}

func (a *App) CaptureSnippet(x, y, width, height float32) (string, error) {
	bounds := NewBoundsOfActiveDisplays()
	fullImage, err := a.ScreenshotEveryDisplayAndMergeIt(
		bounds.NumberOfDisplays,
		bounds.MinX,
		bounds.MinY,
		bounds.TotalWidth,
		bounds.TotalHeight)
	if err != nil {
		return "", err
	}

	croppedImg := *a.CropSectionFromImage(fullImage, x, y, width, height, bounds.TotalWidth, bounds.TotalHeight)
	buffer, err := utils.ConvertToPng(croppedImg)
	if err != nil {
		return "", err
	}
	return utils.EncodeToBase64Image(buffer.Bytes()), nil
}
