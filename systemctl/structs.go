package systemctl

type Options struct {
	UserMode bool `json:"user_mode"`
	Timeout  int  `json:"timeout"`
}
