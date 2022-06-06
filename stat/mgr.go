package stat

import (
	"net/http"
	"sync"
	"time"

	"github.com/jkstack/jkframe/logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Mgr tick management
type Mgr struct {
	sync.RWMutex
	handler http.Handler
	data    map[string]*ticks
}

// New create management
func New(interval time.Duration) *Mgr {
	mgr := &Mgr{
		handler: promhttp.Handler(),
		data:    make(map[string]*ticks),
	}
	go func() {
		for {
			mgr.Collect()
			time.Sleep(interval)
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

// ServeHTTP responds to an HTTP request
func (mgr *Mgr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logging.Info("metrics tick")
	mgr.handler.ServeHTTP(w, r)
}
