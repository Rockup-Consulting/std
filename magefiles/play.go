package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/Rockup-Consulting/go_std/core/buildutil"
	"github.com/Rockup-Consulting/go_std/core/cli"
)

func Play() error {
	info, err := buildutil.NewInfo("", "", "", "", "")
	if err != nil {
		return err
	}

	if err := buildutil.EmbedBuildInfo(info, func(info buildutil.Info) error {
		fmt.Println(info)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func CLI() error {
	app, menu := cli.NewApp("FERDZ", "The FERDZ cli application is siiiick")

	menu.AddFunc("yip", "this does sick shit", func(ctx context.Context, args ...string) error {
		return errors.New("heyo")
	})

	menu.AddGroup("common commands")
	menu.AddFunc("yass", "this does sick shit", func(ctx context.Context, args ...string) error {
		fmt.Println("yasss")
		return nil
	})

	menu.AddFunc("again", "this does more sicker shit", func(ctx context.Context, args ...string) error {
		fmt.Println("yasss")
		return nil
	})

	menu.AddGroup("administration")
	menu.AddFunc("deploy", "this does sick shit", func(ctx context.Context, args ...string) error {
		fmt.Println("yasss")
		return nil
	})

	logs := menu.AddMenu("logs", "logs related commands", "The logs menu provides utilities to inspect or tail logs")

	logs.AddFunc("tail", "this does more sicker shit", func(ctx context.Context, args ...string) error {
		fmt.Println("yasss")
		return nil
	})

	logs.AddGroup("filterby")

	logs.AddFunc("name", "this does more sicker shit", func(ctx context.Context, args ...string) error {
		fmt.Println("yasss")
		return nil
	})

	logs.AddFunc("age", "this does more sicker shit", func(ctx context.Context, args ...string) error {
		fmt.Println("yasss")
		return nil
	})

	return app.Run("logs", "tail")
}
