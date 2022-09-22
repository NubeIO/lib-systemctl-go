package systemd

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	log "github.com/sirupsen/logrus"
	"path"
	"path/filepath"
)

type InstallOpts struct {
	Options systemctl.Options
}

type InstallResponse struct {
	UnlinkServiceFile bool `json:"unlink_service_file"`
	LinkServiceFile   bool `json:"link_service_file"`
	DaemonReload      bool `json:"daemon_reload"`
	Enable            bool `json:"enabled"`
	Restart           bool `json:"restarted"`
}

// TransferSystemdFile a new service
func (inst *Systemd) TransferSystemdFile(sourceFile string) error {
	inst.locker.Lock()
	defer inst.locker.Unlock()
	if filepath.Ext(sourceFile) != ".service" {
		return fmt.Errorf("must add a valid service file, your current file is: %s", sourceFile)
	}
	hasFile := fileutils.FileExists(sourceFile)
	if hasFile == false {
		return errors.New(fmt.Sprintf("file doesn't exist: %s", sourceFile))
	}
	destinationFile := path.Join(inst.systemdDir, inst.service)
	err := fileutils.CopyFile(sourceFile, destinationFile)
	if err != nil {
		return err
	}
	return nil
}

// Install a new service
func (inst *Systemd) Install() (resp *InstallResponse) {
	resp = &InstallResponse{}
	ctl := systemctl.New(inst.Options.UserMode, inst.Options.Timeout)

	log.Info(fmt.Sprintf("soft un-linking linux service: %s...", inst.service))
	err := inst.SoftUnlink()
	if err != nil {
		log.Errorln(err)
	} else {
		resp.UnlinkServiceFile = true
	}

	log.Info(fmt.Sprintf("soft linking linux service: %s...", inst.service))
	err = inst.SoftLink()
	if err != nil {
		log.Errorln(err)
	} else {
		resp.LinkServiceFile = true
	}

	log.Info("hitting daemon-reload...")
	err = ctl.DaemonReload()
	if err != nil {
		log.Errorln(err)
		return
	} else {
		resp.DaemonReload = true
	}

	log.Info(fmt.Sprintf("enabling linux service: %s...", inst.service))
	err = ctl.Enable(inst.service)
	if err != nil {
		log.Errorln(err)
		return
	} else {
		resp.Enable = true
	}

	log.Info(fmt.Sprintf("restarting linux service: %s...", inst.service))
	err = ctl.Restart(inst.service)
	if err != nil {
		log.Errorln(err)
		return
	} else {
		resp.Restart = true
	}
	return resp
}
