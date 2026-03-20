package naijafakergo_test

import (
	naijafakergo "github.com/kodegrenade/naija-faker-go"
	"testing"
)

func TestName(t *testing.T) {
	f := naijafakergo.New()

	t.Run("returns a name with no arguments", func(t *testing.T) {
		name, err := f.Name("", "")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if name == "" {
			t.Error("expected a non-empty name")
		}
	})

	t.Run("returns a yoruba male name", func(t *testing.T) {
		// seeded for determinism
		f.Seed(42)
		name, err := f.Name("yoruba", "male")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if name == "" {
			t.Error("expected a non-empty name")
		}
	})

	t.Run("returns error for invalid language", func(t *testing.T) {
		_, err := f.Name("klingon", "")
		if err == nil {
			t.Fatal("expected an error for invalid language, got nil")
		}
	})
}

func TestTitle(t *testing.T) {
	f := naijafakergo.New()

	t.Run("returns a male title", func(t *testing.T) {
		title, err := f.Title("male")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if title == "" {
			t.Error("expected a non-empty title")
		}
	})

	t.Run("returns error for invalid gender", func(t *testing.T) {
		_, err := f.Title("unknown")
		if err == nil {
			t.Fatal("expected an error for invalid gender, got nil")
		}
	})
}
