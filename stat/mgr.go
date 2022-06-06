package stat

import (
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Mgr tick management
type Mgr struct {
	sync.RWMutex
	handler http.Handler
	ticks   map[string]*ticks
	counter map[string]*Counter
}

// New create management
func New(interval time.Duration) *Mgr {
	mgr := &Mgr{
		handler: promhttp.Handler(),
		ticks:   make(map[string]*ticks),
		counter: make(map[string]*Counter),
	}
	go func() {
		for {
			mgr.Collect()
			time.Sleep(interval)
		}
	}()
	return mgr
}

// NewTick create new tick
func (mgr *Mgr) NewTick(name string) *Tick {
	tick := &Tick{
		begin: time.Now(),
		end:   time.Now(),
	}
	mgr.Lock()
	list, ok := mgr.ticks[name]
	if !ok {
		list = newTicks(name)
		mgr.ticks[name] = list
	}
	mgr.Unlock()
	list.push(tick)
	return tick
}

// NewCounter create new counter
func (mgr *Mgr) NewCounter(name string) *Counter {
	mgr.Lock()
	ct, ok := mgr.counter[name]
	if !ok {
		ct = newCounter(name)
		mgr.counter[name] = ct
	}
	mgr.Unlock()
	return ct
}

// Collect compute default values and export
func (mgr *Mgr) Collect() {
	mgr.RLock()
	defer mgr.RUnlock()
	for _, tks := range mgr.ticks {
		tks.collect()
	}
}

// ServeHTTP responds to an HTTP request
func (mgr *Mgr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mgr.handler.ServeHTTP(w, r)
}
