package cmd

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/builder"
	"github.com/spf13/cobra"
)

var (
	builderPath        string
	builderServiceName string
	builderWrite       bool
)

var builderCmd = &cobra.Command{
	Use:   "builder",
	Short: "builder",
	Long:  ``,
	Run:   runBuilder,
}

func runBuilder(cmd *cobra.Command, args []string) {

	name := "aidans-service"
	user := "aidan"
	directory := "/home/aidan"
	execCmd := "/usr/bin/python3 something.py"

	write := builder.WriteFile{
		Write:    true,
		Path:     builderPath,
		FileName: builderServiceName,
	}

	bld := &builder.SystemDBuilder{
		Description:      name,
		User:             user,
		WorkingDirectory: directory,
		ExecStart:        execCmd,
		SyslogIdentifier: "rubix-bios",
		WriteFile:        write,
	}

	err := bld.Build(0700)
	if err != nil {
		fmt.Println(err)
	}

}

func init() {
	RootCmd.AddCommand(builderCmd)
	builderCmd.Flags().BoolVarP(&builderWrite, "write", "", false, "generate a new systemd file")
	builderCmd.Flags().StringVarP(&builderPath, "builder-path", "", "", "provide the path of the new service file eg: /tmp/")
	builderCmd.Flags().StringVarP(&builderServiceName, "builder-name", "", "", "rubix-updater")

}
