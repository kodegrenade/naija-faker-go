package naijafakergo_test

import (
	naijafakergo "github.com/kodegrenade/naija-faker-go"
	"testing"
)

func TestNextOfKin(t *testing.T) {
	f := naijafakergo.New()

	t.Run("returns a next of kin without arguments", func(t *testing.T) {
		nextOfKin, err := f.NextOfKin("", "")
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		if nextOfKin.FullName == "" {
			t.Error("expected a non-empty next of kin fullname")
		}
		if nextOfKin.Address == "" {
			t.Error("expected a non-empty next of kin address")
		}
		if nextOfKin.Phone == "" {
			t.Error("expected a non-empty next of kin phone number")
		}
		if nextOfKin.Relationship == "" {
			t.Error("expected a non-empty next of kin relationship")
		}
	})

	t.Run("returns a next of kin with arguments", func(t *testing.T) {
		nextOfKin, err := f.NextOfKin("hausa", "male")
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		if nextOfKin.FullName == "" {
			t.Error("expected a non-empty next of kin fullname")
		}
		if nextOfKin.Address == "" {
			t.Error("expected a non-empty next of kin address")
		}
		if nextOfKin.Phone == "" {
			t.Error("expected a non-empty next of kin phone number")
		}
		if nextOfKin.Relationship == "" {
			t.Error("expected a non-empty next of kin relationship")
		}
	})

	t.Run("returns error for invalid language", func(t *testing.T) {
		_, err := f.NextOfKin("tiv", "")
		if err == nil {
			t.Fatalf("expected an error for invalid language")
		}
	})

	t.Run("returns error for invalid gender", func(t *testing.T) {
		_, err := f.NextOfKin("", "boy")
		if err == nil {
			t.Fatalf("expected an error for invalid language")
		}
	})
}
