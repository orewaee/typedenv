package typedenv

import (
	"fmt"
	"testing"
)

func TestLookupTrue(t *testing.T) {
	key := "TEST"
	want := true

	DefaultInt(key, 1)
	defer delete(defaultInt, key)

	got := Lookup(key)

	if got != want {
		t.Fatalf("want %t, got %t", want, got)
	}
}

func TestLookupFalse(t *testing.T) {
	key := "TEST"
	want := false

	got := Lookup(key)

	if got != want {
		fmt.Println(defaultBool, defaultInt, defaultInt64, defaultString, defaultDuration)
		t.Fatalf("want %t, got %t", want, got)
	}
}
