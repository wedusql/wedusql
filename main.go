package main

import (
	"embed"
	"wedusql/backend/app"
	"wedusql/backend/connections"
	"wedusql/backend/queries"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func bind(app *app.App) []any {
	return []any{
		app,
		queries.NewQuery(),
		connections.NewConnection(),
	}
}

func main() {
	// Create an instance of the app structure
	app := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:      "wedusql",
		Fullscreen: false,
		Frameless:  false,
		Width:      1024,
		Height:     768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Menu:             app.MakeMenu(),
		Bind:             bind(app),
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
		LogLevel:           logger.INFO,
		LogLevelProduction: logger.ERROR,
		StartHidden:        false,
		Windows:            &windows.Options{},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: false,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
		},
		Linux: &linux.Options{},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
