package llog

type LogItemString struct {
	LogItemBase
	Value string
}

func (item LogItemString) StringValue() string {
	return item.Value
}

func NewString(key string, value string) LogItemString {
	return LogItemString{
		LogItemBase: LogItemBase{Label: key},
		Value:       value,
	}
}
