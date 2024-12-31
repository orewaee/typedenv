package typedenv

import (
	"testing"
	"time"
)

func TestEmptyDuration(t *testing.T) {
	key := "MINUTE"
	want := time.Duration(0)

	got := Duration(key)

	if got != want {
		t.Fatalf("Duration(%q, %d) = %d", key, want, got)
	}
}

func TestDefaultDuration(t *testing.T) {
	key := "MINUTE"
	want := time.Minute

	got := Duration(key, time.Minute)

	if got != want {
		t.Fatalf("Duration(%q, %d) = %d", key, want, got)
	}
}

func TestGlobalDefaultDuration(t *testing.T) {
	key := "MINUTE"
	want := time.Minute

	DefaultDuration(key, want)

	got := Duration(key)

	if got != want {
		t.Fatalf("Duration(%q, %d) = %d", key, want, got)
	}
}
