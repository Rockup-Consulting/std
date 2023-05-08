package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/Rockup-Consulting/go_std/x/twx"
)

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
