package mapSet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type threadUnsafeSet map[interface{}]struct{}

func newThreadUnsafeSet() threadUnsafeSet {
	return make(threadUnsafeSet)
}

func (set *threadUnsafeSet) Add(i interface{}) bool {
	_, found := (*set)[i]
	if found {
		return false //False if it existed already
	}

	(*set)[i] = struct{}{}
	return true
}

func (set *threadUnsafeSet) Contains(i ...interface{}) bool {
	for _, val := range i {
		if _, ok := (*set)[val]; !ok {
			return false
		}
	}
	return true
}

func (set *threadUnsafeSet) Clear() {
	*set = newThreadUnsafeSet()
}

func (set *threadUnsafeSet) Remove(i interface{})  {
	delete(*set, i)
}

func (set *threadUnsafeSet) RetElementCount() int {
	return len(*set)
}

func (set *threadUnsafeSet) Each(cb func(interface{}) bool) {
	for elem := range *set {
		if cb(elem) {
			break
		}
	}
}

func (set *threadUnsafeSet) Pop() interface{} {
	for item := range *set {
		delete(*set, item)
		return item
	}
	return nil
}

func (set *threadUnsafeSet) Equal(other MapSet) bool {
	_ = other.(*threadUnsafeSet)

	if set.RetElementCount() != other.RetElementCount() {
		return false
	}
	for elem := range *set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (set *threadUnsafeSet) Clone() MapSet {
	clonedSet := newThreadUnsafeSet()
	for elem := range *set {
		clonedSet.Add(elem)
	}
	return &clonedSet
}


func (set *threadUnsafeSet) String(sep string) string {
	items := make([]string, 0, len(*set))

	for elem := range *set {
		items = append(items, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, sep))
}

func (set *threadUnsafeSet) RandomReturn() interface{} {
	for item := range *set {
		return item
	}
	return nil
}

func (set *threadUnsafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, set.RetElementCount())
	for elem := range *set {
		keys = append(keys, elem)
	}

	return keys
}

// MarshalJSON creates a JSON array from the set, it marshals all elements
func (set *threadUnsafeSet) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, set.RetElementCount())

	for elem := range *set {
		b, err := json.Marshal(elem)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ","))), nil
}

// UnmarshalJSON recreates a set from a JSON array, it only decodes
// primitive types. Numbers are decoded as json.Number.
func (set *threadUnsafeSet) UnmarshalJSON(b []byte) error {
	var i []interface{}

	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err := d.Decode(&i)
	if err != nil {
		return err
	}

	for _, v := range i {
		switch t := v.(type) {
		case []interface{}, map[string]interface{}:
			continue
		default:
			set.Add(t)
		}
	}

	return nil
}
