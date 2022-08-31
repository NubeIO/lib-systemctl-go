package systemd

import (
	"path"
	"syscall"
)

func (inst *conf) SoftLink() error {
	actualFile := path.Join(inst.systemdDir, inst.service)
	linkFile := path.Join(inst.systemdSoftLinkDir, inst.service)
	return syscall.Symlink(actualFile, linkFile)
}

func (inst *conf) SoftUnlink() error {
	file := path.Join(inst.systemdSoftLinkDir, inst.service)
	return syscall.Unlink(file)
}
