package systemctl

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// IsEnabled Checks whether any of the specified unit files are enabled (as with enable).
//
// Returns true if the unit is enabled, aliased, static, indirect, generated
// or transient.
//
// Returns false if disabled. Also returns an error if linked, masked, or bad.
//
// See https://www.freedesktop.org/software/systemd/man/systemctl.html#is-enabled%20UNIT%E2%80%A6
// for more information
func (inst *SystemCtl) IsEnabled(unit string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(inst.Timeout)*time.Second)
	defer cancel()
	var args = []string{"is-enabled", "--system", unit}
	if inst.UserMode {
		args[1] = "--user"
	}
	stdout, _, _, err := execute(ctx, args)
	stdout = strings.TrimSuffix(stdout, "\n")
	switch stdout {
	case "enabled":
		return true, nil
	case "enabled-runtime":
		return true, nil
	case "linked":
		return false, ErrLinked
	case "linked-runtime":
		return false, ErrLinked
	case "alias":
		return true, nil
	case "masked":
		return false, ErrMasked
	case "masked-runtime":
		return false, ErrMasked
	case "static":
		return true, nil
	case "indirect":
		return true, nil
	case "disabled":
		return false, nil
	case "generated":
		return true, nil
	case "transient":
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, ErrUnspecified
}

// IsActive Check whether any of the specified units are active (i.e. running).
//
// Returns true if the unit is active, false if inactive or failed.
// Also returns false in an error case.
func (inst *SystemCtl) IsActive(unit string) (active bool, status string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(inst.Timeout)*time.Second)
	defer cancel()
	var args = []string{"is-active", "--system", unit}
	if inst.UserMode {
		args[1] = "--user"
	}
	stdout, _, _, _ := execute(ctx, args)
	stdout = strings.TrimSuffix(stdout, "\n")
	switch stdout {
	case "inactive":
		return false, "inactive", nil
	case "active":
		return true, "active", nil
	case "failed":
		return false, "failed", nil
	case "activating":
		return false, "activating", nil
	default:
		return false, "", errors.New("failed or service is not installed")
	}
}

// IsRunning Check whether specified units is in a "running" state.
func (inst *SystemCtl) IsRunning(unit string) (active bool, status string, err error) {
	stats, err := inst.State(unit)
	if err != nil {
		return false, string(stats.SubState), err
	}
	if stats.SubState != "running" {
		return false, string(stats.SubState), nil
	}
	return true, string(stats.SubState), nil
}

// IsFailed Check whether any of the specified units are in a "failed" state.
func (inst *SystemCtl) IsFailed(unit string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(inst.Timeout)*time.Second)
	defer cancel()
	var args = []string{"is-failed", "--system", unit}
	if inst.UserMode {
		args[1] = "--user"
	}
	stdout, _, _, err := execute(ctx, args)
	if matched, _ := regexp.MatchString(`inactive`, stdout); matched {
		return false, nil
	} else if matched, _ := regexp.MatchString(`active`, stdout); matched {
		return false, nil
	} else if matched, _ := regexp.MatchString(`failed`, stdout); matched {
		return true, nil
	}
	return false, err
}

// IsInstalled checks if the program is installed
func (inst *SystemCtl) IsInstalled(unit string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(inst.Timeout)*time.Second)
	defer cancel()
	var args = []string{"status", "--system", unit}
	if inst.UserMode {
		args[1] = "--user"
	}
	stdout, _, _, _ := execute(ctx, args)
	if stdout == "" {
		return false, errors.New(fmt.Sprintf("%s: service is not installed", unit))
	}
	return true, nil
}
