package signaler

// Receiver is the target for the os.Signal detected.
type Receiver interface {
	OnReload()
	OnStop()
}
