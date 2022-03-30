package vif

import (
	"context"
	"github.com/rs/zerolog/log"
)

type Vif struct {
	// cancelFn is the function to cancel the context.
	cancelFn func()
}

// New returns a Vif shell, this object encapsulate the operation of the inputs, it will
// takes care of the following:
// - Handle operating system signals
// - Connect to Elastic Agent and wait for changes
func New() *Vif {
	return &Vif{}
}

// Run will start the shell and supervise the inputs.
func (v *Vif) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	v.cancelFn = cancel

	log.Info().
		Msg("Vif shell started")

	log.Info().
		Msg("Vif shell ready")

	<-ctx.Done()

	return nil
}

func (v *Vif) stop() {
	v.cancelFn()
	log.Info().
		Msg("Vif shell stopped")
}

func (v *Vif) OnStop() {
	v.stop()
}

func (v *Vif) OnReload() {

}
