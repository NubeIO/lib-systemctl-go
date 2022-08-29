package cmd

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/builder"
	"github.com/spf13/cobra"
)

var (
	write       bool
	servicePath string
	serviceName string
)

var builderCmd = &cobra.Command{
	Use:   "builder",
	Short: "builder",
	Long:  ``,
	Run:   runBuilder,
}

func runBuilder(cmd *cobra.Command, args []string) {
	description := "Service Description"
	user := "aidan"
	directory := "/home/aidan"
	execCmd := "/usr/bin/python3 something.py"

	writeFile := builder.WriteFile{
		Write:    write,
		Path:     path,
		FileName: serviceName,
	}

	builderInstance := &builder.SystemDBuilder{
		ServiceName:      serviceName,
		Description:      description,
		User:             user,
		WorkingDirectory: directory,
		ExecStart:        execCmd,
		SyslogIdentifier: "rubix-bios",
		WriteFile:        writeFile,
	}
	err := builderInstance.Build(0700)
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	RootCmd.AddCommand(builderCmd)
	builderCmd.PersistentFlags().BoolVarP(&write, "write", "", false, "generate a new systemd file")
	builderCmd.PersistentFlags().StringVarP(&servicePath, "service-path", "", "/tmp", "provide the path of the new service file eg: /tmp/")
	builderCmd.PersistentFlags().StringVarP(&serviceName, "service-name", "", "rubix-bios", "rubix-updater")
}
