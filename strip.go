package ansi

import "regexp"

// StripString removes all ANSI-Escape sequences from the given string and
// returns the cleaned version
func StripString(str string) string {
	return DetectionPattern.ReplaceAllString(str, "")
}

var DetectionPattern = regexp.MustCompile(`(?m)((\x1b\[[;\d]*[A-Za-z])*)`)
