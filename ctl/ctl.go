package ctl

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type conf struct {
	service     string
	path        string
	services    []*service // list of managed services
	controlDir  string
	workDir     string
	locker      *sync.Mutex
	InstallOpts InstallOpts
	RemoveOpts  RemoveOpts
}

func New(service, path string) *conf {
	c := newConf()
	c.service = service
	c.path = path
	return c
}

// read from local home directory
func newConf() *conf {
	dir := "/lib/systemd/system"
	c := conf{workDir: dir, controlDir: dir, locker: new(sync.Mutex)}

	stat, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) { //if not exists make dir
			//if err := os.Mkdir(dir, 0755); err != nil {
			//	fmt.Println(" Failed to initialize configuration:", err.Error())
			//	os.Exit(1)
			//}
			return &c
		}
		fmt.Printf(" Failed to load local configuration: %s\n", err.Error())
		os.Exit(1)
	}
	if !stat.IsDir() {
		fmt.Println(" Configuration file is invalid ")
		os.Exit(1)
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf(" Failed to load local configuration: %s\n ", err.Error())
		os.Exit(1)
	}
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".service" {
			c.services = append(c.services, newService(file.Name(), path.Join(dir, file.Name())))
		}
	}
	return &c
}

//Has Whether the service is already managed
func (inst *conf) Has(name string) *service {
	name = strings.TrimSuffix(name, ".service")
	for _, existed := range inst.services {
		if name == existed.Name {
			return existed
		}
	}
	return nil
}

//List  of managed services
func (inst *conf) List() []*service {
	return inst.services
}

// WorkDir get
func (inst *conf) WorkDir() string {
	return inst.workDir
}

// ControlDir get
func (inst *conf) ControlDir() string {
	return inst.controlDir
}
