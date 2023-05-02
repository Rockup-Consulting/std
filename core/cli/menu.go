package cli

import (
	"fmt"
	"io"
)

// Menu stores a collection of Executables, meaning you can add Func(s) and Menu(s) to Menu(s). The
// most important rule is that each Executable must have a unique argument. It is also possible to
// Group Executables on a Menu, but this is strictly a visual grouping, the argument uniqueness rule
// still applies at the Menu level.
type Menu struct {
	a            *App
	this         *command
	defaultGroup *Group
	groups       []*Group
	commands     map[string]*command
}

func (m *Menu) Exec(args []string) error {
	if len(args) < 1 {
		return m.executePrint()
	}

	arg := args[0]

	cmnd, ok := m.commands[arg]
	if !ok {
		return errCmndNotFound(arg)
	}

	if cmnd.e == nil {
		return errInternalNilExec(args)
	}

	return cmnd.e.Exec(args[1:])
}

func (m *Menu) Print(w io.Writer) {
	fmt.Fprintf(w, "  %s\t%s\n", m.this.arg, m.this.d)
}

// group is a helper that groups and orders commands
type Group struct {
	a    *App
	name string
	this *command

	// m is a pointer to the group's parent
	m        *Menu
	commands []*command
}

// AddFunc adds an Executable Func to the Menu's default group
func (m *Menu) AddFunc(arg string, description, help string, f Func) {
	if m.defaultGroup == nil {
		m.defaultGroup = &Group{a: m.a, this: m.this, m: m}
	}

	m.defaultGroup.funcInvariants(arg, f)
	cmnd := m.defaultGroup.m.createFunc(arg, description, help, f)

	m.defaultGroup.addCommand(arg, cmnd)
}

func (m *Menu) AddMenu(arg string, description, help string) *Menu {
	if m.defaultGroup == nil {
		m.defaultGroup = &Group{a: m.a, this: m.this, m: m}
	}

	m.defaultGroup.commandInvariants(arg)

	menuCmnd, menu := m.defaultGroup.createMenu(arg, description, help)
	m.defaultGroup.addCommand(arg, menuCmnd)

	return menu
}

// AddGroup adds a Group of to the Menu. Groups are useful for organizing Executables (functions
// and Menu's)
func (m *Menu) AddGroup(name string) *Group {
	g := &Group{a: m.a, name: name, this: m.this, m: m}

	// sanity check - should've been initialized by now
	if m.groups == nil {
		panic(errInternalNilGroupsSlice().Error())
	}

	// switch around stdCommandsGroup with a new group, then append standardCommands to the end
	//
	// the array access should be safe since standard commands gets added before this ever gets
	// called
	groupCount := len(m.groups)
	stdCommandsGroup := m.groups[groupCount-1]
	m.groups[groupCount-1] = g
	m.groups = append(m.groups, stdCommandsGroup)

	return g
}

// AddMenu adds an Executable Menu to the Group
func (g *Group) AddMenu(arg string, description, help string) *Menu {
	g.commandInvariants(arg)

	menuCmnd, menu := g.createMenu(arg, description, help)
	g.addCommand(arg, menuCmnd)

	return menu
}

// AddFunc adds an Executable Func to the Group
func (g *Group) AddFunc(arg string, description, help string, f Func) {
	g.funcInvariants(arg, f)
	cmnd := g.m.createFunc(arg, description, help, f)

	g.addCommand(arg, cmnd)
}

func (g *Group) addCommand(arg string, cmnd *command) {
	if g.m.commands == nil {
		g.m.commands = make(map[string]*command)
	}

	g.m.commands[arg] = cmnd

	if g.commands == nil {
		g.commands = []*command{cmnd}
	} else {
		g.commands = append(g.commands, cmnd)
	}
}

func (a *App) initMenuInvariants() {
	// in this case we already know that the menu hasn't been initialized yet
	if a.c == nil {
		return
	}

	if a.c.e != nil {
		panic(errAttemptToReInitMenu().Error())
	}
}

func (g *Group) commandInvariants(arg string) {
	if arg == "" {
		var menuName string
		if g.this.arg == "" {
			menuName = "[app menu]"
		} else {
			menuName = g.this.arg
		}
		panic(errEmptyArg(menuName).Error())
	}

	_, isReservedKey := stdCommands[arg]
	if isReservedKey {
		panic(errAttemptToUseReservedKey(arg).Error())
	}

	// if the commands has been initialized, we need to ensure that we're not overwriting anything
	if g.m.commands != nil {
		_, ok := g.m.commands[arg]
		if ok {
			panic(errAttemptToOverwriteArg(arg, g.this.arg).Error())
		}
	}
}

func (g *Group) funcInvariants(arg string, f Func) {
	g.commandInvariants(arg)

	if f == nil {
		panic(errAttemptToCreateExecWithNilFunc(arg).Error())
	}
}

// createMenu returns the command and the menu that is wrapped by the command. Otherwise we have no
// way to access the underlying menu to keep extending it
//
// rethink this, we need to keep the abstraction layer below the exec interface
func (g *Group) createMenu(arg string, d, h string) (*command, *Menu) {
	m := &Menu{a: g.a}
	cmnd := &command{
		a:   g.a,
		arg: arg,
		w:   g.this.w,
		e:   m,
		d:   d,
		h:   h,
	}

	// the menu needs to pass a reference to its own command
	m.this = cmnd

	m.createStandardCommands()

	return cmnd, m
}
