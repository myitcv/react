package coretest_test

import (
	"testing"

	"myitcv.io/immutable/cmd/immutableGen/internal/coretest"
)

func TestStruct(t *testing.T) {
	a1 := new(coretest.A)

	if !a1.IsDeeplyNonMutable(nil) {
		t.Fatalf("a1 is zero value; should be DeeplyNonMutable")
	}

	a2 := a1.AsMutable()

	if a2.IsDeeplyNonMutable(nil) {
		t.Fatalf("a2 should not be DeeplyNonMutable")
	}

	a1 = a1.SetA(a2)

	if a1.IsDeeplyNonMutable(nil) {
		t.Fatalf("a1 references a mutable value; should not be DeeplyNonMutable")
	}

	a3 := new(coretest.A).AsMutable()
	a3.SetA(a3)

	if a3.IsDeeplyNonMutable(nil) {
		t.Fatalf("a3 should not be DeeplyNonMutable")
	}

	a3.AsImmutable(nil)

	if !a3.IsDeeplyNonMutable(nil) {
		t.Fatalf("a3 should be DeeplyNonMutable")
	}

	a4 := new(coretest.A).SetBlah(coretest.BlahNonMutable(struct{}{}))

	if !a4.IsDeeplyNonMutable(nil) {
		t.Fatalf("a4 should be DeeplyNonMutable")
	}

	a5 := new(coretest.A).SetBlah(coretest.BlahMutable(struct{}{}))

	if a5.IsDeeplyNonMutable(nil) {
		t.Fatalf("a5 should not be DeeplyNonMutable")
	}
}

func TestSlice(t *testing.T) {
	s1 := coretest.NewAS()

	if !s1.IsDeeplyNonMutable(nil) {
		t.Fatalf("s1 should be DeeplyNonMutable")
	}

	s2 := coretest.NewAS().AsMutable()

	if s2.IsDeeplyNonMutable(nil) {
		t.Fatalf("s2 should not be DeeplyNonMutable")
	}

	aimm := new(coretest.A)
	amut := new(coretest.A).AsMutable()

	s3 := coretest.NewAS(aimm)
	s4 := coretest.NewAS(amut)

	if !s3.IsDeeplyNonMutable(nil) {
		t.Fatalf("s3 should not be DeeplyNonMutable")
	}

	if s4.IsDeeplyNonMutable(nil) {
		t.Fatalf("s4 should not be DeeplyNonMutable")
	}
}

func TestMap(t *testing.T) {
	m1 := coretest.NewAM()

	if !m1.IsDeeplyNonMutable(nil) {
		t.Fatalf("m1 should be DeeplyNonMutable")
	}

	m2 := coretest.NewAM().AsMutable()

	if m2.IsDeeplyNonMutable(nil) {
		t.Fatalf("m2 should not be DeeplyNonMutable")
	}

	// DeeplyNonMutable
	aimm := new(coretest.A)

	m3 := coretest.NewAM().Set(nil, aimm)
	m4 := coretest.NewAM().Set(aimm, nil)
	m5 := coretest.NewAM().Set(aimm, aimm)

	if !m3.IsDeeplyNonMutable(nil) {
		t.Fatalf("m3 should not be DeeplyNonMutable")
	}

	if !m4.IsDeeplyNonMutable(nil) {
		t.Fatalf("m4 should not be DeeplyNonMutable")
	}

	if !m5.IsDeeplyNonMutable(nil) {
		t.Fatalf("m5 should not be DeeplyNonMutable")
	}

	// not DeeplyNonMutable
	amut := new(coretest.A).AsMutable()

	m6 := coretest.NewAM().Set(nil, amut)
	m7 := coretest.NewAM().Set(amut, nil)
	m8 := coretest.NewAM().Set(amut, amut)

	if m6.IsDeeplyNonMutable(nil) {
		t.Fatalf("m6 should not be DeeplyNonMutable")
	}

	if m7.IsDeeplyNonMutable(nil) {
		t.Fatalf("m7 should not be DeeplyNonMutable")
	}

	if m8.IsDeeplyNonMutable(nil) {
		t.Fatalf("m8 should not be DeeplyNonMutable")
	}
}
