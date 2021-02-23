package display

import (
	"unicode/utf8"

	"git.sr.ht/~poldi1405/go-ansi"
)

func (d *Display) Update(str ...string) {
	d.lineLock.Lock()
	defer d.lineLock.Unlock()

}

func (d *Display) UpdateLine(index int, content string) {
	d.lineLock.Lock()
	defer d.lineLock.Unlock()

}

func (d *Display) InsertLine(position int, content string) {
	d.renderLock.Lock()
	defer d.renderLock.Unlock()

	d.lineLock.RLock()
	defer d.lineLock.RUnlock()

	for i := 0; i < len(d.lines); i++ {

	}
}

func calculateDiff(newline, oldline string) string {
	var transform string

	newRunes := []rune(newline)
	oldRunes := []rune(oldline)

	equal := 0
	for i := 0; i < len(newRunes); i++ {
		if i >= len(oldRunes) {
			transform += string(newRunes[i])
			continue
		}

		if newRunes[i] == oldRunes[i] {
			equal++
			continue
		}

		if equal > 0 {
			transform += ansi.RightX(equal)
			equal = 0
		}
		transform += string(newRunes[i])
	}
	if len(newRunes) < len(oldRunes) {
		transform += ansi.ClearRight()
	}

	clean := cleanString(transform)

	if clean != transform {
		transform = calculateDiff(newline, clean)
	}

	return transform
}

func cleanString(str string) string {
	runes := []rune(str)
	res := ""
	for _, r := range runes {
		if utf8.ValidRune(r) {
			res += string(r)
		}
	}
	return res
}
