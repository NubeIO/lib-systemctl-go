package ctl

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)

type RemoveOpts struct {
	ServiceName string
	FullRemove  bool
	Stop        bool
	Disable     bool
	TestMode    bool
	Options     systemctl.Options
}

func Remove(service RemoveOpts) error {
	if service.Stop || service.FullRemove {
		err := systemctl.Stop(service.ServiceName, service.Options)
		if err != nil {
			return err
		}
	}
	if service.Disable || service.FullRemove {
		err := systemctl.Disable(service.ServiceName, service.Options)
		if err != nil {
			return err
		}
	}
	if service.FullRemove {
		err := systemctl.DaemonReload(service.Options)
		if err != nil {
			return err
		}
		err = systemctl.RestartFailed(service.Options)
		if err != nil {
			return err
		}
	}

	err := C.removeLib(service)
	if err != nil {
		return err
	}
	err = C.removeUsrLib(service)
	if err != nil {
		return err
	}
	return nil

}

//removeLib service from /lib/system
func (c *conf) removeLib(service RemoveOpts) error {
	c.locker.Lock()
	defer c.locker.Unlock()
	name := service.ServiceName
	name = strings.TrimSuffix(name, ".service")
	svc := c.Has(name)
	if svc == nil {
		return errors.New(fmt.Sprintf("remove file no service with that name exists filename:%s", name))
	}
	err := os.Remove(svc.File)
	if err != nil {
		return errors.New(fmt.Sprintf("remove file error err: %t filename:%s", err, name))
	} else {
		log.Infoln("removed file filename:", svc.File)
	}
	return nil
}

//removeUsrLib service from /lib/system
func (c *conf) removeUsrLib(service RemoveOpts) error {
	c.locker.Lock()
	defer c.locker.Unlock()
	name := service.ServiceName
	log.Println("remove", path.Join(serviceDir, newService(name, "").FullName()))
	err := os.Remove(path.Join(serviceDir, newService(name, "").FullName()))
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to remove %s: %s \n ", name, err.Error()))
	} else {
		log.Println("removed file ok", path.Join(serviceDir, newService(name, "").FullName()))

	}
	return nil

}
