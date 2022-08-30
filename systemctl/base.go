package systemctl

type Ctl struct {
	UserMode bool
	Timeout  int
}

func New(userMode bool, timeout int) *Ctl {
	instance := Ctl{UserMode: userMode, Timeout: timeout}
	return &instance
}
