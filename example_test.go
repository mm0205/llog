package llog_test

import (
	"github.com/mm0205/llog"
	"os"
)

func Example() {

	// llog outputs nothing by default.
	// Set the `io.Writer` for output.
	llog.SetOutput(os.Stdout)

	_ = llog.Println("my message")

	// Output: _message:my message
}
