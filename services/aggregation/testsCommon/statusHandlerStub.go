package testsCommon

import (
	"context"

	"github.com/iulianpascalau/api-monitoring/services/aggregation/common"
)

// StatusHandlerStub -
type StatusHandlerStub struct {
	NotifyAppStartHandler      func()
	ErrorEncounteredHandler    func(err error)
	CollectKeysProblemsHandler func(messages []common.OutputMessage)
	ExecuteHandler             func(ctx context.Context) error
	SendCloseMessageHandler    func()
}

// NotifyAppStart -
func (stub *StatusHandlerStub) NotifyAppStart() {
	if stub.NotifyAppStartHandler != nil {
		stub.NotifyAppStartHandler()
	}
}

// ErrorEncountered -
func (stub *StatusHandlerStub) ErrorEncountered(err error) {
	if stub.ErrorEncounteredHandler != nil {
		stub.ErrorEncounteredHandler(err)
	}
}

// CollectKeysProblems -
func (stub *StatusHandlerStub) CollectKeysProblems(messages []common.OutputMessage) {
	if stub.CollectKeysProblemsHandler != nil {
		stub.CollectKeysProblemsHandler(messages)
	}
}

// Execute -
func (stub *StatusHandlerStub) Execute(ctx context.Context) error {
	if stub.ExecuteHandler != nil {
		return stub.ExecuteHandler(ctx)
	}

	return nil
}

// SendCloseMessage -
func (stub *StatusHandlerStub) SendCloseMessage() {
	if stub.SendCloseMessageHandler != nil {
		stub.SendCloseMessageHandler()
	}
}

// IsInterfaceNil -
func (stub *StatusHandlerStub) IsInterfaceNil() bool {
	return stub == nil
}
