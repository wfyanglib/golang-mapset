package mapSet

import "testing"

func makeSet(ints []int) MapSet {
	set := NewMapSet()
	for _, i := range ints {
		set.Add(i)
	}
	return set
}

func makeUnsafeSet(ints []int) MapSet {
	set := NewThreadUnsafeSet()
	for _, i := range ints {
		set.Add(i)
	}
	return set
}

func assertEqual(a, b MapSet, t *testing.T) {
	if !a.Equal(b) {
		t.Errorf("%v != %v\n", a, b)
	}
}

func Test_NewSet(t *testing.T) {
	a := NewMapSet()
	if a.RetElementCount() != 0 {
		t.Error("NewSet should start out as an empty set")
	}

	assertEqual(NewSetFromSlice([]interface{}{}), NewMapSet(), t)
	assertEqual(NewSetFromSlice([]interface{}{1}), NewMapSet(1), t)
	assertEqual(NewSetFromSlice([]interface{}{1, 2}), NewMapSet(1, 2), t)
	assertEqual(NewSetFromSlice([]interface{}{"a"}), NewMapSet("a"), t)
	assertEqual(NewSetFromSlice([]interface{}{"a", "b"}), NewMapSet("a", "b"), t)
}

func Test_NewUnsafeSet(t *testing.T) {
	a := NewThreadUnsafeSet()

	if a.RetElementCount() != 0 {
		t.Error("NewSet should start out as an empty set")
	}
}

func Test_AddSet(t *testing.T) {
	a := makeSet([]int{1, 2, 3})

	if a.RetElementCount() != 3 {
		t.Error("AddSet does not have a size of 3 even though 3 items were added to a new set")
	}
}

func Test_AddUnsafeSet(t *testing.T) {
	a := makeUnsafeSet([]int{1, 2, 3})

	if a.RetElementCount() != 3 {
		t.Error("AddSet does not have a size of 3 even though 3 items were added to a new set")
	}
}

func Test_AddSetNoDuplicate(t *testing.T) {
	a := makeSet([]int{7, 5, 3, 7})

	if a.RetElementCount() != 3 {
		t.Error("AddSetNoDuplicate set should have 3 elements since 7 is a duplicate")
	}

	if !(a.Contains(7) && a.Contains(5) && a.Contains(3)) {
		t.Error("AddSetNoDuplicate set should have a 7, 5, and 3 in it.")
	}
}

func Test_AddUnsafeSetNoDuplicate(t *testing.T) {
	a := makeUnsafeSet([]int{7, 5, 3, 7})

	if a.RetElementCount() != 3 {
		t.Error("AddSetNoDuplicate set should have 3 elements since 7 is a duplicate")
	}

	if !(a.Contains(7) && a.Contains(5) && a.Contains(3)) {
		t.Error("AddSetNoDuplicate set should have a 7, 5, and 3 in it.")
	}
}

func Test_RemoveSet(t *testing.T) {
	a := makeSet([]int{6, 3, 1})

	a.Remove(3)

	if a.RetElementCount() != 2 {
		t.Error("RemoveSet should only have 2 items in the set")
	}

	if !(a.Contains(6) && a.Contains(1)) {
		t.Error("RemoveSet should have only items 6 and 1 in the set")
	}

	a.Remove(6)
	a.Remove(1)

	if a.RetElementCount() != 0 {
		t.Error("RemoveSet should be an empty set after removing 6 and 1")
	}
}

func Test_RemoveUnsafeSet(t *testing.T) {
	a := makeUnsafeSet([]int{6, 3, 1})

	a.Remove(3)

	if a.RetElementCount() != 2 {
		t.Error("RemoveSet should only have 2 items in the set")
	}

	if !(a.Contains(6) && a.Contains(1)) {
		t.Error("RemoveSet should have only items 6 and 1 in the set")
	}

	a.Remove(6)
	a.Remove(1)

	if a.RetElementCount() != 0 {
		t.Error("RemoveSet should be an empty set after removing 6 and 1")
	}
}

func Test_ContainsSet(t *testing.T) {
	a := NewMapSet()

	a.Add(71)

	if !a.Contains(71) {
		t.Error("ContainsSet should contain 71")
	}

	a.Remove(71)

	if a.Contains(71) {
		t.Error("ContainsSet should not contain 71")
	}

	a.Add(13)
	a.Add(7)
	a.Add(1)

	if !(a.Contains(13) && a.Contains(7) && a.Contains(1)) {
		t.Error("ContainsSet should contain 13, 7, 1")
	}
}

func Test_ContainsUnsafeSet(t *testing.T) {
	a := NewThreadUnsafeSet()

	a.Add(71)

	if !a.Contains(71) {
		t.Error("ContainsSet should contain 71")
	}

	a.Remove(71)

	if a.Contains(71) {
		t.Error("ContainsSet should not contain 71")
	}

	a.Add(13)
	a.Add(7)
	a.Add(1)

	if !(a.Contains(13) && a.Contains(7) && a.Contains(1)) {
		t.Error("ContainsSet should contain 13, 7, 1")
	}
}

func Test_ContainsMultipleSet(t *testing.T) {
	a := makeSet([]int{8, 6, 7, 5, 3, 0, 9})

	if !a.Contains(8, 6, 7, 5, 3, 0, 9) {
		t.Error("ContainsAll should contain Jenny's phone number")
	}

	if a.Contains(8, 6, 11, 5, 3, 0, 9) {
		t.Error("ContainsAll should not have all of these numbers")
	}
}

func Test_ContainsMultipleUnsafeSet(t *testing.T) {
	a := makeUnsafeSet([]int{8, 6, 7, 5, 3, 0, 9})

	if !a.Contains(8, 6, 7, 5, 3, 0, 9) {
		t.Error("ContainsAll should contain Jenny's phone number")
	}

	if a.Contains(8, 6, 11, 5, 3, 0, 9) {
		t.Error("ContainsAll should not have all of these numbers")
	}
}

func Test_ClearSet(t *testing.T) {
	a := makeSet([]int{2, 5, 9, 10})

	a.Clear()

	if a.RetElementCount() != 0 {
		t.Error("ClearSet should be an empty set")
	}
}

func Test_ClearUnsafeSet(t *testing.T) {
	a := makeUnsafeSet([]int{2, 5, 9, 10})

	a.Clear()

	if a.RetElementCount() != 0 {
		t.Error("ClearSet should be an empty set")
	}
}

func Test_CardinalitySet(t *testing.T) {
	a := NewMapSet()

	if a.RetElementCount() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(1)

	if a.RetElementCount() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Remove(1)

	if a.RetElementCount() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(9)

	if a.RetElementCount() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Clear()

	if a.RetElementCount() != 0 {
		t.Error("set should have a size of 1")
	}
}

func Test_CardinalityUnsafeSet(t *testing.T) {
	a := NewThreadUnsafeSet()

	if a.RetElementCount() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(1)

	if a.RetElementCount() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Remove(1)

	if a.RetElementCount() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(9)

	if a.RetElementCount() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Clear()

	if a.RetElementCount() != 0 {
		t.Error("set should have a size of 1")
	}
}

func Test_SetEqual(t *testing.T) {
	a := NewMapSet()
	b := NewMapSet()

	if !a.Equal(b) {
		t.Error("Both a and b are empty sets, and should be equal")
	}

	a.Add(10)

	if a.Equal(b) {
		t.Error("a should not be equal to b because b is empty and a has item 1 in it")
	}

	b.Add(10)

	if !a.Equal(b) {
		t.Error("a is now equal again to b because both have the item 10 in them")
	}

	b.Add(8)
	b.Add(3)
	b.Add(47)

	if a.Equal(b) {
		t.Error("b has 3 more elements in it so therefore should not be equal to a")
	}

	a.Add(8)
	a.Add(3)
	a.Add(47)

	if !a.Equal(b) {
		t.Error("a and b should be equal with the same number of elements")
	}
}

func Test_UnsafeSetEqual(t *testing.T) {
	a := NewThreadUnsafeSet()
	b := NewThreadUnsafeSet()

	if !a.Equal(b) {
		t.Error("Both a and b are empty sets, and should be equal")
	}

	a.Add(10)

	if a.Equal(b) {
		t.Error("a should not be equal to b because b is empty and a has item 1 in it")
	}

	b.Add(10)

	if !a.Equal(b) {
		t.Error("a is now equal again to b because both have the item 10 in them")
	}

	b.Add(8)
	b.Add(3)
	b.Add(47)

	if a.Equal(b) {
		t.Error("b has 3 more elements in it so therefore should not be equal to a")
	}

	a.Add(8)
	a.Add(3)
	a.Add(47)

	if !a.Equal(b) {
		t.Error("a and b should be equal with the same number of elements")
	}
}

func Test_SetClone(t *testing.T) {
	a := NewMapSet()
	a.Add(1)
	a.Add(2)

	b := a.Clone()

	if !a.Equal(b) {
		t.Error("Clones should be equal")
	}

	a.Add(3)
	if a.Equal(b) {
		t.Error("a contains one more element, they should not be equal")
	}

	c := a.Clone()
	c.Remove(1)

	if a.Equal(c) {
		t.Error("C contains one element less, they should not be equal")
	}
}

func Test_UnsafeSetClone(t *testing.T) {
	a := NewThreadUnsafeSet()
	a.Add(1)
	a.Add(2)

	b := a.Clone()

	if !a.Equal(b) {
		t.Error("Clones should be equal")
	}

	a.Add(3)
	if a.Equal(b) {
		t.Error("a contains one more element, they should not be equal")
	}

	c := a.Clone()
	c.Remove(1)

	if a.Equal(c) {
		t.Error("C contains one element less, they should not be equal")
	}
}

func Test_Each(t *testing.T) {
	a := NewMapSet()

	a.Add("Z")
	a.Add("Y")
	a.Add("X")
	a.Add("W")

	b := NewMapSet()
	a.Each(func(elem interface{}) bool {
		b.Add(elem)
		return false
	})

	if !a.Equal(b) {
		t.Error("The sets are not equal after iterating (Each) through the first set")
	}

	var count int
	a.Each(func(elem interface{}) bool {
		if count == 2 {
			return true
		}
		count++
		return false
	})
	if count != 2 {
		t.Error("Iteration should stop on the way")
	}
}

func Test_PopSafe(t *testing.T) {
	a := NewMapSet()

	a.Add("a")
	a.Add("b")
	a.Add("c")
	a.Add("d")

	captureSet := NewMapSet()
	captureSet.Add(a.Pop())
	captureSet.Add(a.Pop())
	captureSet.Add(a.Pop())
	captureSet.Add(a.Pop())
	finalNil := a.Pop()

	if captureSet.RetElementCount() != 4 {
		t.Error("unexpected captureSet cardinality; should be 4")
	}

	if a.RetElementCount() != 0 {
		t.Error("unepxected a cardinality; should be zero")
	}

	if !captureSet.Contains("c", "a", "d", "b") {
		t.Error("unexpected result set; should be a,b,c,d (any order is fine")
	}

	if finalNil != nil {
		t.Error("when original set is empty, further pops should result in nil")
	}
}

func Test_PopUnsafe(t *testing.T) {
	a := NewThreadUnsafeSet()

	a.Add("a")
	a.Add("b")
	a.Add("c")
	a.Add("d")

	captureSet := NewThreadUnsafeSet()
	captureSet.Add(a.Pop())
	captureSet.Add(a.Pop())
	captureSet.Add(a.Pop())
	captureSet.Add(a.Pop())
	finalNil := a.Pop()

	if captureSet.RetElementCount() != 4 {
		t.Error("unexpected captureSet cardinality; should be 4")
	}

	if a.RetElementCount() != 0 {
		t.Error("unepxected a cardinality; should be zero")
	}

	if !captureSet.Contains("c", "a", "d", "b") {
		t.Error("unexpected result set; should be a,b,c,d (any order is fine")
	}

	if finalNil != nil {
		t.Error("when original set is empty, further pops should result in nil")
	}
}

func Test_ToSliceUnthreadsafe(t *testing.T) {
	s := makeUnsafeSet([]int{1, 2, 3})
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