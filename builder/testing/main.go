package main

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/builder"
	"github.com/NubeIO/lib-systemctl-go/ctl"
	pprint "github.com/NubeIO/lib-systemctl-go/helpers/print"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	newService := "nubeio-rubix-bios"
	description := "BIOS comes with default OS, non-upgradable"
	user := "root"
	directory := "/data/rubix-bios-app"
	execCmd := "/data/rubix-bios-app/rubix-bios -p 1615 -g /data/rubix-bios -d data -c config -a apps --prod --auth  --device-type amd64 --token 1234"
	bld := &builder.SystemDBuilder{
		Description:      description,
		User:             user,
		WorkingDirectory: directory,
		ExecStart:        execCmd,
		SyslogIdentifier: "rubix-bios",
		WriteFile: builder.WriteFile{
			Write:    true,
			FileName: newService,
			Path:     "/tmp",
		},
	}

	//generate the service file
	//check if service file exists and replace if it does
	//and if the service does exist then restart the service or install it

	err := bld.Build()
	if err != nil {
		fmt.Println(err, "err in build add new file")
		//return
	}

	path := "/tmp/nubeio-rubix-bios.service"

	timeOut := 30
	service := ctl.New(newService, path)
	opts := systemctl.Options{Timeout: timeOut}
	installOpts := ctl.InstallOpts{
		Options: opts,
	}
	service.InstallOpts = installOpts
	err = service.Install()
	fmt.Println("full install error", err)
	if err != nil {
		fmt.Println("full install error", err)
	}

	time.Sleep(8 * time.Second)

	status, err := systemctl.Status(newService, systemctl.Options{})
	if err != nil {
		log.Errorf("service found: %s: %v", newService, err)
	}
	fmt.Println(status)

	res, err := service.Remove()
	fmt.Println("full install error", err)
	if err != nil {
		fmt.Println("full install error", err)
	}
	pprint.PrintJOSN(res)

}
