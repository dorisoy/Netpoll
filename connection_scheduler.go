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
	"math"
	"runtime"
	"sync/atomic"
)

const PAUSED = math.MinInt32

type scheduler struct {
	running int32
}

func (s *scheduler) isEmpty() bool {
	return atomic.LoadInt32(&s.running) == 0 || atomic.LoadInt32(&s.running) == PAUSED
}

func (s *scheduler) add() {
	for {
		x := atomic.AddInt32(&s.running, 1)
		if x > 0 {
			return
		}
		//reset to pause
		atomic.CompareAndSwapInt32(&s.running, x, PAUSED)
		//in pause
		runtime.Gosched()
	}
}

func (s *scheduler) done() {
	atomic.AddInt32(&s.running, -1)
}

func (s *scheduler) pause() {
	for !atomic.CompareAndSwapInt32(&s.running, 0, PAUSED) {
		runtime.Gosched()
	}
}

func (s *scheduler) resume() {
	atomic.StoreInt32(&s.running, 0)
}
