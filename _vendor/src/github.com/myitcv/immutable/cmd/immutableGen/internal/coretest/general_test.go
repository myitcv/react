package coretest

import "testing"

func TestAnonFields(t *testing.T) {
	m := new(MyStruct)

	if v := m.string(); v != nil {
		t.Fatalf("expected zero value to be %v", nil)
	}

	val := "test"
	m = m.setString(&val)

	if v := m.string(); v == nil || *v != val {
		t.Fatalf("expected set value to be %q", val)
	}
}
