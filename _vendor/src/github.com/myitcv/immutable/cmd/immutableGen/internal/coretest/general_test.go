package coretest

import "testing"

func TestAnonFields(t *testing.T) {
	m := new(MyStruct)

	if v := m.string(); v != "" {
		t.Fatalf("expected zero value to be %v", "")
	}

	val := "test"
	m = m.setString(val)

	if v := m.string(); v == "" || v != val {
		t.Fatalf("expected set value to be %q", val)
	}
}
