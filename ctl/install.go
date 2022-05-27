package ctl

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"path/filepath"
)

type InstallOpts struct {
	Options systemctl.Options
}

type InstallResp struct {
	Install      string
	DaemonReload string
	Enable       string
	Restart      string
}

//Install a new service
func (inst *conf) Install() *InstallResp {
	resp := &InstallResp{}
	if err := inst.add(inst.path); err != nil {
		log.Errorf("failed to add %s: %s \n ", inst.path, err.Error())
		resp.Install = err.Error()
		return resp
	}
	log.Infof("added new file %s: \n ", inst.path)
	//reload
	err := systemctl.DaemonReload(inst.InstallOpts.Options)
	if err != nil {
		log.Errorf("failed to DaemonReload%s: err:%s \n ", inst.service, err.Error())
		resp.DaemonReload = err.Error()
		return resp
	}
	//enable
	err = systemctl.Enable(inst.service, inst.InstallOpts.Options)
	if err != nil {
		log.Errorf("failed to enable%s: err:%s \n ", inst.service, err.Error())
		resp.Enable = err.Error()
		return resp
	}
	log.Infof("enable new service:%s \n ", inst.service)
	//start
	err = systemctl.Restart(inst.service, inst.InstallOpts.Options)
	if err != nil {
		log.Errorf("failed to start%s: err:%s \n ", inst.service, err.Error())
		resp.Restart = err.Error()
		return resp
	}
	log.Infof("start new service:%s \n ", inst.service)
	return nil
}

//Add a new service
func (inst *conf) Add(path string) error {
	if err := inst.add(path); err != nil {
		return err
	}
	return nil
}

//Add  service hosting
func (inst *conf) add(file string) error {
	inst.locker.Lock()
	defer inst.locker.Unlock()
	if filepath.Ext(file) != ".service" {
		return fmt.Errorf(" must add a valid service file")
	}
	stat, err := os.Stat(file)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return fmt.Errorf("must add a valid service file")
	}

	replaceFile := false
	if replaceFile { //TODO maybe give the user this option
		if inst.Has(stat.Name()) != nil {
			return fmt.Errorf("%s already exists", stat.Name())
		}

	}
	expected := path.Join(inst.workDir, stat.Name())
	err = copyFile(file, expected)
	if err != nil {
		fmt.Println("copyfile", err)
		return err
	}
	inst.services = append(inst.services, newService(stat.Name(), expected))
	return nil
}

//copyFile copy the file
func copyFile(src, dst string) error {
	var buf = make([]byte, 5*2^20)
	stat, err := os.Stat(src)
	if err != nil {
		fmt.Println("STAT err")
		return err
	}
	if !stat.Mode().IsRegular() {
		return fmt.Errorf("invalid file: %s", src)
	}
	source, err := os.Open(src)
	if err != nil {
		fmt.Println("OPne err")
		return err
	}
	defer func(source *os.File) {
		_ = source.Close()
	}(source)
	destination, err := os.Create(dst)
	if err != nil {
		fmt.Println("Create err")
		return err
	}
	defer func(destination *os.File) {
		_ = destination.Close()
	}(destination)
	for {
		Byte, err := source.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read err")
			return err
		}
		if Byte == 0 {
			break
		}
		_, err = destination.Write(buf[:Byte])
		if err != nil {
			fmt.Println("write err")
			return err
		}
	}
	return nil
}
