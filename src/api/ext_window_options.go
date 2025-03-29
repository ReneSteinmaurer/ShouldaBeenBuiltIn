package api

import (
	"github.com/lxn/win"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"syscall"
)

func (a *App) MakeWindowTransparent() {
	if a.hwnd == 0 {
		a.hwnd = win.FindWindow(nil, syscall.StringToUTF16Ptr(a.appOptions.Title))
	}

	if a.hwnd != 0 {
		exStyle := win.GetWindowLong(a.hwnd, win.GWL_EXSTYLE)
		win.SetWindowLong(a.hwnd, win.GWL_EXSTYLE, exStyle|win.WS_EX_LAYERED)

		user32 := syscall.NewLazyDLL("user32.dll")
		setLayeredWindowAttributes := user32.NewProc("SetLayeredWindowAttributes")

		setLayeredWindowAttributes.Call(
			uintptr(a.hwnd),
			0,
			50, // (0-255, 0 = 100% transparent)
			2,
		)
	}
}

func (a *App) UndoMakeWindowTransparent() {
	hwnd := win.FindWindow(nil, syscall.StringToUTF16Ptr(a.appOptions.Title))
	if hwnd == 0 {
		return
	}

	exStyle := win.GetWindowLong(hwnd, win.GWL_EXSTYLE)
	win.SetWindowLong(hwnd, win.GWL_EXSTYLE, exStyle & ^win.WS_EX_LAYERED)

	user32 := syscall.NewLazyDLL("user32.dll")
	setLayeredWindowAttributes := user32.NewProc("SetLayeredWindowAttributes")
	setLayeredWindowAttributes.Call(
		uintptr(hwnd),
		0,
		255, // (0-255, 255 = 100% undurchsichtig)
		2,
	)

	user32.NewProc("RedrawWindow").Call(
		uintptr(hwnd),
		0,
		0,
		uintptr(0x0001|0x0002|0x0400|0x0008),
	)
}

func (a *App) MaximizeWindowToBounds() {
	coordinates := a.GetMaximizedWindowCoordinates()
	runtime.WindowSetSize(a.ctx, coordinates.TotalWidth, coordinates.TotalHeight)
	runtime.WindowSetPosition(a.ctx, coordinates.MinX, coordinates.MinY)
	runtime.WindowSetAlwaysOnTop(a.ctx, true)
}

func (a *App) ResetWindowSizeToDefault() {
	runtime.WindowSetSize(a.ctx, a.appOptions.Width, a.appOptions.Height)
	runtime.WindowCenter(a.ctx)
	runtime.WindowSetAlwaysOnTop(a.ctx, false)
}
