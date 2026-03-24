package naijafakergo

import (
	"fmt"
	"unicode"

	"github.com/kodegrenade/naija-faker-go/internal/errors"
	"github.com/kodegrenade/naija-faker-go/providers"
)

var validNetworks = map[string]bool{
	"mtn":     true,
	"glo":     true,
	"airtel":  true,
	"9mobile": true,
}

// PhoneNumber generates a Nigerian phone number.
// network is optional — pass empty string for random network.
func (f *Faker) PhoneNumber(network string) (string, error) {
	network = f.resolveNetwork(network)

	if !validNetworks[network] {
		return "", errors.New(
			errors.ErrInvalidNetwork,
			"network must be one of: mtn, glo, airtel, 9mobile",
		)
	}

	prefixes, ok := providers.Numbers.Prefix[capitalizeFirstLetter(network)]
	if !ok || len(prefixes) == 0 {
		return "", errors.New(errors.ErrInvalidNetwork, "no prefixes found for network: "+network)
	}

	prefix := f.pick(prefixes)
	suffix := fmt.Sprintf("%07d", f.random(9999999)+1)

	return "+234" + prefix[1:] + suffix, nil
}

func (f *Faker) resolveNetwork(network string) string {
	if network != "" {
		return network
	}
	if f.network != "" {
		return f.network
	}
	return f.pick(resolveWordCase(providers.Numbers.Networks))
}

func resolveWordCase(networks []string) []string {
	for i := range networks {
		networks[i] = lowercaseFirstLetter(networks[i])
	}

	return networks
}

func lowercaseFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
