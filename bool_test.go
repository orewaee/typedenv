package typedenv

import "testing"

func TestEmptyBool(t *testing.T) {
	key := "BOOL"
	want := false

	got := Bool(key)

	if got != want {
		t.Fatalf("Bool(%q, %v) = %v", key, want, got)
	}
}

func TestDefaultBool(t *testing.T) {
	key := "BOOL"
	want := true

	got := Bool(key, want)

	if got != want {
		t.Fatalf("Bool(%q, %v) = %v", key, want, got)
	}
}

func TestGlobalDefaultBool(t *testing.T) {
	key := "BOOL"
	want := true

	DefaultBool(key, want)

	got := Bool(key)

	if got != want {
		t.Fatalf("Bool(%q, %v) = %v", key, want, got)
	}
}
