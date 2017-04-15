package coretest_test

import (
	"testing"

	"myitcv.io/immutable/cmd/immutableGen/internal/coretest"
)

func TestMySliceZeroValue(t *testing.T) {
	s1 := new(coretest.MySlice)

	if s1.Mutable() {
		t.Fatalf("zero value should be immutable")
	}

	if s1.Len() != 0 {
		t.Fatalf("zero value should have zero length")
	}

	setFailed := false
	func() {
		defer func() {
			if _, ok := recover().(error); ok {
				setFailed = true
			}
		}()

		s1.Set(0, "test")
	}()

	if !setFailed {
		t.Fatalf("should panic when setting on zero value")
	}

	getFailed := false
	func() {
		defer func() {
			if _, ok := recover().(error); ok {
				getFailed = true
			}
		}()

		s1.Get(0)
	}()

	if !getFailed {
		t.Fatalf("should panic when getting on zero value")
	}

	var vals []string
	for _, v := range s1.Range() {
		vals = append(vals, v)
	}
	if vals != nil {
		t.Fatalf("range on zero value should be nil; got %v", vals)
	}

	appendFailed := false
	func() {
		defer func() {
			if _, ok := recover().(error); ok {
				appendFailed = true
			}
		}()

		s1.Append("test")
	}()

	if appendFailed {
		t.Fatalf("should not panic when appending on zero value")
	}

	// assume that the unexported fields behave well from a zero value
	// perspective
}

func TestMySliceAsMutableImmutableReceiver(t *testing.T) {
	s1 := new(coretest.MySlice)
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

func TestMySliceAsMutableMutableReceiver(t *testing.T) {
	s1 := new(coretest.MySlice).AsMutable()
	s2 := s1.AsMutable()

	if s1 != s2 {
		t.Fatalf("s1 and s2 should not be different values; they are")
	}

	if !s2.Mutable() {
		t.Fatalf("s2 should be mutable; it is not")
	}
}

func TestMySliceWithMutableImmutableReceiver(t *testing.T) {
	wasMutable := false

	var s3 *coretest.MySlice

	s1 := new(coretest.MySlice)
	s2 := s1.WithMutable(func(s *coretest.MySlice) {
		wasMutable = s.Mutable()

		// have some side effect
		s.Append(peter)

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

func TestMySliceWithMutableMutableReceiver(t *testing.T) {
	wasMutable := false

	var s3 *coretest.MySlice

	s1 := new(coretest.MySlice).AsMutable()
	s2 := s1.WithMutable(func(s *coretest.MySlice) {
		wasMutable = s.Mutable()

		// have some side effect
		s.Append(peter)

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

func TestMySliceWithImmutableImmutableReceiver(t *testing.T) {
	wasMutable := false

	var s3 *coretest.MySlice

	s1 := new(coretest.MySlice)
	s2 := s1.WithImmutable(func(s *coretest.MySlice) {
		wasMutable = s.Mutable()

		// have some side effect
		s.Append(peter)

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

func TestMySliceWithImmutableMutableReceiver(t *testing.T) {
	wasMutable := false

	var s3 *coretest.MySlice

	s1 := new(coretest.MySlice).AsMutable()
	s2 := s1.WithImmutable(func(s *coretest.MySlice) {
		wasMutable = s.Mutable()

		// have some side effect
		s.Append(peter)

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

func TestMySliceAsImmutableImmutableReceiver(t *testing.T) {

	v := func() *coretest.MySlice {
		return new(coretest.MySlice)
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
		s2 := s1.AsImmutable(new(coretest.MySlice))

		if s2 != s1 {
			t.Fatalf("expected s1 and s2 to be the same value")
		}

		if s2.Mutable() {
			t.Fatalf("s2 should not be mutable")
		}
	}
}

func TestMySliceAsImmutableMutableReceiver(t *testing.T) {

	v := func() *coretest.MySlice {
		return new(coretest.MySlice).AsMutable()
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
		s2 := s1.AsImmutable(new(coretest.MySlice))

		if s2 != s1 {
			t.Fatalf("expected s1 and s2 to be the same value")
		}

		if s2.Mutable() {
			t.Fatalf("s2 should not be mutable")
		}
	}
}

func TestMySliceVariadicConstructor(t *testing.T) {
	a := "a"
	b := "b"
	s1 := coretest.NewMySlice(a, b)

	if v := s1.Len(); v != 2 {
		t.Fatalf("length of s1 should be 2; got %v", v)
	}

	if s1.Mutable() {
		t.Fatalf("s1 should not be mutable; it is")
	}

	if v1, v2 := s1.Get(0), s1.Get(1); v1 != a || v2 != b {
		t.Fatalf("expected values to be (%q, %q), got (%q, %q)", a, b, v1, v2)
	}
}

func TestMySliceConstructorLength(t *testing.T) {
	a := ""
	b := ""
	s1 := coretest.NewMySliceLen(2)

	if v := s1.Len(); v != 2 {
		t.Fatalf("length of s1 should be 2; got %v", v)
	}

	if s1.Mutable() {
		t.Fatalf("s1 should not be mutable; it is")
	}

	if v1, v2 := s1.Get(0), s1.Get(1); v1 != a || v2 != b {
		t.Fatalf("expected values to be (%q, %q), got (%q, %q)", a, b, v1, v2)
	}
}

func TestMySliceSetImmutableReceiver(t *testing.T) {

	s1 := coretest.NewMySliceLen(1)
	s2 := s1.Set(0, peter)

	if s1 == s2 {
		t.Fatalf("s1 and s2 should be different values; they are not")
	}

	if v := s1.Get(0); v != "" {
		t.Fatalf("expected s1.Get(0) to be %q; got %q", "", v)
	}

	if v := s2.Get(0); v != peter {
		t.Fatalf("expected Get(0) to be %q, got %q", peter, v)
	}

	if s2.Mutable() {
		t.Fatalf("s2 should not be mutable")
	}
}

func TestMySliceSetMutableReceiver(t *testing.T) {

	s1 := coretest.NewMySliceLen(1).AsMutable()
	s2 := s1.Set(0, peter)

	if s1 != s2 {
		t.Fatalf("s1 and s2 should not be different values; they are")
	}

	if v := s2.Get(0); v != peter {
		t.Fatalf("expected Get(0) to be %q, got %q", peter, v)
	}

	if !s2.Mutable() {
		t.Fatalf("s2 should be mutable")
	}
}

func TestMySliceAppendImmutableReceiver(t *testing.T) {

	s1 := new(coretest.MySlice)
	s2 := s1.Append(peter)

	if s1 == s2 {
		t.Fatalf("s1 and s2 should be different values; they are not")
	}

	if v := s1.Len(); v != 0 {
		t.Fatalf("expected s1.Len() to be 0; got %v", v)
	}

	if v := s2.Get(0); v != peter {
		t.Fatalf("expected Get(0) to be %q, got %q", peter, v)
	}

	if s2.Mutable() {
		t.Fatalf("s2 should not be mutable")
	}
}

func TestMySliceAppendMutableReceiver(t *testing.T) {

	s1 := new(coretest.MySlice).AsMutable()
	s2 := s1.Append(peter)

	if s1 != s2 {
		t.Fatalf("s1 and s2 should not be different values; they are")
	}

	if v := s2.Get(0); v != peter {
		t.Fatalf("expected Get(0) to be %q, got %q", peter, v)
	}

	if !s2.Mutable() {
		t.Fatalf("s2 should be mutable")
	}
}
