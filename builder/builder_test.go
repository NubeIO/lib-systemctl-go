package builder

import (
	"testing"
)

func TestSystemDBuilder(*testing.T) {
	name := "aidans-service"
	user := "aidan"
	directory := "/home/aidan"
	execCmd := "/usr/bin/python3 something.py"
	bld := &SystemDBuilder{
		Name:      name,
		User:      user,
		Directory: directory,
		ExecCmd:   execCmd,
	}

	bld.build()

}
