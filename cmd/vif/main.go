package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ph/vif/pkg/signaler"
	"github.com/ph/vif/pkg/vif"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.elastic.co/ecszerolog"
)

func main() {
	log.Info().Msg("Process X started")

	// Handle any last resort panic at the root of the application, we will log the error as fatal
	// with the stack trace and we will rely on the supervisor upstream to restart us.
	// TODO(ph) Add the stack trace
	defer panicHandler()

	v := vif.New()

	// Attach the operating system signals to the shell.
	ctx, cancelFn := context.WithCancel(context.Background())
	signaler.Attach(ctx, v)
	defer cancelFn()

	if err := v.Run(); err != nil {
		log.Error().
			Err(err).
			Msg("Process X did not stopped gracefully")

		os.Exit(-1)
		return
	}

	log.Info().Msg("Process X excited gracefully")
	os.Exit(0)
}

func configureDefaultLogging() {
	// TODO(ph) Add information about the agent uuid.
	// TODO(ph) Add build information/version
	// TODO(ph) Make sure we can link logging to the agent policy.
	// TODO(ph) use severity hook with the level.Discard to filter specific event
	logger := ecszerolog.New(os.Stdout)
	log.Logger = logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func panicHandler() {
	if r := recover(); r != nil {
		err := fmt.Errorf("panic: %s", r)

		log.Fatal().
			Stack().
			Err(err).
			Msg("Process X panicked")

		os.Exit(-1)
	}
}
