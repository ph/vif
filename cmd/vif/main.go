package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.elastic.co/ecszerolog"
	"os"
)

func main() {
	// TODO(ph) Add build information/version
	// TODO(ph) Make sure we can link logging to the agent policy.
	// TODO(ph) use severity hook with the level.Discard to filter specific event
	logger := ecszerolog.New(os.Stdout)
	log.Logger = logger

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Info().Msg("hello world")
	log.Debug().Msg("Have some fun")
}
