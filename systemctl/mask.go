package systemctl

import (
	"context"
	"time"
)

// Mask one or more units, as specified on the command line. This will link
// these unit files to /dev/null, making it impossible to start them.
//
// Notably, Mask may return ErrDoesNotExist if a unit doesn't exist, but it will
// continue masking anyway. Calling Mask on a non-existing masked unit does not
// return an error. Similarly, see Unmask.
func (inst *Ctl) Mask(unit string, opts Options) error {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(opts.Timeout)*time.Second)
	defer cancel()
	var args = []string{"mask", "--system", unit}
	if opts.UserMode {
		args[1] = "--user"
	}
	_, _, _, err := execute(ctx, args)
	return err
}

// Unmask one or more unit files, as specified on the command line.
// This will undo the effect of Mask.
//
// In line with systemd, Unmask will return ErrDoesNotExist if the unit
// doesn't exist, but only if it's not already masked.
// If the unit doesn't exist, but it's masked anyway, no error will be
// returned. Gross, I know. Take it up with Pottering.
func (inst *Ctl) Unmask(unit string, opts Options) error {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(opts.Timeout)*time.Second)
	defer cancel()
	var args = []string{"unmask", "--system", unit}
	if opts.UserMode {
		args[1] = "--user"
	}
	_, _, _, err := execute(ctx, args)
	return err
}
