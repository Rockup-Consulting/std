package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/Rockup-Consulting/std/x/twx"
)

// ====================================================================
// CLI

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

// ====================================================================
// FUNC

type Func func(ctx context.Context, args ...string) error

type fn struct {
	f    Func
	name string
	desc string
}

func (f fn) Exec(ctx context.Context, args ...string) error {
	return f.f(ctx, args...)
}

func (f fn) print() string {
	return fmt.Sprintf("\t%s\t%s\n", f.name, f.desc)
}

// ====================================================================
// GROUP

type group string

// Exec is a dummy method implemented to comply with the Executable interface
func (g group) Exec(ctx context.Context, args ...string) error {
	return nil
}

func (g group) print() string {
	return fmt.Sprintf("\n%s:\n", g)
}

// ====================================================================
// MENU

type Menu struct {
	name     string
	desc     string
	overview string
	indexed  map[string]Executable
	ordered  []Executable

	// this is just to decide whether a space must be printed
	// below the overview or not.
	firstEntryIsGroup bool
}

func (m *Menu) AddFunc(name, desc string, f Func) {
	validateName(name)

	ff := fn{
		f:    f,
		name: name,
		desc: desc,
	}

	m.ordered = append(m.ordered, ff)
	m.indexed[name] = ff
}

func (m *Menu) AddGroup(name string) {
	if len(m.ordered) == 0 {
		m.firstEntryIsGroup = true
	}

	m.ordered = append(m.ordered, group(name))
}

func (m *Menu) AddMenu(name, desc, overview string) *Menu {
	mm := &Menu{
		name:     name,
		desc:     desc,
		overview: overview,
		ordered:  make([]Executable, 0),
		indexed:  make(map[string]Executable),
	}

	m.ordered = append(m.ordered, mm)
	m.indexed[name] = mm

	return mm
}

func (m Menu) Exec(ctx context.Context, args ...string) error {
	// if there are no further args, we need to print the menu
	if len(args) == 0 {
		m.printMenu()
		return nil
	}

	// if there is atleast one more arg, we need to execute the
	// executable
	exec, ok := m.indexed[args[0]]
	if !ok {
		return fmt.Errorf("arg %q not found", args[0])
	}

	return exec.Exec(ctx, args[1:]...)
}

func (m *Menu) print() string {
	return fmt.Sprintf("\t%s\t%s", m.name, m.desc)
}

func (m *Menu) printMenu() {
	var builder strings.Builder
	tw := twx.NewWriter(&builder)

	fmt.Fprintf(tw, "%s\n", m.overview)
	if !m.firstEntryIsGroup {
		fmt.Fprint(tw, "\n")
	}

	for _, e := range m.ordered {
		fmt.Fprint(tw, e.print())
	}

	tw.Flush()

	fmt.Print(builder.String())
}

// ====================================================================
// INVARIANTS

func initErr(s string) error {
	return fmt.Errorf("CLI Init Error: %s", s)
}

func validateName(name string) {
	if strings.Contains(name, " ") {
		panic(initErr(fmt.Sprintf("invalid arg name %q, spaces not allowed", name)))
	}
}
