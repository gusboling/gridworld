package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func runGridSimulator() {
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigTerm)

	go func() {
		<-sigTerm

		fmt.Println("GridSimulator received shutdown command") // TODO(gus): replace with logger

		// Perform shutdown-cleanup tasks here

		fmt.Println("GridSimulator shutting down") // TODO(gus): replace with logger
	}()
}

func main() {
	runGridSimulator()
}
