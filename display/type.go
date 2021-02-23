package display

import "sync"

type Display struct {
	lineLock sync.RWMutex
	lines    []string

	renderLock    sync.Mutex
	renderContent []string
}
