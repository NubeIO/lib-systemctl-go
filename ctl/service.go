package ctl

import (
	"strings"
)

const (
	serviceDir = "/usr/lib/systemd/system/"
)

type service struct {
	Name         string
	File         string
	statusOutput string // save to reduce execution times
}

func newService(name, file string) *service {
	svc := service{
		Name: strings.TrimSuffix(name, ".service"),
		File: file,
	}
	return &svc
}

func (s *service) FullName() string {
	if !strings.HasSuffix(s.Name, ".service") {
		return s.Name + ".service"
	}
	return s.Name
}
