package systemd

import (
	"path"
	"syscall"
)

func (inst *Systemd) SoftLink() error {
	actualFile := path.Join(inst.systemdDir, inst.service)
	linkFile := path.Join(inst.systemdSoftLinkDir, inst.service)
	return syscall.Symlink(actualFile, linkFile)
}

func (inst *Systemd) SoftUnlink() error {
	file := path.Join(inst.systemdSoftLinkDir, inst.service)
	return syscall.Unlink(file)
}
