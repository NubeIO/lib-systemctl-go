package systemctl

import (
	"context"
	"strings"
	"time"
)

// UnitFileState is state type for Systemctl
type UnitFileState string

// ActiveState is state type for Systemctl
type ActiveState string

// SubState is state type for Systemctl
type SubState string

const (
	// Enabled is a state reported by systemctl
	Enabled = UnitFileState("enabled")

	// Disabled is a state reported by systemctl
	Disabled = UnitFileState("disabled")

	// Active is a state reported by systemctl
	Active = ActiveState("active")

	// Inactive is a state reported by systemctl
	Inactive = ActiveState("inactive")

	// Running is a substate reported by systemctl
	Running = SubState("running")

	// Dead is a substate reported by systemctl
	Dead = SubState("dead")
)

type SystemState struct {
	ServiceName            string        `json:"service_name,omitempty"`
	State                  UnitFileState `json:"state,omitempty"`        // enabled, disabled
	ActiveState            ActiveState   `json:"active_state,omitempty"` // active, inactive
	SubState               SubState      `json:"sub_state,omitempty"`    // running, dead
	ActiveEnterTimestamp   string        `json:"active_enter_timestamp,omitempty"`
	InactiveEnterTimestamp string        `json:"inactive_enter_timestamp,omitempty"`
	Restarts               string        `json:"restarts,omitempty"` // number of restart
	IsInstalled            bool          `json:"is_installed,omitempty"`
}

// State get status
func (inst *SystemCtl) State(unit string) (SystemState, error) {
	stats := SystemState{
		ServiceName: unit,
	}
	_, err := inst.IsInstalled(unit)
	if err != nil {
		stats.IsInstalled = false
		return stats, nil
	} else {
		stats.IsInstalled = true
	}
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(inst.Timeout)*time.Second)
	defer cancel()
	var args = []string{"show", unit, "--no-page"}
	if inst.UserMode {
		args[1] = "--user"
	}
	stdout, _, _, err := execute(ctx, args)

	unitState := UnitFileState("")
	activeState := ActiveState("")
	subState := SubState("")

	for _, line := range strings.Split(stdout, "\n") {
		fields := strings.SplitN(line, "=", 2)
		if len(fields) != 2 {
			continue
		}
		switch fields[0] {
		case "UnitFileState":
			unitState = UnitFileState(fields[1])
		case "ActiveState":
			activeState = ActiveState(fields[1])
		case "SubState":
			subState = SubState(fields[1])
		case "ActiveEnterTimestamp":
			stats.ActiveEnterTimestamp = fields[1]
		case "InactiveEnterTimestamp":
			stats.InactiveEnterTimestamp = fields[1]
		case "NRestarts":
			stats.Restarts = fields[1]
		default:
			// ignore
		}
	}

	stats.State = unitState
	stats.ActiveState = activeState
	stats.SubState = subState
	return stats, err
}
