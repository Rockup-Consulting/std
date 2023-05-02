package cli

import (
	"fmt"
	"strings"
	"text/template"
)

var (
	stdCommands = map[string]bool{
		"help": true,
		"tree": true,
	}
)

func (m *Menu) createStandardCommands() {
	// sanity check - this should be called before any other groups have been added
	if m.groups != nil {
		panic(errInternal().Error())
	}

	if m.commands != nil {
		panic(errInternal().Error())
	}

	m.commands = make(map[string]*command)

	g := &Group{
		name:     "standard",
		this:     m.this,
		m:        m,
		commands: []*command{},
	}

	m.groups = []*Group{g}

	createHelp(m, g)
}

var (
	helpTextTmplStr = `{{.Description}}
	
Calling 'help' will display the help text for a command or subcommand (if any has been 
provided). 'help' can be called in a couple of ways, examples (with made up commands):

- {{.AppName}} logs help stream {{"\t"}} - display the help text for 'stream'
- {{.AppName}} logs help {{"\t"}}{{"\t"}} - display the help text for 'logs'
- {{.AppName}} help logs {{"\t"}}{{"\t"}} - display the help text for 'logs'
- {{.AppName}} help	{{"\t"}}{{"\t"}} - display the help text for '{{.AppName}}'

Note:
- The second and third examples are the same
- All the commands and subcommands are made up, you might not have those commands in your app

The help command has two possible formats:

- {{.AppName}} <...command> help {{"\t"}}{{"\t"}}{{"\t"}} - displays the help text for <command>
- {{.AppName}} <...command> help <subcommand> {{"\t"}} - displays the help text for <subcommand>
`
)

func createHelp(m *Menu, g *Group) {
	helpTextTmpl, err := template.New("help").Parse(helpTextTmplStr)
	if err != nil {
		panic(err.Error())
	}

	var helpStr strings.Builder

	err = helpTextTmpl.Execute(&helpStr, struct {
		AppName     string
		Description string
	}{
		AppName:     m.a.name,
		Description: "Get further help for commands",
	})

	if err != nil {
		panic(err.Error())
	}

	help := m.createFunc(
		"help",
		"Get further help for commands.",
		helpStr.String(),
		func(args []string) error {

			if len(args) < 1 {
				// no further args, display command help text
				fmt.Fprintln(m.this.w, m.this.h)

			} else if len(args) > 1 {
				// help only accepts one arg
				return errTooManyArgsProvided("help")
			} else {
				arg := args[0]

				cmnd, ok := m.commands[arg]
				if !ok {
					return errCmndNotFound(arg)
				}

				fmt.Fprintln(m.this.w, cmnd.h)
			}

			return nil
		},
	)

	g.commands = append(g.commands, help)
	m.commands["help"] = help
}
