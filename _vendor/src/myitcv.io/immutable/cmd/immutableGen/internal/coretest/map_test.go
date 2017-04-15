package coretest_test

import (
	"testing"

	"myitcv.io/immutable/cmd/immutableGen/internal/coretest"
)

func TestMyMapZeroValue(t *testing.T) {
	// a map cannot be used as a zero value
}

func TestMyMapAsMutableImmutableReceiver(t *testing.T) {
	s1 := coretest.NewMyMap()
	s2 := s1.AsMutable()

	if s1 == s2 {
		t.Fatalf("s1 and s2 should be different values; they are not")
	}

	if s1.Mutable() {
		t.Fatalf("s1 should not be mutable; it is")
	}

	if !s2.Mutable() {
		t.Fatalf("s2 should be mutable; it is not")
	}
}

func TestMyMapAsMutableMutableReceiver(t *testing.T) {
	s1 := coretest.NewMyMap().AsMutable()
	s2 := s1.AsMutable()

	if s1 != s2 {
		t.Fatalf("s1 and s2 should not be different values; they are")
	}

	if !s2.Mutable() {
		t.Fatalf("s2 should be mutable; it is not")
	}
}

func TestMyMapWithMutableImmutableReceiver(t *testing.T) {
	wasMutable := false

	var s3 *coretest.MyMap

	s1 := coretest.NewMyMap()
	s2 := s1.WithMutable(func(s *coretest.MyMap) {
		wasMutable = s.Mutable()

		// have some side effect
		s.Set(peter, age42)

		s3 = s
	})

	if s1 == s2 {
		t.Fatalf("s1 and s2 should be different values; they are not")
	}

	if s3 != s2 {
		t.Fatalf("s3 and s2 should be same values; they were not")
	}

	if !wasMutable {
		t.Fatalf("s should have been mutable; it was not")
	}

	if s2.Mutable() {
		t.Fatalf("s2 should not be mutable")
	}
}

func TestMyMapWithMutableMutableReceiver(t *testing.T) {
	p := "Peter"
	a := 42
	wasMutable := false

	var s3 *coretest.MyMap

	s1 := coretest.NewMyMap().AsMutable()
	s2 := s1.WithMutable(func(s *coretest.MyMap) {
		wasMutable = s.Mutable()

		// have some side effect
		s.Set(p, a)

		s3 = s
	})

	if s1 != s2 {
		t.Fatalf("s1 and s2 should be same values; they were not")
	}

	if s3 != s2 {
		t.Fatalf("s3 and s2 should be same values; they were not")
	}

	if !wasMutable {
		t.Fatalf("s should have been mutable; it was not")
	}

	if !s2.Mutable() {
		t.Fatalf("s2 should be mutable")
	}
}

func TestMyMapWithImmutableImmutableReceiver(t *testing.T) {
	wasMutable := false

	var s3 *coretest.MyMap

	s1 := coretest.NewMyMap()
	s2 := s1.WithImmutable(func(s *coretest.MyMap) {
		wasMutable = s.Mutable()

		// have some side effect
		s.Set(peter, age42)

		s3 = s
	})

	if s1 != s2 {
		t.Fatalf("s1 and s2 should be sames values; they are not")
	}

	if s3 != s2 {
		t.Fatalf("s3 and s2 should be same values; they were not")
	}

	if wasMutable {
		t.Fatalf("s should not have been mutable; it was")
	}

	if s2.Mutable() {
		t.Fatalf("s2 should not be mutable")
	}
}

func TestMyMapWithImmutableMutableReceiver(t *testing.T) {
	wasMutable := false

	var s3 *coretest.MyMap

	s1 := coretest.NewMyMap().AsMutable()
	s2 := s1.WithImmutable(func(s *coretest.MyMap) {
		wasMutable = s.Mutable()

		// have some side effect
		s.Set(peter, age42)

		s3 = s
	})

	if s1 != s2 {
		t.Fatalf("s1 and s2 should be same values; they were not")
	}

	if s3 != s2 {
		t.Fatalf("s3 and s2 should be same values; they were not")
	}

	if wasMutable {
		t.Fatalf("s should not have been mutable; it was")
	}

	if !s2.Mutable() {
		t.Fatalf("s2 should be mutable")
	}
}

func TestMyMapAsImmutableImmutableReceiver(t *testing.T) {

	v := func() *coretest.MyMap {
		return coretest.NewMyMap()
	}

	// self
	{
		s1 := v()
		s2 := s1.AsImmutable(s1)

		if s2 != s1 {
			t.Fatalf("expected s1 and s2 to be the same value")
		}

		if s2.Mutable() {
			t.Fatalf("s2 should not be mutable")
		}
	}

	// nil
	{
		s1 := v()
		s2 := s1.AsImmutable(nil)

		if s2 != s1 {
			t.Fatalf("expected s1 and s2 to be the same value")
		}

		if s2.Mutable() {
			t.Fatalf("s2 should not be mutable")
		}
	}

	// different value
	{
		s1 := v()
		s2 := s1.AsImmutable(coretest.NewMyMap())

		if s2 != s1 {
			t.Fatalf("expected s1 and s2 to be the same value")
		}

		if s2.Mutable() {
			t.Fatalf("s2 should not be mutable")
		}
	}
}

func TestMyMapAsImmutableMutableReceiver(t *testing.T) {

	v := func() *coretest.MyMap {
		return coretest.NewMyMap().AsMutable()
	}

	// self
	{
		s1 := v()
		s2 := s1.AsImmutable(s1)

		if s2 != s1 {
			t.Fatalf("expected s1 and s2 to be the same value")
		}

		if !s2.Mutable() {
			t.Fatalf("s2 should be mutable")
		}
	}

	// nil
	{
		s1 := v()
		s2 := s1.AsImmutable(nil)

		if s2 != s1 {
			t.Fatalf("expected s1 and s2 to be the same value")
		}

		if s2.Mutable() {
			t.Fatalf("s2 should not be mutable")
		}
	}

	// different value
	{
		s1 := v()
		s2 := s1.AsImmutable(coretest.NewMyMap())

		if s2 != s1 {
			t.Fatalf("expected s1 and s2 to be the same value")
		}

		if s2.Mutable() {
			t.Fatalf("s2 should not be mutable")
		}
	}
}

func TestMyMapPlainConstructor(t *testing.T) {
	s1 := coretest.NewMyMap()

	if v := s1.Len(); v != 0 {
		t.Fatalf("length of s1 should be 0; got %v", v)
	}

	if s1.Mutable() {
		t.Fatalf("s1 should not be mutable; it is")
	}

	if v, ok := s1.Get(paul); ok {
		t.Fatalf("should be nothing in the map; found %q with value %v", paul, v)
	}
}

func TestMyMapPlainConstructorWithInitialiser(t *testing.T) {
	s1 := coretest.NewMyMap(func(m *coretest.MyMap) {
		m.Set(peter, age42)
	})

	if v := s1.Len(); v != 1 {
		t.Fatalf("length of s1 should be 1; got %v", v)
	}

	if s1.Mutable() {
		t.Fatalf("s1 should not be mutable; it is")
	}

	if v, ok := s1.Get(peter); !ok || v != age42 {
		t.Fatalf("s1.Get(%q) should be (%v, %v); got (%v, %v)", peter, age42, true, v, ok)
	}
}

func TestMyMapConstructorCapacity(t *testing.T) {
	s1 := coretest.NewMyMapCap(2)

	if v := s1.Len(); v != 0 {
		t.Fatalf("length of s1 should be 0; got %v", v)
	}

	if s1.Mutable() {
		t.Fatalf("s1 should not be mutable; it is")
	}

	if v, ok := s1.Get(paul); ok {
		t.Fatalf("should be nothing in the map; found %q with value %v", paul, v)
	}
}

func TestMyMapSetImmutableReceiver(t *testing.T) {
	s1 := coretest.NewMyMap()
	s2 := s1.Set(peter, age42)

	if s1 == s2 {
		t.Fatalf("s1 and s2 should be different values; they are not")
	}

	if v, ok := s1.Get(peter); ok {
		t.Fatalf("expected s1 to not contain %q; but it did with value %v", peter, v)
	}

	if v, ok := s2.Get(peter); !ok || v != age42 {
		t.Fatalf("expected s2.Get(%q) to be (%v, %v); got (%v, %v)", peter, age42, true, v, ok)
	}

	if s2.Mutable() {
		t.Fatalf("s2 should not be mutable")
	}
}

func TestMyMapSetMutableReceiver(t *testing.T) {
	s1 := coretest.NewMyMap().AsMutable()
	s2 := s1.Set(peter, age42)

	if s1 != s2 {
		t.Fatalf("s1 and s2 should be sames values; they are not")
	}

	if v, ok := s2.Get(peter); !ok || v != age42 {
		t.Fatalf("expected s2.Get(%q) to be (%v, %v); got (%v, %v)", peter, age42, true, v, ok)
	}

	if !s2.Mutable() {
		t.Fatalf("s2 should be mutable")
	}
}
