package llog

import (
	"github.com/pkg/errors"
	"io"
	"strings"
)



type logger struct {
	writer     io.Writer
	messageKey string
	escapeFunc EscapeFunc
}

func (logger *logger) SetOutput(w io.Writer) {
	logger.writer = w
}

func (logger *logger) SetEscapeFunc(escapeFunc EscapeFunc) {
	logger.escapeFunc = escapeFunc
}

func (logger *logger) SetMessageLabel(key string) error {
	if key == "" {
		return errors.New("message key must not be empty")
	}
	logger.messageKey = key
	return nil
}

func (logger *logger) Println(message string) error {
	return logger.PrintlnItems(message)
}

func (logger *logger) PrintlnItems(message string, items ...LogItem) error {
	messageItem := &LogItemString{
		LogItemBase: LogItemBase{
			Label: logger.messageKey,
		},
		Value: message,
	}

	items = append([]LogItem{messageItem}, items...)

	return logger.printItems(items...)
}

func (logger *logger) printItems(items ...LogItem) error {
	if logger.writer == nil {
		return nil
	}

	records := make([]string, 0, len(items))
	var builder strings.Builder
	for _, item := range items {
		builder.Reset()
		err := logger.buildLogRecord(&builder, item)
		if err != nil {
			return errors.Wrapf(err, "failed to serialize key value; key=%s", item.StringLabel())
		}
		records = append(records, builder.String())
	}

	line := strings.Join(records, "\t") + "\n"

	_, err := logger.writer.Write([]byte(line))
	if err != nil {
		return errors.Wrap(err, "failed to write to output")
	}
	return nil
}

func (logger *logger) escape(source string) string {
	if logger.escapeFunc == nil {
		return source
	}
	return logger.escapeFunc(source)
}

func (logger *logger) buildLogRecord(builder *strings.Builder, item LogItem) error {
	_, err := builder.WriteString(
		logger.escape(item.StringLabel()),
	)
	if err != nil {
		return err
	}
	_, err = builder.WriteString(":")
	if err != nil {
		return err
	}
	_, err = builder.WriteString(
		logger.escape(item.StringValue()),
	)
	return err
}

var globalLogger *logger = &logger{
	writer:     nil,
	messageKey: DefaultMessageLabel,
}

