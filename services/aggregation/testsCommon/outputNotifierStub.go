package testsCommon

import (
	"github.com/iulianpascalau/api-monitoring/services/aggregation/common"
)

// NotifierStub -
type NotifierStub struct {
	NameHandler           func() string
	OutputMessagesHandler func(messages ...common.OutputMessage) error
}

// OutputMessages -
func (stub *NotifierStub) OutputMessages(messages ...common.OutputMessage) error {
	if stub.OutputMessagesHandler != nil {
		return stub.OutputMessagesHandler(messages...)
	}

	return nil
}

// Name -
func (stub *NotifierStub) Name() string {
	if stub.NameHandler != nil {
		return stub.NameHandler()
	}

	return ""
}

// IsInterfaceNil -
func (stub *NotifierStub) IsInterfaceNil() bool {
	return stub == nil
}
