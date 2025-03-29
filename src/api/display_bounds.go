package api

import "github.com/kbinani/screenshot"

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
