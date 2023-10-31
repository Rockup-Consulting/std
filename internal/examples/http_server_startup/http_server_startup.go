package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rockup-Consulting/std/core/mid"
	"github.com/Rockup-Consulting/std/core/web"
	"github.com/Rockup-Consulting/std/x/logx"
)

var (
	//go:embed assets
	assets embed.FS
)

func main() {
	l := logx.New("| ")

	if err := run(l); err != nil {
		l.Printf("ERROR: %s", err)
		os.Exit(1)
	}
}

func run(l *log.Logger) error {
	// PARSE FLAGS
	// ====================================================================
	dev := flag.Bool("dev", false, "run application in dev mode")
	flag.Parse()

	// SETUP ADDITIONAL DEPENDENCIES
	// ====================================================================

	// - Auth Service
	// - Local Cache
	// - Database Connections
	// - etc.

	// CREATE HTTP SERVER
	// ====================================================================
	appServer := initApp(l)

	server := http.Server{
		Addr:    ":4000",
		Handler: appServer,
	}

	// SETUP LIFECYCLE CHANNELS
	// ====================================================================
	shutdown := make(chan os.Signal, 1)
	serverErr := make(chan error)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT)

	go func() {
		serverErr <- server.ListenAndServe()
	}()

	if *dev {
		l.Println("listening at http://localhost:4000")
	} else {
		l.Println("listening at :4000")
	}

	select {
	case err := <-serverErr:
		return fmt.Errorf("server error: %s", err)
	case sig := <-shutdown:
		l.Printf("starting server shutdown: %s", sig)
		defer l.Println("server shutdown complete")

		ctx, cancel := context.WithTimeout(context.Background(), web.DefaultShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			server.Close()
			return fmt.Errorf("could not gracefully shutdown server: %s", err)
		}
	}

	return nil
}

func initApp(l *log.Logger) http.Handler {
	app := web.NewApp(mid.Log(l), mid.CatchErr(l), mid.CatchPanic())

	app.Handle(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) error {
		return web.Text(w, http.StatusOK, "This is just an example, go build something real!")
	})

	return app
}
