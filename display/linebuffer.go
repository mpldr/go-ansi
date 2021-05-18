package display

import (
	"fmt"
	"sync"

	"git.sr.ht/~poldi1405/go-ansi"
)

// LineBuffer is the type containing written lines.
type LineBuffer struct {
	// lines contains the actual lines. 0 is the bottom-most line
	lines    []string
	linesMtx sync.RWMutex
}

func (lb *LineBuffer) Set(line int, text string) {
	lb.linesMtx.RLock()
	if line >= len(lb.lines) {
		lb.linesMtx.RUnlock()
		return
	}
	oldline := lb.lines[line]
	lb.linesMtx.RUnlock()

	diff := calculateDiff(text, oldline)
	lb.linesMtx.Lock()
	defer lb.linesMtx.Unlock()

	lb.print(line, diff)
	lb.lines[line] = text
}

func (lb *LineBuffer) Add(text string) {
	lb.linesMtx.Lock()
	defer lb.linesMtx.Unlock()

	lb.lines = append([]string{text}, lb.lines...)
	fmt.Print("\n" + text)
}

func (lb *LineBuffer) print(up int, text string) {
	fmt.Print(ansi.SavePos() + "\r" + ansi.UpX(up) + text + ansi.RestorePos())
}
