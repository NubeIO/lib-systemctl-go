package builder

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSystemDBuilder(t *testing.T) {
	description := "Rubix Edge BIOS comes with default OS, non-upgradable"
	user := "root"
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	execCmd := fmt.Sprintf("%s/rubix-edge-bios server -p 1659 -r /data -a rubix-edge-bios -d data -c config -a apps --prod --auth", wd)
	systemdBuilder := &SystemDBuilder{
		ServiceName:      "rubix-bios",
		Description:      description,
		User:             user,
		WorkingDirectory: wd,
		ExecStart:        execCmd,
		SyslogIdentifier: "rubix-edge-bios",
		WriteFile: WriteFile{
			Write:    false,
			FileName: "nubeio-rubix-edge-bios",
			Path:     "/tmp",
		},
	}

	err = systemdBuilder.Build(0644)
	assert.Nil(t, err)
}
