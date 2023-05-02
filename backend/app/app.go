package app

import (
	"context"
	"runtime"

	_ "github.com/go-mysql-org/go-mysql/driver"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

var db *sqlx.DB

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) MakeMenu() *menu.Menu {
	_menu := menu.NewMenu()

	_menu.Append(menu.AppMenu())

	if runtime.GOOS == "darwin" {
		_menu.Append(menu.EditMenu())
	}

	connectionMenu := _menu.AddSubmenu("Database")
	connectionMenu.AddText("New Connections", keys.CmdOrCtrl("n"), func(_ *menu.CallbackData) {
		wailsRuntime.EventsEmit(a.ctx, "LIST_CONNECTIONS")
	})

	connectionMenu.AddSeparator()
	connectionMenu.AddText("Refresh", keys.CmdOrCtrl("r"), func(_ *menu.CallbackData) {
		wailsRuntime.EventsEmit(a.ctx, "REFRESH_CONNECTION")
	})

	return _menu
}
