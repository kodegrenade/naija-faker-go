package naijafakergo

import (
	"fmt"
	"strings"

	"github.com/kodegrenade/naija-faker-go/internal/errors"
	"github.com/kodegrenade/naija-faker-go/providers"
	"github.com/kodegrenade/naija-faker-go/types"
)

var validLevels = map[string]bool{
	"entry":     true,
	"mid":       true,
	"senior":    true,
	"executive": true,
}

// Salary generates a Nigerian salary for a given level.
func (f *Faker) Salary(level string) (types.Salary, error) {
	if level == "" {
		level = f.pick([]string{"entry", "mid", "senior", "executive"})
	}

	if !validLevels[level] {
		return types.Salary{}, errors.New(
			errors.ErrInvalidLevel,
			"level must be one of: entry, junior, mid, senior, lead, executive",
		)
	}

	band := providers.Salaries.SalaryBands[level]
	amount := f.random(band.Max-band.Min+1) + band.Min

	return types.Salary{
		Amount:    amount,
		Currency:  "NGN",
		Level:     level,
		Frequency: "monthly",
	}, nil
}

// BankAccount generates a Nigerian bank account.
func (f *Faker) BankAccount(bankName string) (types.BankAccount, error) {
    var bank providers.Bank

	if bankName == "" {
		bank = providers.Banks.Banks[f.random(len(providers.Banks.Banks))]
	} else {
		found := false
		for _, b := range providers.Banks.Banks {
			if strings.EqualFold(b.Name, bankName) {
				bank = b
				found = true
				break
			}
		}
        if !found {
            return types.BankAccount{}, errors.New(
                errors.ErrInvalidBank,
                "bank not found: "+bankName,
            )
        }
    }

    accountNumber := fmt.Sprintf("%010d", f.random(9000000000)+1000000000)

    return types.BankAccount{
        BankName:      bank.Name,
        BankCode:      bank.Code,
        AccountNumber: accountNumber,
    }, nil
}

// BVN generates a random 11-digit Bank Verification Number.
func (f *Faker) BVN() string {
    return fmt.Sprintf("2%010d", f.random(9999999999)+1)
}

// NIN generates a random 11-digit National Identification Number.
func (f *Faker) NIN() string {
    return fmt.Sprintf("%011d", f.random(99999999999)+1)
}