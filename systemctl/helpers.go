package systemctl

import (
	"github.com/NubeIO/lib-systemctl-go/systemctl/properties"
	"strconv"
	"time"
)

const dateFormat = "Mon 2006-01-02 15:04:05 MST"

// GetStartTime Get start time of a service (`systemctl show [unit] --property ExecMainStartTimestamp`) as a `Time` type
func (inst *SystemCtl) GetStartTime(unit string) (time.Time, error) {
	value, err := inst.Show(unit, properties.ExecMainStartTimestamp)

	if err != nil {
		return time.Time{}, err
	}
	// ExecMainStartTimestamp returns an empty string if the unit is not running
	if value == "" {
		return time.Time{}, ErrUnitNotActive
	}
	return time.Parse(dateFormat, value)
}

// GetNumRestarts Get the number of times a process restarted (`systemctl show [unit] --property NRestarts`) as an int
func (inst *SystemCtl) GetNumRestarts(unit string) (int, error) {
	value, err := inst.Show(unit, properties.NRestarts)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(value)
}

// GetMemoryUsage Get current memory in bytes (`systemctl show [unit] --property MemoryCurrent`) an an int
func (inst *SystemCtl) GetMemoryUsage(unit string) (int, error) {
	value, err := inst.Show(unit, properties.MemoryCurrent)
	if err != nil {
		return -1, err
	}
	if value == "[not set]" {
		return -1, ErrValueNotSet
	}
	return strconv.Atoi(value)
}

// GetPID Get the PID of the main process (`systemctl show [unit] --property MainPID`) as an int
func (inst *SystemCtl) GetPID(unit string) (int, error) {
	value, err := inst.Show(unit, properties.MainPID)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(value)
}

// setTimeout limit with the timeout can be
func setTimeout(timeOut int) time.Duration {
	if timeOut <= 0 || timeOut >= 120 {
		timeOut = 10
	}
	return time.Duration(timeOut)
}
