package bootstrap

type System struct {
}

// New is returning a new system, the system is initalized but not ready to receive any events,
// users are expected to call Start().
func New() *System {
	return System{}
}

func Start() {
}

func Stop() {

}
