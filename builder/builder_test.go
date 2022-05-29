package builder

import (
	"fmt"
	"strings"
	"testing"
)

func TestSystemDBuilder(*testing.T) {
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
	fmt.Println(err)
	if err != nil {
		return
	}

	fmt.Println(strings.Contains("/home/aidan/aidan-test/aidan_test//", "//"))

	//base := "/home"
	//path := "aa"
	//rel, err := filepath.Rel(base, path)
	//fmt.Printf("Base %q: Path %q: Rel %q Err %v\n", base, path, rel, err)
	//if err != nil {
	//	fmt.Println("PROCEED")
	//	return
	//}

}

func TestFullInstall(*testing.T) {

	//}

}
