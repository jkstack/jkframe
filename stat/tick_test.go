package stat

import (
	"math/rand"
	"testing"
	"time"
)

func TestTicks(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tks := newTicks("test")
	for i := 0; i < 10; i++ {
		tk := &Tick{
			begin: time.Now(),
			end:   time.Now(),
		}
		tks.push(tk)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		tk.Close()
	}
	tks.collect()
}
