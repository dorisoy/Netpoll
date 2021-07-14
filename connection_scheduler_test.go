// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netpoll

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestConnectionScheduler(t *testing.T) {
	var scher scheduler
	var done int32
	gsize := 10
	for i := 0; i < gsize; i++ {
		scher.add()
		go func() {
			defer scher.done()
			time.Sleep(time.Millisecond * 10)
			atomic.AddInt32(&done, 1)
		}()
	}
	scher.pause()
	Equal(t, done, int32(gsize))
	scher.resume()
}
