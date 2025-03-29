package api

import (
	"context"
	"github.com/kbinani/screenshot"
)

type SysInfo struct {
	ctx               context.Context
	WindowCoordinates *WindowCoordinates
}

type WindowCoordinates struct {
	TotalWidth  int
	TotalHeight int
	MinX        int
	MinY        int
}

func NewWindowCoordinates() *WindowCoordinates {
	w := &WindowCoordinates{}
	w.CalcMaximizedWindowCoordinates()
	return w
}

func NewSysInfo(ctx context.Context) *SysInfo {
	s := &SysInfo{
		ctx:               ctx,
		WindowCoordinates: NewWindowCoordinates(),
	}
	return s
}

func (w *WindowCoordinates) CalcMaximizedWindowCoordinates() {
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

	w.TotalWidth = maxX - minX
	w.TotalHeight = maxY - minY
	w.MinX = minX
	w.MinY = minY
}
