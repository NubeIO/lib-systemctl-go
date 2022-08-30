package cmd

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/ctl"

	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/spf13/cobra"
)

var (
	isActive bool
	status   bool
	start    bool
	restart  bool
	stop     bool
	enable   bool
	disable  bool
	add      bool
	remove   bool
	install  bool
	path     string
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "service",
	Long:  ``,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	timeout := 5
	systemctlObject := systemctl.New(false, timeout)
	if isActive {
		out, msg, err := systemctlObject.IsActive(serviceName)
		fmt.Println(out, msg)
		fmt.Println(err)
	}

	if status {
		out, err := systemctlObject.Status(serviceName)
		fmt.Println(out)
		fmt.Println(err)
	}

	if start {
		err := systemctlObject.Start(serviceName)
		fmt.Println(err)
	}

	if restart {
		err := systemctlObject.Restart(serviceName)
		fmt.Println(err)
	}

	if stop {
		err := systemctlObject.Stop(serviceName)
		fmt.Println(err)
	}

	if enable {
		err := systemctlObject.Enable(serviceName)
		fmt.Println(err)
	}

	if disable {
		err := systemctlObject.Disable(serviceName)
		fmt.Println(err)
	}

	service := ctl.New(serviceName, false, timeout)
	if add {
		err := service.Add(path)
		if err != nil {
			fmt.Println("add error", err)
		}
	}

	if install {
		err := service.Install()
		if err != nil {
			fmt.Println("install error", err)
		}
	}

	if remove {
		fmt.Println("try and remove a file:", serviceName)
		service.Remove()
	}
}

func init() {
	RootCmd.AddCommand(serviceCmd)
	serviceCmd.Flags().BoolVarP(&isActive, "active", "", false, "if service is active")
	serviceCmd.Flags().BoolVarP(&status, "status", "", false, "status of a service")
	serviceCmd.Flags().BoolVarP(&start, "start", "", false, "start a service")
	serviceCmd.Flags().BoolVarP(&restart, "restart", "", false, "restart a service")
	serviceCmd.Flags().BoolVarP(&stop, "stop", "", false, "stop a service")
	serviceCmd.Flags().BoolVarP(&enable, "enable", "", false, "enable a service")
	serviceCmd.Flags().BoolVarP(&disable, "disable", "", false, "disable a service")

	serviceCmd.Flags().StringVarP(&path, "path", "", "", "provide the path of the new service file eg: /tmp/rubix-updater.service")
	serviceCmd.Flags().BoolVarP(&add, "add", "", false, "add a new service file")
	serviceCmd.Flags().BoolVarP(&install, "install", "", false, "deamon-reload, enable, start service")
	serviceCmd.Flags().BoolVarP(&remove, "remove", "", false, "remove a new service")
}
