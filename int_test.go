package typedenv

import "testing"

func TestEmptyInt(t *testing.T) {
	key := "INT"
	want := 0

	got := Int(key)

	if got != want {
		t.Fatalf("Int(%q, %d) = %d", key, want, got)
	}
}

func TestDefaultInt(t *testing.T) {
	key := "INT"
	want := 1

	got := Int(key, want)

	if got != want {
		t.Fatalf("Int(%q, %d) = %d", key, want, got)
	}
}

func TestGlobalDefaultInt(t *testing.T) {
	key := "INT"
	want := 1

	DefaultInt(key, want)

	got := Int(key)

	if got != want {
		t.Fatalf("Int(%q, %d) = %d", key, want, got)
	}
}
