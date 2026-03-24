package naijafakergo_test

import (
	naijafakergo "github.com/kodegrenade/naija-faker-go"
	"testing"
)

func TestStates(t *testing.T) {
	f := naijafakergo.New()

	t.Run("returns a list of states", func(t *testing.T) {
		states := f.States()
		if len(states) < 1 {
			t.Error("expected at least one state")
		}

		if len(states) < 37 {
			t.Error("expected at least 37 states with FCT included")
		}
	})
}

func TestLGAs(t *testing.T) {
	f := naijafakergo.New()

	t.Run("returns a list of LGAs", func(t *testing.T) {
		lgas := f.LGAs()
		if len(lgas) < 1 {
			t.Error("expected at least one LGA")
		}

		if len(lgas) < 774 {
			t.Error("expectd at least 774 LGAs in the federation")
		}
	})
}

func TestAddress(t *testing.T) {
	f := naijafakergo.New()

	t.Run("returns an address from the igbo region", func(t *testing.T) {
		address, err := f.Address("igbo")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if address == "" {
			t.Error("expected a non-empty address")
		}
	})

	t.Run("returns an address from the yoruba region", func(t *testing.T) {
		address, err := f.Address("yoruba")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if address == "" {
			t.Error("expected a non-empty address")
		}
	})

	t.Run("returns an address from the hausa region", func(t *testing.T) {
		address, err := f.Address("hausa")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if address == "" {
			t.Error("expected a non-empty address")
		}
	})

	t.Run("returns error for invalid language", func(t *testing.T) {
		_, err := f.Address("klingon")
		if err == nil {
			t.Fatal("expected an error for invalid language, got nil")
		}
	})
}