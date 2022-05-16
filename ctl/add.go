package ctl

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

//Add a new service
func Add(service string) {
	if err := C.add(service); err != nil {
		fmt.Printf("Failed to add %s: %s \n ", service, err.Error())
	}
}

//Add  service hosting
func (c *conf) add(file string) error {
	c.locker.Lock()
	defer c.locker.Unlock()
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
	if c.Has(stat.Name()) != nil {
		return fmt.Errorf("%s already exists", stat.Name())
	}
	expected := path.Join(c.workDir, stat.Name())
	err = copyFile(file, expected)
	if err != nil {
		return err
	}
	c.services = append(c.services, newService(stat.Name(), expected))
	return nil
}

//copyFile copy the file
func copyFile(src, dst string) error {
	var buf = make([]byte, 5*2^20)
	stat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !stat.Mode().IsRegular() {
		return fmt.Errorf("invalid file: %s", src)
	}
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(source *os.File) {
		_ = source.Close()
	}(source)
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(destination *os.File) {
		_ = destination.Close()
	}(destination)
	for {
		Byte, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if Byte == 0 {
			break
		}
		_, err = destination.Write(buf[:Byte])
		if err != nil {
			return err
		}
	}
	return nil
}
