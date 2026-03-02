package executors

import (
	"context"

	"github.com/iulianpascalau/api-monitoring/services/aggregation/common"
)

// Notifier defines the interface for sending alarm notifications
type Notifier interface {
	OutputMessages(messages ...common.OutputMessage) error
	Name() string
	IsInterfaceNil() bool
}

// OutputNotifiersHandler defines the behavior of a component that is able to notify all notifiers
type OutputNotifiersHandler interface {
	NotifyWithRetry(caller string, messages ...common.OutputMessage) error
	IsInterfaceNil() bool
}

// Executor defines the behavior of a component able to execute a certain task
type Executor interface {
	Execute(ctx context.Context) error
	IsInterfaceNil() bool
}
