package naijafakergo_test

import (
	naijafakergo "github.com/kodegrenade/naija-faker-go"
	"testing"
)

func TestFinance(t *testing.T) {
	f := naijafakergo.New()

	t.Run("returns a salary band without argument", func(t *testing.T) {
		salary, err := f.Salary("")
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		if salary.Amount < 0 {
			t.Error("expected a salary value greater than zero")
		}
	})

	t.Run("returns a salary band with argument", func (t *testing.T)  {
		salary, err := f.Salary("entry")
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		if salary.Amount < 0 {
			t.Error("expected a salary value greater than zero")
		}
	})

	t.Run("returns error for invalid level", func (t *testing.T)  {
		_, err := f.Salary("higher")
		if err == nil {
			t.Fatalf("expected an error for invalid salary level, got nil")
		}
	})
}
