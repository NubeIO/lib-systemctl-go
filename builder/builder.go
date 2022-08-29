package builder

import (
	"errors"
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

/*
[Unit]
Description=BIOS comes with default OS, non-upgradable
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=<working_dir>
ExecStart=<working_dir>/rubix-bios -p 1615 -g /data/rubix-bios -d data -c config -a apps --device-type <device_type> --prod --auth
Restart=always
RestartSec=10
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=rubix-bios

[Install]
WantedBy=multi-user.target
*/

var (
	fileUtils = fileutils.New()
)

func (inst *SystemDBuilder) template() string {
	out := `[Unit]
Description=%v
After=%v
[Service]
ExecStartPre=%v
Type=%v
User=%v
WorkingDirectory=%v
ExecStart=%v
Restart=%v
RestartSec=%v
StandardOutput=%v
StandardError=%v
SyslogIdentifier=%v
[Install]
WantedBy=%v`

	return out
}

type SystemDBuilder struct {
	ServiceName string `json:"service_name"` // nubeio-rubix-bios

	// [Unit]
	Description string `json:"description"`
	After       string `json:"after"` // network.target
	// [Service]
	ExecStartPre     string `json:"exec_start_pre"` // ExecStartPre=/bin/sleep 0
	Type             string `json:"type"`           // simple
	User             string `json:"user"`
	WorkingDirectory string `json:"working_directory"`
	ExecStart        string `json:"exec_start"`
	Restart          string `json:"restart"`
	RestartSec       int    `json:"restart_sec"`
	StandardOutput   string `json:"standard_output"`
	StandardError    string `json:"standard_error"`
	SyslogIdentifier string `json:"syslog_identifier"`
	// [Install]
	WantedBy string `json:"wanted_by"` // multi-user.target

	// write the file to a location
	WriteFile WriteFile `json:"write_file"`
}

type WriteFile struct {
	Write    bool   `json:"write"`
	Path     string `json:"path"`
	FileName string `json:"file_name"` // nubeio-rubix-bios NOT nubeio-rubix-bios.service
}

func checkPath(path string) error {
	if strings.Contains(path, "//") {
		return fmt.Errorf("path is formed incorrect, path cant have // in the path:%s", path)
	}
	return nil
}

func (inst *SystemDBuilder) Build(permissions os.FileMode) error {
	err := checkPath(inst.WorkingDirectory)
	if err != nil {
		return err
	}
	if inst.ServiceName == "" {
		return errors.New("systemctl service builder please provide a ServiceName")
	}
	if inst.WriteFile.Write {
		path := inst.WriteFile.Path
		name := inst.WriteFile.FileName
		if name == "" {
			return errors.New("service file name can not be nil")
		}
		err := fileUtils.DirExistsErr(path)
		if err != nil {
			return err
		}
	}

	if inst.Description == "" {
		inst.Description = fmt.Sprintf("nube-io app %s", inst.ServiceName)
	}
	if inst.After == "" {
		inst.After = "network.target"
	}
	if inst.User == "" {
		inst.User = "root"
	}
	if inst.ExecStartPre == "" {
		inst.ExecStartPre = "/bin/sleep 0"
	}
	if inst.Type == "" {
		inst.Type = "simple"
	}
	if inst.Restart == "" {
		inst.Restart = "always"
	}
	if inst.RestartSec == 0 {
		inst.RestartSec = 10
	}
	if inst.StandardOutput == "" {
		inst.StandardOutput = "syslog"
	}
	if inst.StandardError == "" {
		inst.StandardError = "syslog"
	}
	if inst.SyslogIdentifier == "" {
		inst.SyslogIdentifier = strings.ToLower(inst.ServiceName)
	}
	if inst.WantedBy == "" {
		inst.WantedBy = "multi-user.target"
	}

	serviceFile := fmt.Sprintf(inst.template(),
		inst.Description,
		inst.After,
		inst.ExecStartPre,
		inst.Type,
		inst.User,
		inst.WorkingDirectory,
		inst.ExecStart,
		inst.Restart,
		inst.RestartSec,
		inst.StandardOutput,
		inst.StandardError,
		inst.SyslogIdentifier,
		inst.WantedBy,
	)
	log.Infoln("-----------build service file--------------")
	log.Infoln(serviceFile)
	log.Infoln("-----------build service file--------------")
	if inst.WriteFile.Write {
		path := inst.WriteFile.Path
		name := inst.WriteFile.FileName
		servicePath := fmt.Sprintf("%s/%v.service", path, name)
		log.Infoln("build and add new file here:", servicePath)
		err = fileUtils.WriteFile(servicePath, serviceFile, permissions)
		if err != nil {
			log.Errorf("write service file error %s", err.Error())
			return err
		}
	}
	return nil
}
