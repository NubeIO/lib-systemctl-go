package systemctl

import (
	"context"
	"time"
)

// RestartFailed to remove the failed status. To reset all units with failed status:
func (inst *SystemCtl) RestartFailed() error {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(inst.Timeout)*time.Second)
	defer cancel()
	var args = []string{"reset-failed", "--system"}
	if inst.UserMode {
		args[1] = "--user"
	}
	_, _, _, err := execute(ctx, args)
	return err
}

// DaemonReload Reload systemd manager configuration.
//
// This will rerun all generators (see systemd. generator(7)), reload all unit
// files, and recreate the entire dependency tree. While the daemon is being
// reloaded, all sockets systemd listens on behalf of user configuration will
// stay accessible.
func (inst *SystemCtl) DaemonReload() error {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(inst.Timeout)*time.Second)
	defer cancel()
	var args = []string{"daemon-reload", "--system"}
	if inst.UserMode {
		args[1] = "--user"
	}
	_, _, _, err := execute(ctx, args)
	return err
}

// Restart Stop and then start one or more units specified on the command line.
// If the units are not running yet, they will be started.
func (inst *SystemCtl) Restart(unit string) error {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(inst.Timeout)*time.Second)
	defer cancel()
	var args = []string{"restart", "--system", unit}
	if inst.UserMode {
		args[1] = "--user"
	}
	_, _, _, err := execute(ctx, args)
	return err
}
