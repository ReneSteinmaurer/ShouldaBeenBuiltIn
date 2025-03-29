package api

import (
	"context"
	"github.com/lxn/win"
	"github.com/wailsapp/wails/v2/pkg/options"
	"syscall"
	"time"
)

// App struct
type App struct {
	ctx        context.Context
	appOptions *options.App
	hwnd       win.HWND
	screenshot *Screenshot
	sysInfo    *SysInfo
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.screenshot = NewScreenshot(ctx)
	a.sysInfo = NewSysInfo(ctx)
	go func() {
		time.Sleep(100 * time.Millisecond)
		a.hwnd = win.FindWindow(nil, syscall.StringToUTF16Ptr("ShouldaBeenBuiltIn"))
	}()
}

func (a *App) SetAppOptions(options *options.App) {
	a.appOptions = options
}

func (a *App) CaptureSnippet(x, y, width, height float32) (string, error) {
	return a.screenshot.CaptureSnippet(x, y, width, height)
}

func (a *App) CaptureArea() ([]string, error) {
	return a.screenshot.CaptureArea()
}

func (a *App) GetWindowSize() map[string]int {
	return map[string]int{
		"height": a.sysInfo.WindowCoordinates.TotalHeight,
		"width":  a.sysInfo.WindowCoordinates.TotalWidth,
	}
}
