// Package cli is a light-weight cli application framework.

// The goal of cli is simple - make it easy to create *bug free* cli applications. cli doesn't allow
// user's to make any logical errors by:
//
//   - not allowing overwriting existing commands
//
//   - not allowing nil commands
//
//   - running strict invariant checks on initialisation
//
// Each cli application should be created by running `cli.NewApp()`. This is where you should start
// reading to get up and running quickly. The rest of this section will be dedicated to covering the
// package jargon in depth. Feel free to refer back to the jargon when you're more familiar with the
// package API.
//
// There are 6 main terms that must be grasped to aquire a deep understanding of this package. They
// are: Executable, command, arg, App, Func, Menu and Group/Default Group. This is quite a lot, but
// to create complicated applications you should understand all the terms:
//
//   - Executable is the main interface used by the framework. Technically the framework would be
//     able to run anything that implements Executable.
//
//   - command wraps an Executable. As a user of the API, you never actually have to work with
//     commands, but it is useful to understand the hierarchy.
//
//   - arg(ument) is used to invoke a command, for example, if we call `app menu1 func1` then
//     `menu1` is an arg that invokes a command on the app menu, and func1 is an arg that invokes a
//     command on `menu1`.
//
//   - App is the type that the user of this package will work with. It is the entrypoint to the
//     framework. In reality, the app just wraps a command and that command has a Menu. So when you
//     call `app, menu := cli.NewApp(...)` that is the Menu that you will receive.
//
//   - Func is the most basic type of Executable you can add to an App. It basically just runs a
//     function for you.
//
//   - Menu  is slightly more complicated. A Menu stores a collection of Executables. Meaning you
//     can either add Funcs or Menu's to a Menu.
//
//   - Group(s) are simply used for visual organisation of Menu's. Rules regarding uniqueness still
//     apply at a Menu level. When adding an Executable to a Group, you should consider it the same
//     as adding it to the parent Menu.
//
//   - Default Group is completely internal to the package. It basically means you can add
//     Executables directly to the Menu, instead of having to create a Group.
//
// As mentioned before, there is a lot of jargon. The recommendation is to get started with the
// NewApp example. That is the most complete reference you'll find.
package cli
