package naijafakergo

import (
	"fmt"
	"github.com/kodegrenade/naija-faker-go/internal/validate"
	"github.com/kodegrenade/naija-faker-go/providers"
)

// Address generates a Nigerian street address.
// lang is optional — pass empty string for random.
func (f *Faker) Address(lang string) (string, error) {
	lang = f.resolveLanguage(lang)

	if err := validate.Language(lang); err != nil {
		return "", err
	}

	return f.buildAddress(lang), nil
}

func (f *Faker) States() []string {
	return providers.Locations.States
}

func (f *Faker) LGAs() []string {
	return providers.Locations.Lgas
}

func (f *Faker) buildAddress(lang string) string {
	streetType := f.pick(providers.Addresses.StreetTypes)
	locale := f.pick(providers.Addresses.Locale)
	prefix := f.pick(providers.Addresses.StreetPrefixes)
	place := f.pick(providers.Addresses.Places[locale])
	number := f.random(100) + 1
	streetName := f.generateStreetName(lang)

	return fmt.Sprintf("%s %d, %s %s, %s", prefix, number, streetName, streetType, place)
}

func (f *Faker) generateStreetName(lang string) string {
	// street names are derived from person names — authentic Nigerian street naming
	firstName, _ := f.firstName(lang, f.pick(genders))
	lastName, _ := f.lastName(lang, "male")
	return firstName + " " + lastName
}