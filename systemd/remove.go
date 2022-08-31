package systemd

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
)

type RemoveRes struct {
	ServiceWasInstalled bool `json:"service_was_installed"`
	Stop                bool `json:"stop"`
	DaemonReload        bool `json:"daemon_reload"`
	UnlinkServiceFile   bool `json:"unlink_service_file"`
	DeleteServiceFile   bool `json:"delete_service_file"`
}

func (inst *conf) Remove() (res *RemoveRes) {
	res = &RemoveRes{}
	ctl := systemctl.New(inst.Options.UserMode, inst.Options.Timeout)

	wasInstalled, err := ctl.IsInstalled(inst.service)
	res.ServiceWasInstalled = wasInstalled

	log.Info(fmt.Sprintf("stopping linux service: %s...", inst.service))
	err = ctl.Stop(inst.service)
	if err != nil {
		log.Errorln(err)
	} else {
		res.Stop = true
	}

	log.Info(fmt.Sprintf("un-linking linux service: %s...", inst.service))
	err = inst.SoftUnlink()
	if err != nil {
		log.Errorln(err)
	} else {
		res.UnlinkServiceFile = true
	}

	log.Info(fmt.Sprintf("removing linux service: %s...", inst.service))
	err = inst.RemoveLib()
	if err != nil {
		log.Errorln(err)
	} else {
		res.DeleteServiceFile = true
	}

	log.Info("hitting daemon-reload...")
	err = ctl.DaemonReload()
	if err != nil {
		log.Errorln(err)
	} else {
		res.DaemonReload = true
	}
	return
}

func (inst *conf) RemoveLib() error {
	file := path.Join(inst.systemdDir, inst.service)
	hasFile := fileutils.FileExists(file)
	if hasFile == false {
		return errors.New(fmt.Sprintf("file doesn't exist: %s", file))
	}
	err := os.Remove(file)
	if err != nil {
		return errors.New(fmt.Sprintf("remove file error err: %t on file: %s", err, file))
	}
	log.Infoln("removed file filename:", file)
	return nil
}
