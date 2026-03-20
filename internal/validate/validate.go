package validate

import "github.com/kodegrenade/naija-faker-go/internal/errors"

var validLanguages = map[string]bool{
	"yoruba": true,
	"igbo":   true,
	"hausa":  true,
}

var validGenders = map[string]bool{
	"male":   true,
	"female": true,
}

// Language checks that the provided language is supported.
// Returns nil if valid, a NaijaFakerError if not.
func Language(lang string) error {
	if !validLanguages[lang] {
		return errors.New(
			errors.ErrInvalidLanguage,
			"language must be one of: yoruba, igbo, hausa",
		)
	}
	return nil
}

// Gender checks that the provided gender is supported.
func Gender(gender string) error {
	if !validGenders[gender] {
		return errors.New(
			errors.ErrInvalidGender,
			"gender must be one of: male, female",
		)
	}
	return nil
}
