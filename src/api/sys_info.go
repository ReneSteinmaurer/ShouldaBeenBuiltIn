package api

import (
	"fmt"
	"github.com/kbinani/screenshot"
)

type WindowCoordinates struct {
	TotalWidth  int
	TotalHeight int
	MinX        int
	MinY        int
}

func (a *App) GetMaximizedWindowCoordinates() *WindowCoordinates {
	n := screenshot.NumActiveDisplays()
	var minX, minY, maxX, maxY int

	if n > 0 {
		bounds := screenshot.GetDisplayBounds(0)
		minX = bounds.Min.X
		minY = bounds.Min.Y
		maxX = bounds.Max.X
		maxY = bounds.Max.Y
	}

	for i := 1; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		if bounds.Min.X < minX {
			minX = bounds.Min.X
		}
		if bounds.Min.Y < minY {
			minY = bounds.Min.Y
		}
		if bounds.Max.X > maxX {
			maxX = bounds.Max.X
		}
		if bounds.Max.Y > maxY {
			maxY = bounds.Max.Y
		}
	}

	totalWidth := maxX - minX
	totalHeight := maxY - minY

	fmt.Printf("Virtuelles Display: Min(%d,%d) Max(%d,%d) - Größe: %dx%d\n",
		minX, minY, maxX, maxY, totalWidth, totalHeight)
	return &WindowCoordinates{
		TotalWidth:  totalWidth,
		TotalHeight: totalHeight,
		MinX:        minX,
		MinY:        minY,
	}
}
