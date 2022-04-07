package ansi

import (
	"regexp"
	"unicode/utf8"
)

func GetLengthWithoutCodes(content string) int {
	re := regexp.MustCompile(`(?m)\x1b(\[([\d;]+m|[\d;]+(H|A|B|C|D|J|K|S|T)|s|u)|\](8;;[^\x1b]*\x1b\\|777;notify;[^;]*;[^\a]*\a))`)
	return utf8.RuneCountInString(re.ReplaceAllString(content, ""))
}
