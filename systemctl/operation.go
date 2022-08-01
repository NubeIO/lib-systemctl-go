package systemctl

import (
	"errors"
	"fmt"
)

type SystemResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type MassSystemResponse struct {
	Service string `json:"service"`
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type MassSystemResponseChecks struct {
	Service string `json:"service"`
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// ServiceMassAction mass start, stop, enable, disable, a service
func (inst *Ctl) ServiceMassAction(serviceNames []string, action string, timeout int) ([]MassSystemResponse, error) {
	var out []MassSystemResponse
	var msg MassSystemResponse
	for _, name := range serviceNames {
		msg.Service = name
		_, err := inst.CtlAction(action, name, timeout)
		if err != nil {
			msg.Ok = false
			msg.Message = err.Error()
		} else {
			msg.Ok = true
			msg.Message = fmt.Sprintf("%s ok!", action)
		}
		out = append(out, msg)
	}
	return out, nil
}

// ServiceMassCheck check if a service isRunning, isEnabled and so on
func (inst *Ctl) ServiceMassCheck(serviceNames []string, action string, timeout int) ([]MassSystemResponseChecks, error) {
	systemOpts.Timeout = timeout
	var out []MassSystemResponseChecks
	var msg MassSystemResponseChecks
	for _, name := range serviceNames {
		msg.Service = name
		ctlAction, err := inst.CtlStatus(action, name, timeout)
		if err != nil {
			msg.Ok = false
			msg.Message = err.Error()
		} else {
			msg.Ok = true
			msg.Message = ctlAction.Message
		}
		out = append(out, msg)
	}
	return out, nil
}

func (inst *Ctl) CtlAction(action, unit string, timeout int) (*SystemResponse, error) {
	systemOpts.Timeout = timeout
	resp := &SystemResponse{}
	var err error
	switch action {
	case start.String():
		err = inst.Start(unit, systemOpts)
	case stop.String():
		err = inst.Stop(unit, systemOpts)
	case enable.String():
		err = inst.Enable(unit, systemOpts)
	case disable.String():
		err = inst.Disable(unit, systemOpts)
	default:
		return nil, errors.New("no valid action found try, start, stop, enable or disable")
	}
	if err == nil {
		resp.Message = fmt.Sprintf("service:%s failed to:%s", unit, action)
		return resp, err
	} else {
		resp.Ok = true
		resp.Message = fmt.Sprintf("service:%s action:%s ok!", unit, action)
		return resp, err
	}
}

type SystemResponseChecks struct {
	Is      bool   `json:"is"`
	Message string `json:"message"`
}

func (inst *Ctl) CtlStatus(action, unit string, timeout int) (*SystemResponseChecks, error) {
	systemOpts.Timeout = timeout
	actionResp := &SystemResponseChecks{}
	switch action {
	case isRunning.String():
		running, status, err := inst.IsRunning(unit, systemOpts)
		if err != nil {
			return nil, err
		}
		actionResp.Is = running
		actionResp.Message = status

	case isInstalled.String():
		installed, err := inst.IsInstalled(unit, systemOpts)
		if err != nil {
			actionResp.Message = "is not installed"
			return nil, err
		}
		actionResp.Is = installed
		actionResp.Message = "is installed"

	case isEnabled.String():
		enabled, err := inst.IsEnabled(unit, systemOpts)
		if err != nil {
			actionResp.Message = "is not enabled"
			return nil, err
		}
		actionResp.Is = enabled
		actionResp.Message = "is enabled"

	case isActive.String():
		active, sts, err := inst.IsActive(unit, systemOpts)
		if err != nil {
			actionResp.Message = sts
			return nil, err
		}
		actionResp.Is = active
		actionResp.Message = sts

	case isFailed.String():
		failed, err := inst.IsFailed(unit, systemOpts)
		if err != nil {
			actionResp.Message = "is not failed"
			return nil, err
		}
		actionResp.Is = failed
		actionResp.Message = "is failed"

	}
	return actionResp, errors.New("no valid action found try, isRunning, isInstalled, isEnabled, isActive or isFailed")
}

func (inst *Ctl) ServiceStats(serviceName string, timeout int) (resp SystemState, err error) {
	systemOpts.Timeout = timeout
	resp, err = inst.State(serviceName, systemOpts)
	if err != nil {
		return resp, err
	}
	return resp, nil
}