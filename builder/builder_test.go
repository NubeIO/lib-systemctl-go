package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSystemDBuilder(t *testing.T) {
	description := "BIOS comes with default OS, non-upgradable"
	user := "root"
	directory := "/data/rubix-bios-app/v1.5.2"
	execCmd := "./data/rubix-bios -p 1615 -g /data/rubix-bios -d data -c config -a apps --prod --auth  --device-type amd64 --token 1234"
	bld := &SystemDBuilder{
		ServiceName:      "rubix-bios",
		Description:      description,
		User:             user,
		WorkingDirectory: directory,
		ExecStart:        execCmd,
		SyslogIdentifier: "rubix-bios",
		WriteFile: WriteFile{
			Write:    false,
			FileName: "nubeio-rubix-bios",
			Path:     "/tmp",
		},
	}

	err := bld.Build(0700)
	assert.Nil(t, err)
}
