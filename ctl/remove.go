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
	RemoveOpts systemctl.Options
}

type RemoveRes struct {
	Stop                 bool `json:"stop"`
	Disable              bool `json:"disable"`
	DaemonReload         bool `json:"daemon_reload"`
	RestartFailed        bool `json:"restart_failed"`
	DeleteServiceFile    bool `json:"delete_service_file"`
	DeleteServiceFileUsr bool `json:"delete_service_file_usr"`
}

func (inst *conf) Remove() (*RemoveRes, error) {
	res := &RemoveRes{}
	err := systemctl.Stop(inst.service, inst.RemoveOpts.RemoveOpts)
	if err != nil {
		log.Errorln("failed to stop:", inst.service)
	} else {
		res.Stop = true
	}
	err = systemctl.Disable(inst.service, inst.RemoveOpts.RemoveOpts)
	if err != nil {
		log.Errorln("failed to disable:", inst.service)
	} else {
		res.Disable = true
	}
	err = systemctl.DaemonReload(inst.RemoveOpts.RemoveOpts)
	if err != nil {
		log.Errorln("failed to reload-demon:", inst.service)
	} else {
		res.DaemonReload = true
	}
	err = systemctl.RestartFailed(inst.RemoveOpts.RemoveOpts)
	if err != nil {
		log.Errorln("failed to restart-failed:", inst.service)
	} else {
		res.RestartFailed = true
	}
	//remove service file from /lib/system
	err = inst.removeLib()
	if err != nil {
		log.Errorln("failed to delete-file /lib/systemd/system/", inst.service)
	} else {
		res.DeleteServiceFile = true
	}
	//remove service file from /usr/lib/system
	err = inst.removeUsrLib()
	if err != nil {
		log.Errorln("failed to delete-file usr/lib/systemd/system/", inst.service)
	} else {
		res.DeleteServiceFileUsr = true
	}
	return res, nil

}

//removeLib service from /lib/system
func (inst *conf) removeLib() error {
	inst.locker.Lock()
	defer inst.locker.Unlock()
	name := inst.service
	name = strings.TrimSuffix(name, ".service")
	svc := inst.Has(name)
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
func (inst *conf) removeUsrLib() error {
	inst.locker.Lock()
	defer inst.locker.Unlock()
	name := inst.service
	log.Println("remove", path.Join(serviceDir, newService(name, "").FullName()))
	err := os.Remove(path.Join(serviceDir, newService(name, "").FullName()))
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to remove %s: %s \n ", name, err.Error()))
	} else {
		log.Println("removed file ok", path.Join(serviceDir, newService(name, "").FullName()))

	}
	return nil

}
