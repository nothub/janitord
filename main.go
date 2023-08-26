package main

import (
	"context"
	"fmt"
	"github.com/coreos/go-systemd/v22/dbus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var done = make(chan any, 1)

func main() {
	fmt.Println(cfg.Motd)

	bus, err := dbus.NewWithContext(context.Background())
	if err != nil {
		log.Fatalf("error connecting to dbus: %s\n", err.Error())
	}
	defer bus.Close()

	_, err = bus.ListUnitsContext(context.Background())
	if err != nil {
		log.Fatalf("error fetching units: %s\n", err.Error())
	}

	// TODO: dbus event handlers

	// TODO: https://github.com/coreos/go-systemd/tree/main/daemon

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-done:
			log.Println("work is done, shutting down janitord...")
			os.Exit(0)
		case <-signals:
			log.Println("signal received, shutting down janitord...")
			os.Exit(0)
		}
	}
}
