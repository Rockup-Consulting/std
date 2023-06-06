package cli

import "context"

type Executable interface {
	Exec(ctx context.Context, args ...string) error
	print() string
}

type App struct {
	menu *Menu
}

func NewApp(name, overview string) (*App, *Menu) {
	menu := &Menu{
		name:     name,
		overview: overview,
		ordered:  make([]Executable, 0),
		indexed:  make(map[string]Executable),
	}

	app := &App{menu}

	return app, menu
}

func (a *App) Run(args ...string) error {
	ctx := context.Background()

	return a.menu.Exec(ctx, args...)
}
