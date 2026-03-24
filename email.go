package naijafakergo

import (
	"fmt"
	"github.com/kodegrenade/naija-faker-go/providers"
	"strings"
)

var emailProviders []string = providers.Emails.Extensions

// Email generates a Nigerian email address.
// If a full name is provided, it derives the email from it.
// Otherwise it generates a name first then derives from that.
func (f *Faker) Email(fullName string) (string, error) {
	if fullName == "" {
		name, err := f.Name("", "")
		if err != nil {
			return "", err
		}
		fullName = name
	}

	return f.deriveEmail(fullName), nil
}

// deriveEmail builds an email from a full name.
func (f *Faker) deriveEmail(fullName string) string {
	parts := strings.Fields(strings.ToLower(fullName))

	var local string
	switch f.random(3) {
	case 0:
		local = strings.Join(parts, ".")
	case 1:
		local = strings.Join(parts, "_")
	case 2:
		local = fmt.Sprintf("%s.%s%d", parts[0], parts[len(parts)-1], f.random(99)+1)
	}

	provider := f.pick(emailProviders)
	return local + "@" + provider
}
