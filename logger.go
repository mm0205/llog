package llog

import (
	"io"
)

// DefaultMessageLabel is the label for the message.
const DefaultMessageLabel string = "_message"

// Logger is the interface of logger.
type Logger interface {
	// SetOutput sets the output writer for log.
	SetOutput(w io.Writer)

	// SetEscapeFunc sets the escape function.
	SetEscapeFunc(escapeFunc EscapeFunc)

	// SetMessageLabel sets the label of message.
	SetMessageLabel(key string) error

	// Println writes a message to the output writer.
	// The format is following
	//	_message:{message}
	// The label of message can be changed by `SetMessageLabel`.
	Println(message string) error

	// PrintlnItems writes a message and additional items to the output writer.
	// The format is following
	//	_message:{message}	{items[0].StringLabel()}:{items[0].StringValue}	...
	PrintlnItems(message string, items ...LogItem) error
}

// New returns a new logger.
func New(w io.Writer) Logger {
	return &logger{
		writer:     w,
		messageKey: DefaultMessageLabel,
		escapeFunc: nil,
	}
}

// Global returns a global instance of the Logger.
func Global() Logger {
	return globalLogger
}

// SetOutput sets the output writer of the global logger.
func SetOutput(w io.Writer) {
	globalLogger.SetOutput(w)
}

// SetEscapeFunc sets the escape function of the global logger.
func SetEscapeFunc(escapeFunc EscapeFunc) {
	globalLogger.SetEscapeFunc(escapeFunc)
}

// Println calls `Println` of the global logger.
func Println(message string) error {
	return globalLogger.Println(message)
}

// PrintlnItems calls `PrintlnItems` of the global logger.
func PrintlnItems(message string, items ...LogItem) error {
	return globalLogger.PrintlnItems(message, items...)
}
