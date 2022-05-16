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

//var C = newConf()
var C *conf

type conf struct {
	services   []*service // list of managed services
	controlDir string
	workDir    string
	locker     *sync.Mutex
}

type Options struct {
	WorkDir string
}

func New(opts *Options) {

	C = newConf()

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
func (c *conf) Has(name string) *service {
	name = strings.TrimSuffix(name, ".service")
	for _, existed := range c.services {
		if name == existed.Name {
			return existed
		}
	}
	return nil
}

//List  of managed services
func (c *conf) List() []*service {
	return c.services
}

// WorkDir get
func (c *conf) WorkDir() string {
	return c.workDir
}

// ControlDir get
func (c *conf) ControlDir() string {
	return c.controlDir
}
