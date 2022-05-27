package builder

import (
	"errors"
	"fmt"
	"io/ioutil"
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

func (inst *SystemDBuilder) template() string {
	out := `[Unit]
Description=%v
After=%v
[Service]
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
	ServiceName string `json:"service_name"` //nubeio-rubix-bios

	//[Unit]
	Description string `json:"description"`
	After       string `json:"after"` //network.target
	//[Service]
	Type             string `json:"type"` //simple
	User             string `json:"user"`
	WorkingDirectory string `json:"working_directory"`
	ExecStart        string `json:"exec_start"`
	Restart          string `json:"restart"`
	RestartSec       int    `json:"restart_sec"`
	StandardOutput   string `json:"standard_output"`
	StandardError    string `json:"standard_error"`
	SyslogIdentifier string `json:"syslog_identifier"`
	//[Install]
	WantedBy string `json:"wanted_by"` //multi-user.target

	// write the file to a location
	WriteFile WriteFile `json:"write_file"`
}

type WriteFile struct {
	Write    bool   `json:"write"`
	Path     string `json:"path"`
	FileName string `json:"file_name"` //nubeio-rubix-bios NOT nubeio-rubix-bios.service
}

func (inst *SystemDBuilder) Build() error {
	if inst.Description == "" {
		return errors.New("please provide a description")
	}
	if inst.After == "" {
		inst.After = "network.target"
	}
	if inst.User == "" {
		inst.User = "root"
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
		inst.SyslogIdentifier = strings.ToLower(inst.Description)
	}
	if inst.WantedBy == "" {
		inst.WantedBy = "multi-user.target"
	}

	serviceFile := fmt.Sprintf(inst.template(),
		inst.Description,
		inst.After,
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
	fmt.Println("------------------------------")
	fmt.Println(serviceFile)
	fmt.Println("------------------------------")
	if inst.WriteFile.Write {
		path := inst.WriteFile.Path
		name := inst.WriteFile.FileName
		servicePath := fmt.Sprintf("%s/%v.service", path, name)
		fmt.Println("------------------------------")
		fmt.Println("build and add new file here:", servicePath)
		fmt.Println("------------------------------")
		err := ioutil.WriteFile(servicePath, []byte(serviceFile), os.ModePerm)
		if err != nil {
			fmt.Println("write file error", err)
			return err
		}
	}
	return nil

}
