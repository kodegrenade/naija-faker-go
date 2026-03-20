package naijafakergo

import (
	"github.com/kodegrenade/naija-faker-go/internal/errors"
	"github.com/kodegrenade/naija-faker-go/internal/validate"
	"github.com/kodegrenade/naija-faker-go/providers"
)

// Name generates a full Nigerian name.
// lang and gender are optional — pass empty strings to use configured defaults or random.
func (f *Faker) Name(lang, gender string) (string, error) {
	lang = f.resolveLanguage(lang)
	gender = f.resolveGender(gender)

	if err := validate.Language(lang); err != nil {
		return "", err
	}
	if err := validate.Gender(gender); err != nil {
		return "", err
	}

	firstName, err := f.firstName(lang, gender)
	if err != nil {
		return "", err
	}

	lastName, err := f.lastName(lang, "male")
	if err != nil {
		return "", err
	}

	return firstName + " " + lastName, nil
}

// FirstName generates a Nigerian first name.
func (f *Faker) FirstName(lang, gender string) (string, error) {
	lang = f.resolveLanguage(lang)
	gender = f.resolveGender(gender)

	if err := validate.Language(lang); err != nil {
		return "", err
	}
	if err := validate.Gender(gender); err != nil {
		return "", err
	}

	return f.firstName(lang, gender)
}

// LastName generates a Nigerian last name for a given ethnic group.
func (f *Faker) LastName(lang string) (string, error) {
	lang = f.resolveLanguage(lang)

	if err := validate.Language(lang); err != nil {
		return "", err
	}

	// use male as default gender when generating last name
	return f.lastName(lang, "male")
}

// Title generates a Nigerian title.
// gender is optional — pass empty string for random.
func (f *Faker) Title(gender string) (string, error) {
	gender = f.resolveGender(gender)

	if err := validate.Gender(gender); err != nil {
		return "", err
	}

	titles, ok := providers.Titles.Title[gender]
	if !ok || len(titles) == 0 {
		return "", errors.New(errors.ErrInvalidParam, "no titles found for gender: "+gender)
	}

	return f.pick(titles), nil
}

// firstName picks a random first name from the data for a given language and gender.
func (f *Faker) firstName(lang, gender string) (string, error) {
	langData, ok := providers.Names.Yoruba[gender]

	// select the right language map
	switch lang {
	case "yoruba":
		langData, ok = providers.Names.Yoruba[gender]
	case "igbo":
		langData, ok = providers.Names.Igbo[gender]
	case "hausa":
		langData, ok = providers.Names.Hausa[gender]
	}

	if !ok || len(langData) == 0 {
		return "", errors.New(errors.ErrInvalidParam, "no names found for: "+lang+"/"+gender)
	}

	return f.pick(langData), nil
}

// lastName picks a random last name from the data for a given language.
func (f *Faker) lastName(lang, gender string) (string, error) {
	lastNames, ok := providers.Names.Yoruba[gender]

	// select the right language map
	switch lang {
	case "yoruba":
		lastNames, ok = providers.Names.Yoruba[gender]
	case "igbo":
		lastNames, ok = providers.Names.Igbo[gender]
	case "hausa":
		lastNames, ok = providers.Names.Hausa[gender]
	}

	if !ok || len(lastNames) == 0 {
		return "", errors.New(errors.ErrInvalidParam, "no last names found for: "+lang)
	}
	return f.pick(lastNames), nil
}
