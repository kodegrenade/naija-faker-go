package naijafakergo

import (
	"github.com/kodegrenade/naija-faker-go/internal/validate"
	"github.com/kodegrenade/naija-faker-go/providers"
	"github.com/kodegrenade/naija-faker-go/types"
)

var languageToRegions = map[string][]string{
	"yoruba": {"east"},
	"hausa":  {"north"},
	"igbo":   {"east", "south"},
}

// Person generates a single fake Nigerian person.
// lang and gender are optional — pass empty strings for random.
func (f *Faker) Person(lang, gender string) (types.Person, error) {
	lang = f.resolveLanguage(lang)
	gender = f.resolveGender(gender)

	if err := validate.Language(lang); err != nil {
		return types.Person{}, err
	}
	if err := validate.Gender(gender); err != nil {
		return types.Person{}, err
	}

	return f.buildPerson(lang, gender)
}

// People generates a slice of fake Nigerian persons.
// count defaults to 10 if 0 is passed.
// lang and gender are optional.
func (f *Faker) People(count int, lang, gender string) ([]types.Person, error) {
	if count <= 0 {
		count = 10
	}

	people := make([]types.Person, 0, count)

	for range count {
		p, err := f.Person(lang, gender)
		if err != nil {
			return nil, err
		}
		people = append(people, p)
	}

	return people, nil
}

// ConsistentPerson generates a person where name, address, state, and LGA
// are all geographically coherent for the given ethnic group.
func (f *Faker) ConsistentPerson(lang, gender string) (types.ConsistentPerson, error) {
	lang = f.resolveLanguage(lang)
	gender = f.resolveGender(gender)

	if err := validate.Language(lang); err != nil {
		return types.ConsistentPerson{}, err
	}
	if err := validate.Gender(gender); err != nil {
		return types.ConsistentPerson{}, err
	}

	person, err := f.buildPerson(lang, gender)
	if err != nil {
		return types.ConsistentPerson{}, err
	}

	regions := languageToRegions[lang]
	region := f.pick(regions)
	regionData := providers.Geo.RegionMap[region]
	state := f.pick(regionData.States)
	lga := f.pick(providers.Geo.StateLgas[state])

	return types.ConsistentPerson{
		Person: person,
		State:  state,
		LGA:    lga,
	}, nil
}

// ConsistentPeople generates multiple geographically consistent persons.
func (f *Faker) ConsistentPeople(count int, lang, gender string) ([]types.ConsistentPerson, error) {
	if count <= 0 {
		count = 10
	}

	people := make([]types.ConsistentPerson, 0, count)

	for range count {
		p, err := f.ConsistentPerson(lang, gender)
		if err != nil {
			return nil, err
		}
		people = append(people, p)
	}

	return people, nil
}

func (f *Faker) buildPerson(lang, gender string) (types.Person, error) {
	title, err := f.Title(gender)
	if err != nil {
		return types.Person{}, err
	}

	firstName, err := f.firstName(lang, gender)
	if err != nil {
		return types.Person{}, err
	}

	lastName, err := f.lastName(lang, "male")
	if err != nil {
		return types.Person{}, err
	}

	fullName := firstName + " " + lastName

	email, err := f.Email(fullName)
	if err != nil {
		return types.Person{}, err
	}

	phone, err := f.PhoneNumber("")
	if err != nil {
		return types.Person{}, err
	}

	address, err := f.Address(lang)
	if err != nil {
		return types.Person{}, err
	}

	return types.Person{
		Title:     title,
		FirstName: firstName,
		LastName:  lastName,
		FullName:  fullName,
		Email:     email,
		Phone:     phone,
		Address:   address,
	}, nil
}
