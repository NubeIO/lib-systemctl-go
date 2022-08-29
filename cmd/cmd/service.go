package cmd

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/ctl"

	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/spf13/cobra"
)

var (
	isActive    bool
	status      bool
	start       bool
	restart     bool
	stop        bool
	enable      bool
	disable     bool
	add         bool
	remove      bool
	fullRemove  bool // stop, disable and remove the file
	fullInstall bool // add the new file start, enable
	path        string
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "service",
	Long:  ``,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {

	timeOut := 5

	service := ctl.New(serviceName, path)
	opts := systemctl.Options{Timeout: timeOut}
	installOpts := ctl.InstallOpts{
		Options: opts,
	}
	removeOpts := ctl.RemoveOpts{RemoveOpts: opts}
	service.InstallOpts = installOpts
	service.RemoveOpts = removeOpts

	// if isActive {
	// 	out, msg, err := systemctl.IsActive(serviceName, opts)
	// 	fmt.Println(out, msg)
	// 	fmt.Println(err)
	// }
	//
	// if status {
	// 	out, err := systemctl.Status(serviceName, opts)
	// 	fmt.Println(out)
	// 	fmt.Println(err)
	// }
	//
	// if start {
	// 	err := systemctl.Start(serviceName, opts)
	// 	fmt.Println(err)
	// }
	//
	// if restart {
	// 	err := systemctl.Restart(serviceName, opts)
	// 	fmt.Println(err)
	// }
	//
	// if stop {
	// 	err := systemctl.Stop(serviceName, opts)
	// 	fmt.Println(err)
	// }
	//
	// if enable {
	// 	err := systemctl.Enable(serviceName, opts)
	// 	fmt.Println(err)
	// }
	//
	// if disable {
	// 	err := systemctl.Disable(serviceName, opts)
	// 	fmt.Println(err)
	// }

	if add {
		err := service.Add(path)
		if err != nil {
			fmt.Println("full Add error", err)
		}
	}

	if fullInstall {
		err := service.Install()
		if err != nil {
			fmt.Println("full install error", err)
		}
	}

	if remove {
		fmt.Println("try and remove a file:", serviceName)
		// err := service.Remove()
		// fmt.Println(err)
	}
	if fullRemove {
		fmt.Println("try and remove a file:", serviceName)
		// err := service.Remove()
		// fmt.Println(err)
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

	serviceCmd.Flags().BoolVarP(&add, "add", "", false, "add a new service file")
	serviceCmd.Flags().BoolVarP(&fullInstall, "install", "", false, "add a new service file and do the full install")

	serviceCmd.Flags().StringVarP(&path, "path", "", "", "provide the path of the new service file eg: /tmp/rubix-updater.service")

	serviceCmd.Flags().BoolVarP(&remove, "remove", "", false, "remove a new service")
	serviceCmd.Flags().BoolVarP(&fullRemove, "remove-force", "", false, "remove a the service, actions are stop, disable, delete the files and daemon-reload")
}
