package mapSet

import (
	"sync"
)

type threadSafeSet struct {
	s threadUnsafeSet
	sync.RWMutex
}

func newThreadSafeSet() threadSafeSet {
	return threadSafeSet{s: newThreadUnsafeSet()}
}

func (set *threadSafeSet) Add(i interface{}) bool {
	set.Lock()
	ret := set.s.Add(i)
	set.Unlock()
	return ret
}

func (set *threadSafeSet) Contains(i ...interface{}) bool {
	set.RLock()
	ret := set.s.Contains(i...)
	set.RUnlock()
	return ret
}

func (set *threadSafeSet) Clear() {
	set.Lock()
	set.s = newThreadUnsafeSet()
	set.Unlock()
}

func (set *threadSafeSet) Remove(i interface{}) {
	set.Lock()
	delete(set.s, i)
	set.Unlock()
}

func (set *threadSafeSet) RetElementCount() int {
	set.RLock()
	defer set.RUnlock()
	return len(set.s)
}

func (set *threadSafeSet) Each(cb func(interface{}) bool) {
	set.RLock()
	for elem := range set.s {
		if cb(elem) {
			break
		}
	}
	set.RUnlock()
}

func (set *threadSafeSet) Equal(other MapSet) bool {
	o := other.(*threadSafeSet)

	set.RLock()
	o.RLock()

	ret := set.s.Equal(&o.s)
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *threadSafeSet) Clone() MapSet {
	set.RLock()

	unsafeClone := set.s.Clone().(*threadUnsafeSet)
	ret := &threadSafeSet{s: *unsafeClone}
	set.RUnlock()
	return ret
}

func (set *threadSafeSet) Pop() interface{} {
	set.Lock()
	defer set.Unlock()
	return set.s.Pop()
}


func (set *threadSafeSet) String(sep string) string {
	set.RLock()
	ret := set.s.String(sep)
	set.RUnlock()
	return ret
}

func (set *threadSafeSet) RandomReturn() interface{} {
	set.Lock()
	defer set.Unlock()
	return set.s.RandomReturn()
}

func (set *threadSafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, set.RetElementCount())
	set.RLock()
	for elem := range set.s {
		keys = append(keys, elem)
	}
	set.RUnlock()
	return keys
}
