package typedenv

import "testing"

func TestEmptyString(t *testing.T) {
	key := "STRING"
	want := ""

	got := String(key)

	if got != want {
		t.Fatalf("String(%q, %q) = %q", key, want, got)
	}
}

func TestDefaultString(t *testing.T) {
	key := "STRING"
	want := "hello world"

	got := String(key, want)

	if got != want {
		t.Fatalf("String(%q, %q) = %q", key, want, got)
	}
}

func TestGlobalDefaultString(t *testing.T) {
	key := "STRING"
	want := "hello world"

	DefaultString(key, want)

	got := String(key)

	if got != want {
		t.Fatalf("String(%q, %q) = %q", key, want, got)
	}
}
