/*
llog provides simple logging functionality.

Output Destination

`llog` outputs nothing by default.
You should set appropriate `io.Writer`.

	llog.SetOutput(os.Stdout)

Log format

The output format is lstv.
	llog.Println("string value")
	// Output: _message:string value

Additional Information For Log

If you need additional information, use `llog.PrintlnItems`.
	llog.PrintlnItems("main message", llog.NewString("additional", "value"))
	// Output: _message:main message	additional:value

Escape

`llog` doesn't escape any characters by default.
If you need to escape, use `llog.SetEscapeFunc`.
	llog.SetEscapeFunc(llog.SimpleEscape)
	llog.Println("a\tb")
	// Output: _message:a<TAB>b

`llog` formats `error` using '%v'.

 */
package llog // import "github.com/mm0205/llog"
