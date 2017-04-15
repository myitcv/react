package coretest_test

import (
	"testing"

	"myitcv.io/immutable/cmd/immutableGen/internal/coretest"
)

func TestMyStructZeroValue(t *testing.T) {
	s1 := new(coretest.MyStruct)

	if s1.Mutable() {
		t.Fatalf("zero value should be immutable")
	}

	if v := s1.Name(); v != "" {
		t.Fatalf("zero value for Name() should be \"\", got %q", v)
	}

	// assume that the unexported fields behave well from a zero value
	// perspective
}

func TestMyStructAsMutableImmutableReceiver(t *testing.T) {
	s1 := new(coretest.MyStruct)
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

func TestMyStructAsMutableMutableReceiver(t *testing.T) {
	s1 := new(coretest.MyStruct).AsMutable()
	s2 := s1.AsMutable()

	if s1 != s2 {
		t.Fatalf("s1 and s2 should not be different values; they are")
	}

	if !s2.Mutable() {
		t.Fatalf("s2 should be mutable; it is not")
	}
}

func TestMyStructWithMutableImmutableReceiver(t *testing.T) {
	wasMutable := false

	var s3 *coretest.MyStruct

	s1 := new(coretest.MyStruct)
	s2 := s1.WithMutable(func(s *coretest.MyStruct) {
		wasMutable = s.Mutable()

		// have some side effect
		s.SetName(peter)

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

func TestMyStructWithMutableMutableReceiver(t *testing.T) {
	wasMutable := false

	var s3 *coretest.MyStruct

	s1 := new(coretest.MyStruct).AsMutable()
	s2 := s1.WithMutable(func(s *coretest.MyStruct) {
		wasMutable = s.Mutable()

		// have some side effect
		s.SetName(peter)

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

func TestMyStructWithImmutableImmutableReceiver(t *testing.T) {
	wasMutable := false

	var s3 *coretest.MyStruct

	s1 := new(coretest.MyStruct)
	s2 := s1.WithImmutable(func(s *coretest.MyStruct) {
		wasMutable = s.Mutable()

		// have some side effect
		s.SetName(peter)

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

func TestMyStructWithImmutableMutableReceiver(t *testing.T) {
	wasMutable := false

	var s3 *coretest.MyStruct

	s1 := new(coretest.MyStruct).AsMutable()
	s2 := s1.WithImmutable(func(s *coretest.MyStruct) {
		wasMutable = s.Mutable()

		// have some side effect
		s.SetName(peter)

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

func TestMyStructAsImmutableImmutableReceiver(t *testing.T) {

	v := func() *coretest.MyStruct {
		return new(coretest.MyStruct)
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
		s2 := s1.AsImmutable(new(coretest.MyStruct))

		if s2 != s1 {
			t.Fatalf("expected s1 and s2 to be the same value")
		}

		if s2.Mutable() {
			t.Fatalf("s2 should not be mutable")
		}
	}
}

func TestMyStructAsImmutableMutableReceiver(t *testing.T) {

	v := func() *coretest.MyStruct {
		return new(coretest.MyStruct).AsMutable()
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
		s2 := s1.AsImmutable(new(coretest.MyStruct))

		if s2 != s1 {
			t.Fatalf("expected s1 and s2 to be the same value")
		}

		if s2.Mutable() {
			t.Fatalf("s2 should not be mutable")
		}
	}
}

func TestMyStructSetNameImmutableReceiver(t *testing.T) {

	s1 := new(coretest.MyStruct)
	s2 := s1.SetName(peter)

	if s1 == s2 {
		t.Fatalf("s1 and s2 should be different values; they are not")
	}

	if v := s1.Name(); v != "" {
		t.Fatalf("expected s1.Name() to be %q; got %q", "", v)
	}

	if v := s2.Name(); v != peter {
		t.Fatalf("expected Name() to be %q, got %q", peter, v)
	}

	if s2.Mutable() {
		t.Fatalf("s2 should not be mutable")
	}
}

func TestMyStructSetNameMutableReceiver(t *testing.T) {

	s1 := new(coretest.MyStruct).AsMutable()
	s2 := s1.SetName(peter)

	if s1 != s2 {
		t.Fatalf("s1 and s2 should not be different values; they are")
	}

	if v := s2.Name(); v != peter {
		t.Fatalf("expected Name() to be %q, got %q", peter, v)
	}

	if !s2.Mutable() {
		t.Fatalf("s2 should be mutable")
	}
}

func TestSpecialVersionBump(t *testing.T) {
	s1 := new(coretest.MyStruct)

	if v := s1.Key().Version; v != 0 {
		t.Fatalf("expected version 0 as initial value, got %v", v)
	}

	s2 := s1.AsMutable()

	if v := s1.Key().Version; v != 0 {
		t.Fatalf("expected version 0 as initial value, got %v", v)
	}

	if v := s2.Key().Version; v != 1 {
		t.Fatalf("expected version 1 as initial value, got %v", v)
	}

	s2.AsImmutable(s1)

	if v := s1.Key().Version; v != 0 {
		t.Fatalf("expected version 0 as initial value, got %v", v)
	}

	if v := s2.Key().Version; v != 1 {
		t.Fatalf("expected version 1 as initial value, got %v", v)
	}

	s3 := s2.SetName("test")

	if v := s2.Key().Version; v != 1 {
		t.Fatalf("expected version 1 as initial value, got %v", v)
	}

	if v := s3.Key().Version; v != 2 {
		t.Fatalf("expected version 2 as initial value, got %v", v)
	}

	s3.AsImmutable(s2)

	if v := s2.Key().Version; v != 1 {
		t.Fatalf("expected version 1 as initial value, got %v", v)
	}

	if v := s3.Key().Version; v != 2 {
		t.Fatalf("expected version 2 as initial value, got %v", v)
	}
}
