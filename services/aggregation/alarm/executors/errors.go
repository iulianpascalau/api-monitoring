package executors

import "errors"

var (
	errNilOutputNotifier            = errors.New("nil output notifier")
	errInvalidTimeBetweenRetries    = errors.New("invalid time between retries")
	errNotificationsSendingProblems = errors.New("notification sending problems")
	errNilOutputNotifiersHandler    = errors.New("nil output notifiers handler")
	errNilTimeFunc                  = errors.New("nil pointer for the current time function")
	errNilExecutor                  = errors.New("nil executor")
	errInvalidWeekDay               = errors.New("invalid week day")
	errInvalidHour                  = errors.New("invalid hour")
	errInvalidMinute                = errors.New("invalid minute")
)
