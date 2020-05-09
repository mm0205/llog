package llog

import (
	"fmt"
	"reflect"
)

// DefaultErrorLabel is the default key of logItemError.
const DefaultErrorLabel = "_error"

// DefaultErrorFormat is the format which is used for formatting the error value.
const DefaultErrorFormat = "%v"

// DefaultErrorStackTraceFormat is the format for output stack trace of error.
const DefaultErrorStackTraceFormat = "%+v"

// logItemError is a log item for output error.
type logItemError struct {
	LogItemBase
	value            error
	format           string
	autoFormat       bool
	escapeStackTrace EscapeFunc
}

// NewError returns an instance of logItemError
func NewError(err error) LogItem {
	return NewErrorKey(err, DefaultErrorLabel)
}

func NewErrorKey(err error, key string) LogItem {
	return NewErrorKeyFormat(err, key, DefaultErrorFormat)
}

func NewErrorKeyFormat(err error, key string, format string) LogItem {
	return logItemError{
		LogItemBase:      LogItemBase{Label: key},
		value:            err,
		format:           format,
		autoFormat:       false,
		escapeStackTrace: nil,
	}
}

func NewStackTrace(err error, escapeFunc EscapeFunc) LogItem {
	return logItemError{
		LogItemBase:      LogItemBase{Label: DefaultErrorLabel},
		value:            err,
		format:           DefaultErrorFormat,
		autoFormat:       true,
		escapeStackTrace: escapeFunc,
	}
}

func (item logItemError) StringValue() string {
	if !item.autoFormat {
		return fmt.Sprintf(item.format, item.value)
	}

	t := reflect.TypeOf(item.value)
	if _, ok := t.MethodByName("StackTrace"); ok {
		st := fmt.Sprintf(DefaultErrorStackTraceFormat, item.value)
		if item.escapeStackTrace == nil {
			return st
		}
		return item.escapeStackTrace(st)
	}
	return fmt.Sprintf(DefaultErrorFormat, item.value)
}
