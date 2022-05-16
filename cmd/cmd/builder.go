package cmd

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/builder"
	"github.com/spf13/cobra"
)

var (
	//isActive   bool
	//status     bool
	//start      bool
	//stop       bool
	//enable     bool
	//disable    bool
	//add        bool
	//remove     bool
	//fullRemove bool //stop, disable and remove the file
	//path       string

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
		Name:      name,
		User:      user,
		Directory: directory,
		ExecCmd:   execCmd,
		WriteFile: write,
	}

	err := bld.Build()
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
