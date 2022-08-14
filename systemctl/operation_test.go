package systemctl

import (
	"fmt"
	pprint "github.com/NubeIO/lib-systemctl-go/print"
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
	action, err := service.CtlAction("restart", "nubeio-flow-framework", 10)
	fmt.Println(err)
	if err != nil {
		return
	}
	fmt.Println(action.Ok)
	fmt.Println(action.Message)
}

func TestCtl_ServiceState(t *testing.T) {
	service := New(&Ctl{
		UserMode: false,
		Timeout:  30,
	})
	action, err := service.ServiceState("nubeio-flow-framework", 10)
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(action)

}

func TestCtl_ServiceStateMass(t *testing.T) {
	service := New(&Ctl{
		UserMode: false,
		Timeout:  30,
	})
	action, err := service.ServiceStateMass([]string{"nubeio-flow-framework"}, 10)
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(action)

}
