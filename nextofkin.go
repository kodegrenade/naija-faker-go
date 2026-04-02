package naijafakergo

import (
	"github.com/kodegrenade/naija-faker-go/internal/validate"
	"github.com/kodegrenade/naija-faker-go/types"
)

var relationshipTitles = map[string][]string{
	"male":   {"Father", "Brother", "Spouse", "Uncle", "Son", "Friend"},
	"female": {"Mother", "Sister", "Spouse", "Aunt", "Daughter", "Friend"},
}

func (f *Faker) NextOfKin(lang, gender string) (types.NextOfKin, error) {
	lang = f.resolveLanguage(lang)
	gender = f.resolveGender(gender)

	if err := validate.Language(lang); err != nil {
		return types.NextOfKin{}, err
	}
	if err := validate.Gender(gender); err != nil {
		return types.NextOfKin{}, err
	}

	fullName, err := f.Name(lang, gender)
	if err != nil {
		return types.NextOfKin{}, err
	}

	phone, err := f.PhoneNumber("")
	if err != nil {
		return types.NextOfKin{}, err
	}

	address, err := f.Address(lang)
	if err != nil {
		return types.NextOfKin{}, err
	}

	relationship := f.pick(relationshipTitles[gender])

	return types.NextOfKin{
		FullName:     fullName,
		Relationship: relationship,
		Phone:        phone,
		Address:      address,
	}, nil
}
