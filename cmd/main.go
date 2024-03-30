package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/coreos/go-systemd/v22/dbus"
)

var done = make(chan any, 1)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		logE.Fatalf("error loading config: %s\n", err.Error())
	}
	logI.Println(cfg.Motd)

	// TODO: https://github.com/coreos/go-systemd/tree/main/daemon

	attachDbus()

	handleSignals()
}

func attachDbus() {
	ctx := context.Background()

	bus, err := dbus.NewWithContext(ctx)
	if err != nil {
		logE.Fatalf("error connecting to dbus: %s\n", err.Error())
	}
	defer bus.Close()

	_, err = bus.ListUnitsContext(ctx)
	if err != nil {
		logE.Fatalf("error fetching units: %s\n", err.Error())
	}

	// TODO: dbus event handlers
}

func handleSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-done:
			logI.Println("work is done, shutting down janitord...")
			os.Exit(0)
		case sig := <-signals:
			switch sig {
			case syscall.SIGHUP:
				// TODO: reload config
			case syscall.SIGINT:
				fallthrough
			case syscall.SIGTERM:
				logI.Printf("signal %s received, shutting down janitord...\n", sig.String())
				os.Exit(0)
			default:
				logE.Printf("signal %s received, no handler found, ignoring...\n", sig.String())
			}
		}
	}
}
