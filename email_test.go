package naijafakergo_test

import (
	naijafakergo "github.com/kodegrenade/naija-faker-go"
	"testing"
)

func TestEmail (t *testing.T) {
	f := naijafakergo.New()

	t.Run("return an email with no argument", func(t *testing.T) {
		email, err := f.Email("")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if email == "" {
			t.Error("expected a non-empty email address")
		}
	})

	t.Run("return an email with argument", func(t *testing.T) {
		email, err := f.Email("Temitope Ayotunde")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if email == "" {
			t.Error("expected a non-empty email address")
		}
	})
}
