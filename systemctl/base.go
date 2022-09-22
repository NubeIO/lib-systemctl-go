package systemctl

type SystemCtl struct {
	UserMode bool
	Timeout  int
}

func New(userMode bool, timeout int) *SystemCtl {
	instance := SystemCtl{UserMode: userMode, Timeout: timeout}
	return &instance
}
