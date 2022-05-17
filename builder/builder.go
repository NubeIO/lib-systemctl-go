package builder

import (
	"errors"
	"fmt"
	"io/ioutil"
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

	//[Unit]
	Description string
	After       string //network.target
	//[Service]
	Type             string //simple
	User             string
	WorkingDirectory string
	ExecStart        string
	Restart          string
	RestartSec       int
	StandardOutput   string
	StandardError    string
	SyslogIdentifier string
	//[Install]
	WantedBy string //multi-user.target

	// write the file to a location
	WriteFile WriteFile
}

type WriteFile struct {
	Write    bool
	Path     string
	FileName string
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
		err := ioutil.WriteFile(servicePath, []byte(serviceFile), 0644)
		if err != nil {
			return err
		}
	}
	return nil

}
