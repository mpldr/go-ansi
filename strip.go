package ansi

import (
	"bytes"
	"fmt"
)

// StripString removes all ANSI-Escape sequences from the given string and
// returns the cleaned version
func StripString(str string) string {
	bts := []byte(str)
	bts = stripStandard(bts)
	bts = stripLink(bts)

	return string(bts)
}

func stripLink(bts []byte) []byte {
	for { // stop if the last run did not match anything
		index := bytes.Index(bts, []byte{0x1b, '\\'})
		if index == -1 {
			break
		}
		removeUntil := bytes.Index(bts[index+1:], []byte{0x1b})
		if removeUntil == -1 {
			break
		}

		bts = append(bts[:index], bts[index+removeUntil+1:]...)
	}
	bts = bytes.ReplaceAll(bts, []byte{0x1b, ']', '8', ';', ';', 0x1b, '\\'}, []byte{})
	bts = bytes.ReplaceAll(bts, []byte{0x1b, ']', '8', ';', ';'}, []byte{})
	return bts
}

func stripStandard(bts []byte) []byte {
	matched := true
	for matched { // stop if the last run did not match anything
		index := bytes.Index(bts, []byte{0x1b, '['})
		if index == -1 {
			break
		}
		removeUntil := index + 2
		for i := removeUntil; i < len(bts); i++ {
			if (bts[i] >= 0x30 && bts[i] <= 0x39) || bts[i] == 0x3b {
				removeUntil++
				continue
			}
			if (bts[i] >= 0x41 && bts[i] <= 0x5a) || (bts[i] >= 0x61 && bts[i] <= 0x7a) {
				removeUntil++
				break
			}
			// we shouldn't be here in valid codes. anyway. stop
			// what we're doing before be break something
			fmt.Print("oppsie! ")
			fmt.Println(bts[index:removeUntil])
			matched = false
			break
		}

		bts = append(bts[:index], bts[removeUntil:]...)
	}
	return bts
}
