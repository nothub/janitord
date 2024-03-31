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

const usage string = `janitord - The janitor juggles services`

var done = make(chan any, 1)

func main() {

	flag.Usage = func() {
		log.Print(usage)
		os.Exit(1)
	}

	flag.StringVar(&cfgPath, "config", "/etc/janitord.yaml", "")

	flag.Parse()

	err := loadConfig()
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

	sysState, err := bus.SystemStateContext(ctx)
	if err != nil {
		logE.Fatalf("error fetching system state: %s\n", err.Error())
	}
	logI.Printf("%s=%+v (%T)\n", sysState.Name, sysState.Value, sysState.Value)

	unitStates, err := bus.ListUnitsContext(ctx)
	if err != nil {
		logE.Fatalf("error fetching unit states: %s\n", err.Error())
	}
	logI.Printf("found %v unit states:\n", len(unitStates))
	for _, s := range unitStates {
		logI.Printf("  - %s (%s) %s %s %s\n", s.Name, s.JobType, s.LoadState, s.ActiveState, s.SubState)
	}

	unitFiles, err := bus.ListUnitFilesContext(ctx)
	if err != nil {
		logE.Fatalf("error fetching unit files: %s\n", err.Error())
	}
	logI.Printf("found %v unit files:\n", len(unitStates))
	for _, f := range unitFiles {
		logI.Printf("  - %s %s\n", f.Type, f.Path)
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
				err := loadConfig()
				if err != nil {
					logE.Fatalf("error loading config: %s\n", err.Error())
				}
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
