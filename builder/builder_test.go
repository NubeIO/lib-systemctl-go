package builder

import (
	"testing"
)

func TestSystemDBuilder(*testing.T) {
	description := "BIOS comes with default OS, non-upgradable"
	user := "root"
	directory := "/data/tmp/rubix-bios"
	execCmd := "./data/rubix-bios -p 1615 -g /data/rubix-bios -d data -c config -a apps --prod --auth  --device-type amd64 --token 1234"
	bld := &SystemDBuilder{
		Description:      description,
		User:             user,
		WorkingDirectory: directory,
		ExecStart:        execCmd,
		SyslogIdentifier: "rubix-bios",
	}

	bld.Build()

}
