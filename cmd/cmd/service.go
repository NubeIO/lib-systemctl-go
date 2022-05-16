package cmd

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/ctl"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/spf13/cobra"
)

var (
	isActive   bool
	status     bool
	start      bool
	stop       bool
	enable     bool
	disable    bool
	add        bool
	remove     bool
	fullRemove bool //stop, disable and remove the file
	path       string
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "service",
	Long:  ``,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {

	timeOut := 5

	ctl.New(&ctl.Options{WorkDir: ""})

	if isActive {
		out, msg, err := systemctl.IsActive(serviceName, systemctl.Options{Timeout: timeOut})
		fmt.Println(out, msg)
		fmt.Println(err)
	}

	if status {
		out, err := systemctl.Status(serviceName, systemctl.Options{Timeout: timeOut})
		fmt.Println(out)
		fmt.Println(err)
	}

	if start {
		err := systemctl.Start(serviceName, systemctl.Options{Timeout: timeOut})
		fmt.Println(err)
	}

	if stop {
		err := systemctl.Stop(serviceName, systemctl.Options{Timeout: timeOut})
		fmt.Println(err)
	}

	if enable {
		err := systemctl.Enable(serviceName, systemctl.Options{Timeout: timeOut})
		fmt.Println(err)
	}

	if disable {
		err := systemctl.Disable(serviceName, systemctl.Options{Timeout: timeOut})
		fmt.Println(err)
	}

	if add {
		ctl.Add(path)
	}
	if remove {
		fmt.Println("try and remove a file:", serviceName)
		err := ctl.Remove(ctl.RemoveOpts{ServiceName: serviceName})
		fmt.Println(err)
	}

}

func init() {
	RootCmd.AddCommand(serviceCmd)
	serviceCmd.Flags().BoolVarP(&isActive, "active", "", false, "if service is active")
	serviceCmd.Flags().BoolVarP(&status, "status", "", false, "status of a service")
	serviceCmd.Flags().BoolVarP(&start, "start", "", false, "start a service")
	serviceCmd.Flags().BoolVarP(&stop, "stop", "", false, "stop a service")
	serviceCmd.Flags().BoolVarP(&enable, "enable", "", false, "enable a service")
	serviceCmd.Flags().BoolVarP(&disable, "disable", "", false, "disable a service")

	serviceCmd.Flags().BoolVarP(&add, "add", "", false, "add a new service")
	serviceCmd.Flags().StringVarP(&path, "path", "", "", "provide the path of the new service file eg: /tmp/rubix-updater.service")

	serviceCmd.Flags().BoolVarP(&remove, "remove", "", false, "remove a new service")
}
