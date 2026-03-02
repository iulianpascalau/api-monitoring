package testsCommon

import (
	"github.com/iulianpascalau/api-monitoring/services/aggregation/common"
)

// OutputNotifiersHandlerStub -
type OutputNotifiersHandlerStub struct {
	NotifyWithRetryHandler func(caller string, messages ...common.OutputMessage) error
}

// NotifyWithRetry -
func (stub *OutputNotifiersHandlerStub) NotifyWithRetry(caller string, messages ...common.OutputMessage) error {
	if stub.NotifyWithRetryHandler != nil {
		return stub.NotifyWithRetryHandler(caller, messages...)
	}

	return nil
}

// IsInterfaceNil -
func (stub *OutputNotifiersHandlerStub) IsInterfaceNil() bool {
	return stub == nil
}
