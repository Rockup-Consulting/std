package cli

import (
	"io"
)

// Executable is the most fundamental unit of work on the framework. Currently there are only two
// types of Executables - Menu and Func.
type Executable interface {
	Exec(args []string) error
	Print(w io.Writer)
}

// App is the top level unit that a user will work with. The use this package, the user must create
// an App - prefferably by calling cli.NewApp(...)
type App struct {
	w    io.Writer
	name string
	d    string
	h    string
	c    *command // top level command is a menu
}

// NewApp creates and returns a *cli.app and the top-level Menu
func NewApp(w io.Writer, name, description, help string) (*App, *Menu) {
	app := &App{w: w, name: name, d: description, h: help}
	menu := app.initMenu()

	return app, menu
}

// command wraps an Executable and is invoked by passing a argument to the application.
type command struct {
	a   *App
	arg string
	w   io.Writer
	e   Executable
	d   string
	h   string
}

// initMenu initializes the app with the very first command being a menu, the menu is returned for
// further setup
func (a *App) initMenu() *Menu {
	a.initMenuInvariants()
	cmnd := &command{w: a.w, a: a, arg: a.name, d: a.d, h: a.h}
	menu := &Menu{a: a, this: cmnd}

	menu.createStandardCommands()

	cmnd.e = menu
	a.c = cmnd

	return menu
}

// Run passes arguments to the application which starts the invocation of the chain of commands.
//
// Basically it runs your app.
func (a *App) Run(args ...string) error {
	var err error
	if a.c == nil {
		err = errAppMenuNotInitialized()
	} else {
		err = a.c.e.Exec(args)
	}

	return err
}
