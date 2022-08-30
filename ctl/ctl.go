package ctl

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type conf struct {
	service    string
	services   []*service // list of managed services
	systemdDir string
	locker     *sync.Mutex
	Options    systemctl.Options
}

func New(service string, userMode bool, timeout int) *conf {
	c := newConf()
	c.service = service
	c.Options = systemctl.Options{UserMode: userMode, Timeout: timeout}
	return c
}

// read from local home directory
func newConf() *conf {
	dir := "/lib/systemd/system"
	c := &conf{systemdDir: dir, locker: new(sync.Mutex)}

	stat, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			panic(err.Error())
		}
		fmt.Printf("Failed to load local configuration: %s\n", err.Error())
		return c
	}
	if !stat.IsDir() {
		fmt.Println("Configuration file is invalid")
		return c
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("Failed to load local configuration: %s\n", err.Error())
		return c
	}
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".service" {
			c.services = append(c.services, newService(file.Name(), path.Join(dir, file.Name())))
		}
	}
	return c
}

// Has Whether the service is already managed
func (inst *conf) Has(name string) *service {
	name = strings.TrimSuffix(name, ".service")
	for _, existed := range inst.services {
		if name == existed.Name {
			return existed
		}
	}
	return nil
}

// List  of managed services
func (inst *conf) List() []*service {
	return inst.services
}

// SystemdDir get
func (inst *conf) SystemdDir() string {
	return inst.systemdDir
}
