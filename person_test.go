package naijafakergo_test

import (
    "testing"
    naijafaker "github.com/kodegrenade/naija-faker-go"
)

func TestPerson(t *testing.T) {
    f := naijafaker.New()

    t.Run("returns a person with no arguments", func(t *testing.T) {
        person, err := f.Person("", "")
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        if person.FullName == "" {
            t.Error("expected a non-empty full name")
        }
        if person.Email == "" {
            t.Error("expected a non-empty email")
        }
        if person.Phone == "" {
            t.Error("expected a non-empty phone number")
        }
    })

    t.Run("returns a yoruba female person", func(t *testing.T) {
        person, err := f.Person("yoruba", "female")
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        if person.Title == "" {
            t.Error("expected a non-empty title")
        }
    })

    t.Run("returns error for invalid language", func(t *testing.T) {
        _, err := f.Person("klingon", "")
        if err == nil {
            t.Fatal("expected an error for invalid language, got nil")
        }
    })
}

func TestPeople(t *testing.T) {
    f := naijafaker.New()

    t.Run("returns 10 persons by default", func(t *testing.T) {
        people, err := f.People(0, "", "")
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        if len(people) != 10 {
            t.Errorf("expected 10 people, got %d", len(people))
        }
    })

    t.Run("returns exact count requested", func(t *testing.T) {
        people, err := f.People(5, "igbo", "male")
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        if len(people) != 5 {
            t.Errorf("expected 5 people, got %d", len(people))
        }
    })
}

func TestConsistentPerson(t *testing.T) {
    f := naijafaker.New()

    t.Run("state and LGA are present", func(t *testing.T) {
        person, err := f.ConsistentPerson("hausa", "male")
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        if person.State == "" {
            t.Error("expected a non-empty state")
        }
        if person.LGA == "" {
            t.Error("expected a non-empty LGA")
        }
    })
}