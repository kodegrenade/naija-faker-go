package naijafakergo_test

import (
	naijafakergo "github.com/kodegrenade/naija-faker-go"
	"testing"
)

func TestPersonalData(t *testing.T) {
	f := naijafakergo.New()

	t.Run("returns date of birth min and max age arguments", func(t *testing.T) {
		dateOfBirth := f.DateOfBirth(18, 99)
		if dateOfBirth.Age < 18 {
			t.Error("expected an age value not less than 18")
		}

		if dateOfBirth.Age > 99 {
			t.Error("expected an age value not greater than 99")
		}

		if dateOfBirth.Date == "" {
			t.Error("expected a non-empty date of birth")
		}
	})

	t.Run("returns a value for marital status", func(t *testing.T) {
		maritalStatus := f.MaritalStatus()
		if maritalStatus == "" {
			t.Error("expected a non-empty marital status")
		}
	})

	t.Run("returns a value for bloodgroup", func(t *testing.T) {
		bloodGroup := f.BloodGroup()
		if bloodGroup == "" {
			t.Error("expected a non-empty bloodgroup")
		}
	})

	t.Run("returns a value for genotype", func(t *testing.T) {
		genotype := f.Genotype()
		if genotype == "" {
			t.Error("expected a non-empty genotype")
		}
	})
}
