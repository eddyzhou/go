// Copyright 2015 Reborndb Org. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package atomic2

import (
	"sync"
	"sync/atomic"
)

type Int64 struct {
	v, s int64
}

func (a *Int64) Get() int64 {
	return atomic.LoadInt64(&a.v)
}

func (a *Int64) Set(v int64) {
	atomic.StoreInt64(&a.v, v)
}

func (a *Int64) Reset() int64 {
	return atomic.SwapInt64(&a.v, 0)
}

func (a *Int64) Add(v int64) int64 {
	return atomic.AddInt64(&a.v, v)
}

func (a *Int64) Sub(v int64) int64 {
	return a.Add(-v)
}

func (a *Int64) Snapshot() {
	a.s = a.Get()
}

func (a *Int64) Delta() int64 {
	return a.Get() - a.s
}

func (a *Int64) Incr() int64 {
	return a.Add(1)
}

func (a *Int64) Decr() int64 {
	return a.Add(-1)
}

func (a *Int64) CompareAndSwap(oldval, newval int64) (swapped bool) {
	return atomic.CompareAndSwapInt64((*int64)(&a.v), oldval, newval)
}

type String struct {
	mu  sync.Mutex
	str string
}

func (s *String) Set(str string) {
	s.mu.Lock()
	s.str = str
	s.mu.Unlock()
}

func (s *String) Get() string {
	s.mu.Lock()
	str := s.str
	s.mu.Unlock()
	return str
}

func (s *String) CompareAndSwap(oldval, newval string) (swqpped bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.str == oldval {
		s.str = newval
		return true
	}
	return false
}
