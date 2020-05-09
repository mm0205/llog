package llog

import "strings"

const (
	CharTab      = "\t"
	AltCharTab   = "<TAB>"
	CharColon    = ":"
	AltCharColon = "<COLON>"
	CharCR       = "\r"
	AltCharCR    = "<CR>"
	CharLF       = "\n"
	AltCharLF    = "<LF>"
)

func SimpleEscape(source string) string {
	return strings.Replace(strings.Replace(strings.Replace(strings.Replace(
		source, CharTab, AltCharTab, -1),
		CharColon, AltCharColon, -1),
		CharLF, AltCharLF, -1),
		CharCR, AltCharCR, -1)
}

