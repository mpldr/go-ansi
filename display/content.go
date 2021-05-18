package display

import (
	"strings"
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

	// no, I do not want to make a dumb wordplay like oldrune schoolscape
	oldrunes := []rune(oldline)
	newrunes := []rune(newline)

	equal := 0
	for i := 0; i < len(newrunes); i++ {
		if i >= len(oldrunes) {
			transform += string(newrunes[i])
			continue
		}

		if newrunes[i] == oldrunes[i] {
			equal++
			continue
		}

		if equal > 5 {
			transform += ansi.RightX(equal)
			equal = 0
		}

		if equal > 0 {
			transform += strings.Repeat(" ", equal)
			equal = 0
		}

		if !utf8.ValidRune(rune(newrunes[i])) {
			transform += "?"
			continue
		}

		transform += string(newrunes[i])
	}
	if len(newrunes) < len(oldrunes) {
		transform += ansi.ClearRight()
	}

	// transform = fixUnicode(newline, oldline, transform)

	return transform
}
