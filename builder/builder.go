package builder

import (
	"fmt"
	"io/ioutil"
)

type SystemDBuilder struct {
	Name      string
	Directory string
	ExecCmd   string
	User      string
	WriteFile WriteFile
}

type WriteFile struct {
	Write    bool
	Path     string
	FileName string
}

func (inst *SystemDBuilder) Build() error {
	serviceFile := fmt.Sprintf(inst.template(), inst.Name, inst.User, inst.Directory, inst.ExecCmd)
	fmt.Println("------------------------------")
	fmt.Println(serviceFile)
	fmt.Println("------------------------------")
	if inst.WriteFile.Write {
		path := inst.WriteFile.Path
		name := inst.WriteFile.FileName
		servicePath := fmt.Sprintf("%s/%v.service", path, name)
		fmt.Println("------------------------------")
		fmt.Println("build and add new file here:", servicePath)
		fmt.Println("------------------------------")
		err := ioutil.WriteFile(servicePath, []byte(serviceFile), 0644)
		if err != nil {
			return err
		}
	}
	return nil

}

func (inst *SystemDBuilder) template() string {
	out := `[Unit]
Description=%v Service
After=network.target
[Service]
User=%v
WorkingDirectory=%v
ExecStart=%v
Restart=always
[Install]
WantedBy=multi-user.target`
	return out

}

//func template() {
//
//	name := "aidans-service"
//	user := "aidan"
//	directory := "/home/aidan"
//	execCmd := "/usr/bin/python3 something.py"
//
//	template := `[Unit]
//Description=%v Service
//After=network.target
//[Service]
//User=%v
//WorkingDirectory=%v
//ExecStart=%v
//Restart=always
//[Install]
//WantedBy=multi-user.target`
//
//	serviceFile := fmt.Sprintf(template, name, user, directory, execCmd)
//	//homedir
//	servicePath := fmt.Sprintf("/home/aidan/%v.service", name)
//	ioutil.WriteFile(servicePath, []byte(serviceFile), 0644)
//	fmt.Println("Service File Created.")
//}
