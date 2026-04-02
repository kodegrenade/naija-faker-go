package naijafakergo

import (
	"fmt"
	"time"

	"github.com/kodegrenade/naija-faker-go/providers"
    "github.com/kodegrenade/naija-faker-go/types"
)

var maritalStatuses = [5]string{"Single", "Married", "Divorced", "Widowed", "Separated"}

// DateOfBirth generates a Nigerian date of birth.
func (f *Faker) DateOfBirth(minAge, maxAge int) types.DateOfBirth {
    if minAge <= 0 {
        minAge = 18
    }
    if maxAge <= 0 || maxAge <= minAge {
        maxAge = 80
    }

    age := f.random(maxAge-minAge+1) + minAge
    year := time.Now().Year() - age

    month := f.random(12) + 1
    day := f.random(28) + 1

    return types.DateOfBirth{
        Date: fmt.Sprintf("%d-%02d-%02d", year, month, day),
        Age:  age,
    }
}

// MaritalStatus returns a random marital status.
func (f *Faker) MaritalStatus() string {
    return f.pick(maritalStatuses[:])
}

// BloodGroup returns a random blood group.
func (f *Faker) BloodGroup() string {
    return f.pick(providers.Medical.BloodGroups)
}

// Genotype returns a random genotype.
func (f *Faker) Genotype() string {
    return f.pick(providers.Medical.Genotypes)
}