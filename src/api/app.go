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
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	go func() {
		time.Sleep(100 * time.Millisecond)
		a.hwnd = win.FindWindow(nil, syscall.StringToUTF16Ptr("ShouldaBeenBuiltIn"))
	}()
}

func (a *App) SetAppOptions(options *options.App) {
	a.appOptions = options
}
