package llog

// EscapeFunc is a type for escape function.
// This is used for escaping key and value of log items.
//
//	llog.SetEscapeFunc(func(source string) {
//		return strings.Replace(source, "a", "A", -1)
//	})
//	llog.Println("abc")
//	// Output: _message:Abc
//
type EscapeFunc func(source string) string
