package stat

import (
	"sync"
	"time"
)

// Mgr tick management
type Mgr struct {
	sync.RWMutex
	data map[string]*ticks
}

// New create management
func New(interval time.Duration) *Mgr {
	mgr := &Mgr{data: make(map[string]*ticks)}
	go func() {
		for {
			time.Sleep(interval)
			mgr.Collect()
		}
	}()
	return mgr
}

// New create new tick
func (mgr *Mgr) New(name string) *Tick {
	tick := &Tick{
		begin: time.Now(),
		end:   time.Now(),
	}
	mgr.Lock()
	list, ok := mgr.data[name]
	if !ok {
		list = newTicks(name)
		mgr.data[name] = list
	}
	mgr.Unlock()
	list.push(tick)
	return tick
}

// Collect compute default values and export
func (mgr *Mgr) Collect() {
	mgr.RLock()
	defer mgr.RUnlock()
	for _, tks := range mgr.data {
		tks.collect()
	}
}
