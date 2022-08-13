package systemctl

type Ctl struct {
	UserMode bool
	Timeout  int
}

var defaultTimeout = 3
var userMode = false

var systemOpts = Options{
	UserMode: false,
	Timeout:  defaultTimeout,
}

func New(inst *Ctl) *Ctl {
	defaultTimeout = inst.Timeout
	userMode = userMode
	return inst
}
