package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/Rockup-Consulting/go_std/core/netutil"
	"github.com/Rockup-Consulting/go_std/x/hashx"
	"github.com/Rockup-Consulting/go_std/x/randx"
)

// Now prints the current date with a timestamp
func Now() {
	now := time.Now()

	fmt.Printf("%d-%d-%d\n%d:%d:%d\n", now.Year(), now.Day(), now.Hour(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(now.Unix())

}

// MyIP prints your current IP address
func MyIP() error {
	resp, err := http.Get("http://ip-api.com/json")
	if err != nil {
		return err
	}

	respBody := struct {
		Status      string
		Country     string
		CountryCode string
		Region      string
		RegionName  string
		City        string
		Zip         string
		Lat         float32
		Lon         float32
		Timezone    string
		Isp         string
		Org         string
		As          string
		Query       string
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return err
	}

	fmt.Printf("Public:\t%s\n", respBody.Query)

	local, hasLocal := netutil.GetHostIP()
	if hasLocal {
		fmt.Printf("Local:\t%s\n", local)
	}

	return nil
}

// Token prints a random 32byte string
func UID() error {
	if uid, err := randx.UID(); err != nil {
		return err
	} else {
		fmt.Println(uid)
		return nil
	}
}

// PIN prints a random 6 character alphanumeric string
func PIN() error {
	pin, err := randx.UID()
	fmt.Println(pin)
	return err
}

func Kill(port int) error {
	cmd := exec.Command("lsof", "-i", ":4000")
	var buf bytes.Buffer
	cmd.Stdout = &buf

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println(buf.String())

	return nil
}

func Hash(password string) {
	fmt.Println(hashx.BcryptNew(password))
}

func MD5(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	out, err := hashx.MD5(f)
	if err != nil {
		return err
	}

	fmt.Println(out)
	return nil
}

// func Build() error {
// 	info, err := buildutil.NewInfo("rockup", "site", "0.0.1", "linux", "amd64")
// 	if err != nil {
// 		return err
// 	}

// 	buildutil.EmbedBuildInfo(info, func(info buildutil.Info) error {
// 		return buildutil.BuildGoBin(info)
// 	})

// 	return nil
// }
