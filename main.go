package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "world-clock",
		Width:  200,
		Height: 100,
		Frameless: true,
		AlwaysOnTop: true,

		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
            WebviewIsTransparent:              true,
            WindowIsTranslucent:               true,
            BackdropType:                      windows.Mica,
        },
        Mac: &mac.Options{
            TitleBar: &mac.TitleBar{
                TitlebarAppearsTransparent: true,
                HideTitle:                  false,
                HideTitleBar:               false,
                FullSizeContent:            false,
                UseToolbar:                 false,
                HideToolbarSeparator:       true,
            },
            Appearance:           mac.NSAppearanceNameDarkAqua,
            WebviewIsTransparent: true,
            WindowIsTranslucent:  true,
            About: &mac.AboutInfo{
                Title:   "My Application",
                Message: "© 2023 Me",
            },
        },
        Linux: &linux.Options{
            WindowIsTranslucent: true,
            WebviewGpuPolicy: linux.WebviewGpuPolicyAlways,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
