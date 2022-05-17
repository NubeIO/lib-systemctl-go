package builder

import (
	"testing"
)

func TestSystemDBuilder(*testing.T) {
	description := "BIOS comes with default OS, non-upgradable"
	user := "root"
	directory := "/data/rubix-bios-app/v1.5.2"
	execCmd := "./data/rubix-bios -p 1615 -g /data/rubix-bios -d data -c config -a apps --prod --auth  --device-type amd64 --token 1234"
	bld := &SystemDBuilder{
		Description:      description,
		User:             user,
		WorkingDirectory: directory,
		ExecStart:        execCmd,
		SyslogIdentifier: "rubix-bios",
		WriteFile: WriteFile{
			Write:    true,
			FileName: "nubeio-rubix-bios",
			Path:     "/tmp",
		},
	}

	bld.Build()

}

func TestFullInstall(*testing.T) {

	//}

}
