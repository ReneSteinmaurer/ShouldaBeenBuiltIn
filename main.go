package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"win_tools/src/api"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := api.NewApp()

	appOptions := &options.App{
		Title:  "ShouldaBeenBuiltIn",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
			BackdropType:         windows.Acrylic,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	}
	app.SetAppOptions(appOptions)

	err := wails.Run(appOptions)

	if err != nil {
		println("Error:", err.Error())
	}
}
