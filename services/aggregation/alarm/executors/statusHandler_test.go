package executors

import (
	"context"
	"errors"
	"testing"

	"github.com/iulianpascalau/api-monitoring/services/aggregation/common"
	"github.com/iulianpascalau/api-monitoring/services/aggregation/testsCommon"
	"github.com/stretchr/testify/assert"
)

func TestNewStatusHandler(t *testing.T) {
	t.Parallel()

	t.Run("nil notifiers handler should error", func(t *testing.T) {
		t.Parallel()

		handler, err := NewStatusHandler(nil)
		assert.Nil(t, handler)
		assert.Equal(t, errNilOutputNotifiersHandler, err)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		handler, err := NewStatusHandler(&testsCommon.OutputNotifiersHandlerStub{})
		assert.NotNil(t, handler)
		assert.Nil(t, err)
	})
}

func TestStatusHandler_NotifyAppStart(t *testing.T) {
	t.Parallel()

	var returnedErr error
	sentMessages := make([]common.OutputMessage, 0)
	outputNotifiersHandler := &testsCommon.OutputNotifiersHandlerStub{
		NotifyWithRetryHandler: func(caller string, messages ...common.OutputMessage) error {
			sentMessages = append(sentMessages, messages...)

			return returnedErr
		},
	}

	handler, _ := NewStatusHandler(outputNotifiersHandler)
	assert.Equal(t, 0, len(sentMessages)) // should not notify at startup

	t.Run("notifiers handler does not error", func(t *testing.T) {
		returnedErr = nil

		handler.NotifyAppStart()

		assert.Equal(t, 1, len(sentMessages))
		assert.Equal(t, common.ExecutorName, sentMessages[0].ExecutorName)
		assert.Equal(t, common.InfoMessageOutputType, sentMessages[0].Type)
		assert.Contains(t, sentMessages[0].Identifier, "Application started on")
		assert.Zero(t, handler.NumErrorsEncountered())
	})
	t.Run("notifiers handler errors", func(t *testing.T) {
		returnedErr = errors.New("expected error")

		handler.NotifyAppStart()

		assert.Equal(t, 2, len(sentMessages))
		assert.Equal(t, common.ExecutorName, sentMessages[0].ExecutorName)
		assert.Equal(t, common.InfoMessageOutputType, sentMessages[0].Type)
		assert.Contains(t, sentMessages[0].Identifier, "Application started on")
		assert.Equal(t, uint32(1), handler.NumErrorsEncountered())
	})
}

func TestStatusHandler_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	var instance *statusHandler
	assert.True(t, instance.IsInterfaceNil())

	instance = &statusHandler{}
	assert.False(t, instance.IsInterfaceNil())
}

func TestStatusHandler_Execute(t *testing.T) {
	t.Parallel()

	sentMessages := make([]common.OutputMessage, 0)
	outputNotifiersHandler := &testsCommon.OutputNotifiersHandlerStub{
		NotifyWithRetryHandler: func(caller string, messages ...common.OutputMessage) error {
			sentMessages = append(sentMessages, messages...)

			return nil
		},
	}

	notifier, _ := NewStatusHandler(outputNotifiersHandler)
	sentMessages = make([]common.OutputMessage, 0) // reset the constructor sent messages

	t.Run("empty state should return info messages", func(t *testing.T) {
		err := notifier.Execute(context.Background())
		assert.Nil(t, err)

		expectedMessageErr := common.OutputMessage{
			Type:         common.InfoMessageOutputType,
			Identifier:   "No application errors occurred",
			ExecutorName: common.ExecutorName,
		}
		expectedMessageKeys := common.OutputMessage{
			Type:         common.InfoMessageOutputType,
			Identifier:   "All monitored metrics are performing as expected",
			ExecutorName: common.ExecutorName,
		}

		assert.Equal(t, []common.OutputMessage{expectedMessageErr, expectedMessageKeys}, sentMessages)

		sentMessages = make([]common.OutputMessage, 0)
	})
	t.Run("2 errors should return warn messages", func(t *testing.T) {
		notifier.ErrorEncountered(nil)
		notifier.ErrorEncountered(errors.New("error 1"))
		notifier.ErrorEncountered(errors.New("error 2"))
		err := notifier.Execute(context.Background())
		assert.Nil(t, err)

		expectedMessageErr := common.OutputMessage{
			Type:         common.WarningMessageOutputType,
			Identifier:   "2 application error(s) occurred, please check the app logs",
			ExecutorName: common.ExecutorName,
		}
		expectedMessageKeys := common.OutputMessage{
			Type:         common.InfoMessageOutputType,
			Identifier:   "All monitored metrics are performing as expected",
			ExecutorName: common.ExecutorName,
		}

		assert.Equal(t, []common.OutputMessage{expectedMessageErr, expectedMessageKeys}, sentMessages)

		sentMessages = make([]common.OutputMessage, 0)
	})
	t.Run("2 problematic keys found should return warn messages", func(t *testing.T) {
		key1 := common.OutputMessage{
			Identifier: "vm1",
		}
		key2 := common.OutputMessage{
			Identifier: "vm2",
		}

		notifier.CollectKeysProblems([]common.OutputMessage{key2})
		notifier.CollectKeysProblems([]common.OutputMessage{key1, key2})
		notifier.CollectKeysProblems([]common.OutputMessage{key1})
		notifier.CollectKeysProblems(nil)

		err := notifier.Execute(context.Background())
		assert.Nil(t, err)

		expectedMessageErr := common.OutputMessage{
			Type:         common.InfoMessageOutputType,
			Identifier:   "No application errors occurred",
			ExecutorName: common.ExecutorName,
		}
		expectedMessageKeys := common.OutputMessage{
			Type:         common.WarningMessageOutputType,
			Identifier:   "2 monitored metrics encountered problems",
			ExecutorName: common.ExecutorName,
		}

		assert.Equal(t, []common.OutputMessage{expectedMessageErr, expectedMessageKeys}, sentMessages)

		sentMessages = make([]common.OutputMessage, 0)
	})
	t.Run("empty state should return info messages", func(t *testing.T) {
		err := notifier.Execute(context.Background())
		assert.Nil(t, err)

		expectedMessageErr := common.OutputMessage{
			Type:         common.InfoMessageOutputType,
			Identifier:   "No application errors occurred",
			ExecutorName: common.ExecutorName,
		}
		expectedMessageKeys := common.OutputMessage{
			Type:         common.InfoMessageOutputType,
			Identifier:   "All monitored metrics are performing as expected",
			ExecutorName: common.ExecutorName,
		}

		assert.Equal(t, []common.OutputMessage{expectedMessageErr, expectedMessageKeys}, sentMessages)

		sentMessages = make([]common.OutputMessage, 0)
	})
}

func TestStatusHandler_SendCloseMessage(t *testing.T) {
	t.Parallel()

	sentMessages := make([]common.OutputMessage, 0)
	outputNotifiersHandler := &testsCommon.OutputNotifiersHandlerStub{
		NotifyWithRetryHandler: func(caller string, messages ...common.OutputMessage) error {
			sentMessages = append(sentMessages, messages...)

			return nil
		},
	}

	notifier, _ := NewStatusHandler(outputNotifiersHandler)
	sentMessages = make([]common.OutputMessage, 0) // reset the constructor sent messages

	notifier.SendCloseMessage()

	expectedMessage := common.OutputMessage{
		Type:         common.WarningMessageOutputType,
		Identifier:   "Application closing",
		ExecutorName: common.ExecutorName,
	}

	assert.Equal(t, []common.OutputMessage{expectedMessage}, sentMessages)
}
