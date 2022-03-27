package main

import (
	"atlas-csc/buff"
	"atlas-csc/kafka"
	"atlas-csc/logger"
	"atlas-csc/skill"
	tasks "atlas-csc/task"
	"atlas-csc/tracing"
	"context"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const serviceName = "atlas-csc"
const consumerGroupId = "Character Skill Coordinator"

func main() {
	l := logger.CreateLogger(serviceName)
	l.Infoln("Starting main service.")

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	tc, err := tracing.InitTracer(l)(serviceName)
	if err != nil {
		l.WithError(err).Fatal("Unable to initialize tracer.")
	}
	defer func(tc io.Closer) {
		err := tc.Close()
		if err != nil {
			l.WithError(err).Errorf("Unable to close tracer.")
		}
	}(tc)

	kafka.CreateConsumers(l, ctx, wg,
		skill.ApplySkillConsumer(consumerGroupId),
		skill.ApplyMonsterMagnetConsumer(consumerGroupId))

	go tasks.Register(buff.ExpireTask(l))

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Infof("Initiating shutdown with signal %s.", sig)
	cancel()
	wg.Wait()
	l.Infoln("Service shutdown.")
}
