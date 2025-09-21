package typedenv

import "testing"

func TestEmptyInt64(t *testing.T) {
	key := "INT64"
	var want int64

	got := Int64(key)

	if got != want {
		t.Fatalf("Int64(%q, %d) = %d", key, want, got)
	}
}

func TestDefaultInt64(t *testing.T) {
	key := "INT64"
	want := int64(1)

	got := Int64(key, want)

	if got != want {
		t.Fatalf("Int64(%q, %d) = %d", key, want, got)
	}
}

func TestGlobalDefaultInt64(t *testing.T) {
	key := "INT64"
	want := int64(1)

	DefaultInt64(key, want)

	got := Int64(key)

	if got != want {
		t.Fatalf("Int64(%q, %d) = %d", key, want, got)
	}
}
