package main

import (
	"atlas-csc/buff"
	consumers "atlas-csc/kafka/consumer"
	"atlas-csc/logger"
	tasks "atlas-csc/task"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l := logger.CreateLogger()

	consumers.CreateEventConsumers(l)

	go tasks.Register(buff.ExpireTask(l))

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Infoln("Shutting down via signal:", sig)
}
