package llog

// LogItem is the interface for items of log.
// The log item is formatted as following
//	item.StringLabel() + ":" + item.StringValue()
type LogItem interface {

	// StringLabel returns the label string of this item.
	StringLabel() string

	// StringValue returns the string value of this item.
	StringValue() string
}

// LogItemBase is base class of log items.
// This has only `Label` field, which is used as label of lstv.
type LogItemBase struct {
	// Label is the label of this item.
	Label string
}

// StringLabel returns the items label itself.
func (item LogItemBase) StringLabel() string {
	return item.Label
}

