package mapSet

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

const N = 1000

func Test_AddConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewMapSet()
	ints := rand.Perm(N)

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for i := 0; i < len(ints); i++ {
		go func(i int) {
			s.Add(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	for _, i := range ints {
		if !s.Contains(i) {
			t.Errorf("Set is missing element: %v", i)
		}
	}
}

func Test_CardinalityConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewMapSet()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		elems := s.RetElementCount()
		for i := 0; i < N; i++ {
			newElems := s.RetElementCount()
			if newElems < elems {
				t.Errorf("Cardinality shrunk from %v to %v", elems, newElems)
			}
		}
		wg.Done()
	}()

	for i := 0; i < N; i++ {
		s.Add(rand.Int())
	}
	wg.Wait()
}

func Test_ClearConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewMapSet()
	ints := rand.Perm(N)

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for i := 0; i < len(ints); i++ {
		go func() {
			s.Clear()
			wg.Done()
		}()
		go func(i int) {
			s.Add(i)
		}(i)
	}

	wg.Wait()
}

func Test_CloneConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewMapSet()
	ints := rand.Perm(N)

	for _, v := range ints {
		s.Add(v)
	}

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for i := range ints {
		go func(i int) {
			s.Remove(i)
			wg.Done()
		}(i)
	}

	s.Clone()
}

func Test_ContainsConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewMapSet()
	ints := rand.Perm(N)
	interfaces := make([]interface{}, 0)
	for _, v := range ints {
		s.Add(v)
		interfaces = append(interfaces, v)
	}

	var wg sync.WaitGroup
	for range ints {
		wg.Add(1)
		go func() {
			s.Contains(interfaces...)
			wg.Done()
		}()
	}
	wg.Wait()
}


func Test_EqualConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s, ss := NewMapSet(), NewMapSet()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(v)
		ss.Add(v)
	}

	var wg sync.WaitGroup
	for range ints {
		wg.Add(1)
		go func() {
			s.Equal(ss)
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_EachConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)
	concurrent := 10

	s := NewMapSet()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(v)
	}

	var count int64
	wg := new(sync.WaitGroup)
	wg.Add(concurrent)
	for n := 0; n < concurrent; n++ {
		go func() {
			defer wg.Done()
			s.Each(func(elem interface{}) bool {
				atomic.AddInt64(&count, 1)
				return false
			})
		}()
	}
	wg.Wait()

	if count != int64(N*concurrent) {
		t.Errorf("%v != %v", count, int64(N*concurrent))
	}
}

func Test_RemoveConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewMapSet()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(v)
	}

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for _, v := range ints {
		go func(i int) {
			s.Remove(i)
			wg.Done()
		}(v)
	}
	wg.Wait()

	if s.RetElementCount() != 0 {
		t.Errorf("Expected cardinality 0; got %v", s.RetElementCount())
	}
}

func Test_StringConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewMapSet()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(v)
	}

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for range ints {
		go func() {
			_ = s.String(",")
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_ToSlice(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewMapSet()
	ints := rand.Perm(N)

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for i := 0; i < len(ints); i++ {
		go func(i int) {
			s.Add(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	setAsSlice := s.ToSlice()
	if len(setAsSlice) != s.RetElementCount() {
		t.Errorf("Set length is incorrect: %v", len(setAsSlice))
	}

	for _, i := range setAsSlice {
		if !s.Contains(i) {
			t.Errorf("Set is missing element: %v", i)
		}
	}
}

// Test_ToSliceDeadlock - fixes issue: https://github.com/deckarep/golang-set/issues/36
// This code reveals the deadlock however it doesn't happen consistently.
func Test_ToSliceDeadlock(t *testing.T) {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	set := NewMapSet()
	workers := 10
	wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				set.Add(1)
				set.ToSlice()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
