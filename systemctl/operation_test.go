package systemctl

import (
	"fmt"
	"testing"
)

func TestCtl_CtlStatus(t *testing.T) {
	service := New(&Ctl{
		UserMode: false,
		Timeout:  30,
	})

	action, err := service.CtlStatus("isRunning", "nubeio-flow-framework", 10)
	fmt.Println(err)
	if err != nil {
		return
	}
	fmt.Println(action)
}

func TestCtl_CtlAction(t *testing.T) {
	service := New(&Ctl{
		UserMode: false,
		Timeout:  30,
	})

	action, err := service.CtlAction("start", "nubeio-flow-framework", 10)
	fmt.Println(err)
	if err != nil {
		return
	}
	fmt.Println(action.Ok)
	fmt.Println(action.Message)
}
