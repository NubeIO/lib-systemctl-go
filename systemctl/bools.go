package systemctl

import (
	"context"
	"errors"
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
func IsEnabled(unit string, opts Options) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(opts.Timeout)*time.Second)
	defer cancel()
	var args = []string{"is-enabled", "--system", unit}
	if opts.UserMode {
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
func IsActive(unit string, opts Options) (active bool, status string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(opts.Timeout)*time.Second)
	defer cancel()
	var args = []string{"is-active", "--system", unit}
	if opts.UserMode {
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

// IsFailed Check whether any of the specified units are in a "failed" state.
func IsFailed(unit string, opts Options) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(opts.Timeout)*time.Second)
	defer cancel()
	var args = []string{"is-failed", "--system", unit}
	if opts.UserMode {
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

//IsInstalled checks if the program is installed
func IsInstalled(unit string, opts Options) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(opts.Timeout)*time.Second)
	defer cancel()
	var args = []string{"status", "--system", unit}
	if opts.UserMode {
		args[1] = "--user"
	}
	_, _, _, err := execute(ctx, args)

	if err != nil {
		return false, errors.New("service is not installed")
	}
	return true, nil
}
