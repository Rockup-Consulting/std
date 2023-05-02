package cli_test

import (
	"errors"
	"fmt"
	"os"
	"rockup/core/cli"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("basic exec func test", func(t *testing.T) {
		var ok bool

		var sb strings.Builder

		app, menu := cli.NewApp(&sb, "test", "test app", "help")
		testGroup := menu.AddGroup("test group")

		// Menu -> Func (on default group)

		ok = false
		menu.AddFunc("menu-func", "test", "help", func(args []string) error {
			ok = true
			return nil
		})
		app.Run("menu-func")
		if !ok {
			t.Error("expected 'ok' to be true, but got false")
		}

		// Menu -> Menu -> Func (menu on default Group)

		ok = false
		menuGroup := menu.AddMenu("menu", "test", "help").AddGroup("menuGroup")
		menuGroup.AddFunc("menu-func", "test", "help", func(args []string) error {
			ok = true
			return nil
		})
		app.Run("menu", "menu-func")
		if !ok {
			t.Error("expected 'ok' to be true, but got false")
		}

		// Menu -> Group -> Func

		ok = false

		testGroup.AddFunc("menu-group-func", "test", "help", func(args []string) error {
			ok = true
			return nil
		})
		app.Run("menu-group-func")
		if !ok {
			t.Error("expected 'ok' to be true, but got false")
		}

		// Menu -> Group -> Menu -> Func

		ok = false
		menuGroup = testGroup.AddMenu("menu2", "test", "help").AddGroup("menuGroup")
		menuGroup.AddFunc("group-menu-func", "test", "help", func(args []string) error {
			ok = true
			return nil
		})
		app.Run("menu2", "group-menu-func")
		if !ok {
			t.Error("expected 'ok' to be true, but got false")
		}

		// add some arg
		menuGroup.AddFunc("group-menu-func-arg", "test", "help", func(args []string) error {
			if len(args) < 1 {
				t.Fatalf("expected len(args) == 1, but got %d", len(args))
			}

			if args[0] != "arg" {
				t.Errorf("expected args[0] to be 'arg', but got %s", args[0])
			}

			return nil
		})

		app.Run("menu", "group-menu-func-arg", "arg")

	})
}

func ExampleNewApp() {
	app, menu := cli.NewApp(os.Stdout, "example", "this is an example cli app", "help me")

	// add a simple Func to the App Menu. Note that args are unique on a Menu, so if you ran
	// menu.AddFunc("greet", ...) again after this, the framework would panic, this is to avoid
	// unnescessary bugs.
	menu.AddFunc(
		"greet",
		"greet's the user on their name",
		"you should pass your name in as an argument",
		func(args []string) error {
			if len(args) != 1 {
				return errors.New("not enough args")
			}

			fmt.Println("Hello", args[0]+"!")

			return nil
		},
	)

	// typically when running a CLI application, you would want to run app.Run(os.Args[1:]...)
	// this would pass all the command-line arguments (without the application's name) through to
	// the framework
	app.Run("greet", "Ferdi")

	// add a Menu with a Func to the App Menu
	nestedMenu := menu.AddMenu(
		"nestedMenu",
		"this menu is nested",
		"I don't know what to say here",
	)

	// since we are calling AddFunc on a nested menu, we can re-use the previous arg. We can nest
	// menus as deep as we want.
	nestedMenu.AddFunc(
		"greet",
		"greet's the user on their name",
		"you should pass your name in as an argument",
		func(args []string) error {
			if len(args) != 1 {
				return errors.New("not enough args")
			}

			fmt.Println("Hello again", args[0]+"!")

			return nil
		},
	)
	app.Run("nestedMenu", "greet", "Ferdi")

	// Output:
	// Hello Ferdi!
	// Hello again Ferdi!
}

func ExampleGroup() {
	app, menu := cli.NewApp(os.Stdout, "example", "this is an example cli app", "help me")

	// we can visially group Executables by adding a Group to a Menu. Note that the argument
	// uniqueness rule still applies at a Menu level, meaning you can't add an arg to a Menu and
	// then add the same arg to a Group, this would cause the framework to panic.
	group := menu.AddGroup("coolGroup")

	group.AddFunc(
		"greet",
		"greet's the user on their name",
		"you should pass your name in as an argument",
		func(args []string) error {
			if len(args) != 1 {
				return errors.New("not enough args")
			}

			fmt.Println("Hello", args[0]+"!")

			return nil
		},
	)

	// note that we don't add a group name when invoking a command on a Group. As mentioned before,
	// a Group is strictly for visual organisation.
	app.Run("greet", "Ferdi")

	// Output:
	// Hello Ferdi!
}
