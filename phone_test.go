package naijafakergo_test

import (
	naijafakergo "github.com/kodegrenade/naija-faker-go"
	"testing"
)

func TestPhoneNumber(t *testing.T) {
	f := naijafakergo.New()

	t.Run("return a phone number with no network argument", func(t *testing.T) {
		phoneNumber, err := f.PhoneNumber("")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if phoneNumber == "" {
			t.Error("expected a non-empty phone number")
		}
	})

	t.Run("return a phone number with network argument", func(t *testing.T) {
		phoneNumber, err := f.PhoneNumber("mtn")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(phoneNumber) < 13 {
			t.Error("expected a phone number value with length greater than 13")
		}
	})
}