package netpoll

import (
	"testing"
	"time"
)

func TestConnectionScheduler(t *testing.T) {
	var scher scheduler
	done := 0
	gsize := 10
	for i := 0; i < gsize; i++ {
		scher.add()
		go func() {
			defer scher.done()
			time.Sleep(time.Millisecond * 10)
			done++
		}()
	}
	scher.pause()
	Assert(t, done == gsize, done, gsize)
	scher.resume()
}
