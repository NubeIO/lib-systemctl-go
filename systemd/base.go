package systemd

import (
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"sync"
)

type conf struct {
	service            string
	systemdDir         string
	systemdSoftLinkDir string
	locker             *sync.Mutex
	Options            systemctl.Options
}

// New it creates a ctl object tool for install/remove services
// service: it needs a service file with .service on suffix or exactly same file as service file name
func New(service string, userMode bool, timeout int) *conf {
	systemdDir := "/lib/systemd/system"
	systemdSoftLinkDir := "/etc/systemd/system/multi-user.target.wants"
	c := &conf{systemdDir: systemdDir, systemdSoftLinkDir: systemdSoftLinkDir, locker: new(sync.Mutex)}
	c.Options = systemctl.Options{UserMode: userMode, Timeout: timeout}
	c.service = service
	return c
}
