package gemini

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type dType = map[string][]string

type instuct struct {
	mu      sync.RWMutex
	data    dType
	lastMod time.Time
	path    string
}

func (i *instuct) reloadIfNeeded() error {
	info, err := os.Stat(i.path)
	if err != nil {
		return err
	}

	if info.ModTime().Equal(i.lastMod) {
		return nil
	}

	file, err := os.ReadFile(i.path)
	if err != nil {
		return err
	}

	var newData dType
	if err := json.Unmarshal(file, &newData); err != nil {
		return err
	}

	i.mu.Lock()
	i.data = newData
	i.lastMod = info.ModTime()
	i.mu.Unlock()

	return nil
}

func (i *instuct) get() dType {
	i.reloadIfNeeded()
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.data
}
