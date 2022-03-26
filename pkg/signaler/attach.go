package signaler

import (
	"context"
)

// +build: !windows

// Attach is keeping tack of the operating system system signals and will call the receiver
// callbacks, the attach assume that the receiver' callbacks are reentrant,
// it's the responsability of the receiver to managed his current state and making sure the
// operation is not happening twice if needed.
func Attach(ctx context.Context, s Receiver) {
	// c := make(chan os.Signal, 1)
	// go func() {
	// 	for {
	// 		select {}
	// 	}
	// }()
}
