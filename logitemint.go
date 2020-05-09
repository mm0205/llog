package llog

import "fmt"

const defaultIntFormat = "%d"

type LogItemInt struct {
	LogItemBase
	Value  int
	Format string
}

func (item LogItemInt) StringValue() string {
	return fmt.Sprintf(item.Format, item.Value)
}

func NewInt(key string, value int) LogItemInt {
	return NewIntFormat(key, value, defaultIntFormat)
}

func NewIntFormat(key string, value int, format string) LogItemInt {
	return LogItemInt{
		LogItemBase: LogItemBase{
			Label: key,
		},
		Value:       value,
		Format:      format,
	}
}
