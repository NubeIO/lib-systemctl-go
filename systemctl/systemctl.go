package systemctl

import (
	"context"
	"github.com/NubeIO/lib-systemctl-go/systemctl/properties"
	"strings"
	"time"
)

// Disable one or more units.
//
// This removes all symlinks to the unit files backing the specified units from
// the unit configuration directory, and hence undoes any changes made by
// enable or link.
func (inst *Ctl) Disable(unit string, opts Options) error {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(opts.Timeout)*time.Second)
	defer cancel()
	var args = []string{"disable", "--system", unit}
	if opts.UserMode {
		args[1] = "--user"
	}
	_, _, _, err := execute(ctx, args)
	return err
}

// Enable one or more units or unit instances.
//
// This will create a set of symlinks, as encoded in the [Install] sections of
// the indicated unit files. After the symlinks have been created, the system
// manager configuration is reloaded (in a way equivalent to daemon-reload),
// in order to ensure the changes are taken into account immediately.
func (inst *Ctl) Enable(unit string, opts Options) error {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(opts.Timeout)*time.Second)
	defer cancel()
	var args = []string{"enable", "--system", unit}
	if opts.UserMode {
		args[1] = "--user"
	}
	_, _, _, err := execute(ctx, args)
	return err
}

// Show a selected property of a unit. Accepted properties are predefined in the
// properties' subpackage to guarantee properties are valid and assist code-completion.
func (inst *Ctl) Show(unit string, property properties.Property, opts Options) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(opts.Timeout)*time.Second)
	defer cancel()
	var args = []string{"show", "--system", unit, "--property", string(property)}
	if opts.UserMode {
		args[1] = "--user"
	}
	stdout, _, _, err := execute(ctx, args)
	stdout = strings.TrimPrefix(stdout, string(property)+"=")
	stdout = strings.TrimSuffix(stdout, "\n")
	return stdout, err
}

// Start (activate) a given unit
func (inst *Ctl) Start(unit string, opts Options) error {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(opts.Timeout)*time.Second)
	defer cancel()
	var args = []string{"start", "--system", unit}
	if opts.UserMode {
		args[1] = "--user"
	}
	_, _, _, err := execute(ctx, args)
	return err
}

// Status Get back the status string which would be returned by running
// `systemctl status [unit]`.
//
// Generally, it makes more sense to programmatically retrieve the properties
// using Show, but this command is provided for the sake of completeness
func (inst *Ctl) Status(unit string, opts Options) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(opts.Timeout)*time.Second)
	defer cancel()
	var args = []string{"status", "--system", unit}
	if opts.UserMode {
		args[1] = "--user"
	}
	stdout, _, _, err := execute(ctx, args)
	if stdout == "" {
		stdout = "service:" + unit + " not found"
	}
	return stdout, err
}

// Stop (deactivate) a given unit
func (inst *Ctl) Stop(unit string, opts Options) error {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(opts.Timeout)*time.Second)
	defer cancel()
	var args = []string{"stop", "--system", unit}
	if opts.UserMode {
		args[1] = "--user"
	}
	_, _, _, err := execute(ctx, args)
	return err
}
