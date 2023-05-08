package main

import (
	"fmt"

	"github.com/Rockup-Consulting/go_std/core/buildutil"
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
