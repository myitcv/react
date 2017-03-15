package coretest

import (
	"reflect"
	"testing"
	"unsafe"
)

func slicesEqual(a, b []string) bool {
	return *(*reflect.SliceHeader)(unsafe.Pointer(&a)) == *(*reflect.SliceHeader)(unsafe.Pointer(&b))
}

func mapsEqual(a, b map[string]int) bool {
	return *(*uintptr)(unsafe.Pointer(&a)) == *(*uintptr)(unsafe.Pointer(&b))
}

func TestSliceToSliceCopies(t *testing.T) {
	s1 := new(MySlice)

	if s1c := s1.ToSlice(); s1c != nil {
		t.Fatalf("expected s1.ToSlice() to be nil; got %v", s1c)
	}

	s2 := NewMySlice("a", "b")
	s2c := s2.ToSlice()

	if slicesEqual(s2.theSlice, s2c) {
		t.Fatalf("expected ToSlice to return a copy; it did not")
	}

	if v := len(s2c); v != 2 {
		t.Fatalf("expected copy to have 2 elements; had %v", v)
	}

	for i := range s2c {
		orig := s2.Get(i)
		cpy := s2c[i]

		if orig != cpy {
			t.Fatalf("expected %q == %q", orig, cpy)
		}
	}
}

func TestMapToMapCopies(t *testing.T) {
	m1 := NewMyMap()

	m1c := m1.ToMap()

	if m1c == nil {
		t.Fatalf("expected ToMap of empty map not to be nil")
	}

	if len(m1c) != 0 {
		t.Fatalf("expected ToMap of empty map to not be empty")
	}

	m2 := NewMyMap(func(m *MyMap) {
		m.Set("a", 1)
		m.Set("b", 2)
	})

	m2c := m2.ToMap()

	if mapsEqual(m2.theMap, m2c) {
		t.Fatalf("expected ToMap to make a copy; it did not")
	}

	if m2c == nil {
		t.Fatalf("expected ToMap of non-empty map not to be nil")
	}

	if v := len(m2c); v != 2 {
		t.Fatalf("expected ToMap map to have 2 values; had %v", v)
	}

	for k := range m2c {
		orig, ok := m2.Get(k)
		cpy := m2c[k]

		if !ok {
			t.Fatalf("expected to find %v in orig", k)
		}

		if orig != cpy {
			t.Fatalf("expected %v == %v", orig, cpy)
		}
	}
}
