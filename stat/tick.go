package stat

import (
	"container/list"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Tick tick object
type Tick struct {
	begin time.Time
	end   time.Time
}

func (tk *Tick) Close() {
	tk.end = time.Now()
}

type ticks struct {
	sync.Mutex
	list.List
	name string
	vec  *prometheus.GaugeVec
}

func newTicks(name string) *ticks {
	list := &ticks{
		name: name,
		vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: name,
		}, []string{"tag"}),
	}
	prometheus.MustRegister(list.vec)
	return list
}

func (tks *ticks) push(tick *Tick) {
	tks.Lock()
	defer tks.Unlock()
	tks.PushBack(tick)
}

type collectResult struct {
	qps                     float64
	avg                     float64
	p0, p50, p90, p99, p100 int64
}

func (tks *ticks) collect() {
	elements := make([]int64, 0, tks.Len())

	var next *list.Element
	var begin, end time.Time
	tks.Lock()
	for element := tks.Front(); element != nil; element = next {
		tk := element.Value.(*Tick)
		if begin.IsZero() {
			begin = tk.begin
		}
		if end.IsZero() {
			end = tk.end
		}
		if tk.begin.Before(begin) {
			begin = tk.begin
		}
		if tk.end.After(end) {
			end = tk.end
		}
		elements = append(elements, tk.end.Sub(tk.begin).Microseconds())
		next = element.Next()
		tks.Remove(element)
	}
	tks.Unlock()

	if end.Sub(begin).Seconds() < 1 {
		end = begin.Add(time.Second)
	}

	var result collectResult
	if len(elements) > 0 {
		result = collect(begin, end, elements)
	}

	tks.vec.With(prometheus.Labels{"tag": "qps"}).Set(result.qps)
	tks.vec.With(prometheus.Labels{"tag": "avg"}).Set(result.avg)
	tks.vec.With(prometheus.Labels{"tag": "p0"}).Set(float64(result.p0))
	tks.vec.With(prometheus.Labels{"tag": "p50"}).Set(float64(result.p50))
	tks.vec.With(prometheus.Labels{"tag": "p90"}).Set(float64(result.p90))
	tks.vec.With(prometheus.Labels{"tag": "p99"}).Set(float64(result.p99))
	tks.vec.With(prometheus.Labels{"tag": "p100"}).Set(float64(result.p100))
}

func collect(begin, end time.Time, elements []int64) collectResult {
	qps := float64(len(elements)) / end.Sub(begin).Seconds()
	avg := sum(elements) / float64(len(elements))
	sort.Slice(elements, func(i, j int) bool {
		return elements[i] < elements[j]
	})
	p50 := len(elements) >> 1
	p90 := int(math.Floor(float64(len(elements)) * 9 / 10))
	p99 := int(math.Floor(float64(len(elements)) * 99 / 100))

	return collectResult{
		qps:  qps,
		avg:  avg,
		p0:   elements[0],
		p50:  elements[p50],
		p90:  elements[p90],
		p99:  elements[p99],
		p100: elements[len(elements)-1],
	}
}

func sum(elements []int64) float64 {
	var ret float64
	for _, e := range elements {
		ret += float64(e)
	}
	return ret
}
