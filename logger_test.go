package llog_test

import (
	"bytes"
	"fmt"
	"github.com/mm0205/llog"
	"github.com/pkg/errors"
	"runtime"
	"strings"
	"testing"
)

type testCase struct {
	message  string
	items    []llog.LogItem
	expected string
}

func Test(t *testing.T) {

	testCases := []testCase{
		{
			message: "test",
			items: []llog.LogItem{
				llog.NewString("additional", "value"),
			},
			expected: "_message:test\tadditional:value\n",
		},
		{
			message: "test",
			items: []llog.LogItem{
				llog.NewString("additional", "value"),
				llog.NewInt("intValue", 10),
			},
			expected: "_message:test\tadditional:value\tintValue:10\n",
		},
		{
			message: "test",
			items: []llog.LogItem{
				llog.NewString("additional", "value"),
				llog.NewIntFormat("intValue", 10, "%05d"),
			},
			expected: "_message:test\tadditional:value\tintValue:00010\n",
		},
		{
			message: "te\tst",
			items: []llog.LogItem{
				llog.NewString("additional", "va\tl:ue"),
				llog.NewIntFormat("intValue", 10, "%05d"),
			},
			expected: "_message:te<TAB>st\tadditional:va<TAB>l<COLON>ue\tintValue:00010\n",
		},
		{
			message: "te\tst",
			items: []llog.LogItem{
				llog.NewString("additional", "va\tl:ue"),
				llog.NewIntFormat("intValue", 10, "%05d"),
				llog.NewError(fmt.Errorf("myerror")),
			},
			expected: "_message:te<TAB>st\tadditional:va<TAB>l<COLON>ue\tintValue:00010\t_error:myerror\n",
		},
	}

	for _, testCase := range testCases {
		var buf bytes.Buffer
		logger := llog.New(&buf)
		logger.SetEscapeFunc(llog.SimpleEscape)

		err := logger.PrintlnItems(
			testCase.message,
			testCase.items...
		)

		if err != nil {
			t.Fatal(err)
		}

		line := buf.String()
		if testCase.expected != line {
			t.Errorf("expected: %s, actual: %s", testCase.expected, line)
		}
	}
}

func TestStackTrace(t *testing.T) {
	var buf bytes.Buffer
	logger := llog.New(&buf)

	err := logger.PrintlnItems("test", llog.NewStackTrace(errors.New("test"), nil))
	if err != nil {
		t.Fatal(err)
	}
	const expectedPrefix string = "_message:test\t_error:test"
	message := buf.String()
	if !strings.HasPrefix(message, expectedPrefix) {
		t.Errorf("has no expected prefix, expected=%s, actual=%s", expectedPrefix, message)
	}
	if len(message) <= len(expectedPrefix) {
		t.Errorf("has no stack trace, actual=%s", message)
	}
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to get caller")
	}
	if !strings.Contains(message, file) {
		t.Errorf("has no file name, actual=%s", message)
	}
}
