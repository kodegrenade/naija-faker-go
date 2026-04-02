package naijafakergo

import (
	"fmt"
	"time"

	"github.com/kodegrenade/naija-faker-go/internal/errors"
	"github.com/kodegrenade/naija-faker-go/internal/validate"
	"github.com/kodegrenade/naija-faker-go/providers"
	"github.com/kodegrenade/naija-faker-go/types"
)

// EducationRecord generates a Nigerian education record.
func (f *Faker) EducationRecord(lang string) (types.EducationRecord, error) {
	lang = f.resolveLanguage(lang)

	if err := validate.Language(lang); err != nil {
		return types.EducationRecord{}, err
	}

	university := f.University()
	degree := providers.Jobs.Degrees[f.random(len(providers.Jobs.Degrees))]
	course := f.pick(providers.Jobs.Courses)

	currentYear := time.Now().Year()
	graduationYear := f.random(15) + (currentYear - 20)

	return types.EducationRecord{
		University:     university.Name,
		Abbreviation:   university.Abbreviation,
		Degree:         degree.Code,
		Course:         course,
		GraduationYear: graduationYear,
	}, nil
}

// WorkRecord generates a Nigerian work record.
// level is optional — pass empty string for random level.
func (f *Faker) WorkRecord(level string) (types.WorkRecord, error) {
	if level == "" {
		level = f.pick([]string{"entry", "mid", "senior", "executive"})
	}

	if !validLevels[level] {
		return types.WorkRecord{}, errors.New(
			errors.ErrInvalidLevel,
			"level must be one of: entry, junior, mid, senior, lead, executive",
		)
	}

	company, err := f.Company()
	if err != nil {
		return types.WorkRecord{}, err
	}

	positions := providers.Jobs.Positions
	position := f.pick(positions)
	industry := f.pick(providers.Companies.CompanyIndustries)

	currentYear := time.Now().Year()
	startYear := f.random(10) + (currentYear - 12)

	return types.WorkRecord{
		Company:   company.Name,
		Position:  position,
		Industry:  industry,
		StartYear: startYear,
	}, nil
}

// VehicleRecord generates a Nigerian vehicle record.
// stateName is optional — used to generate a geographically correct license plate.
func (f *Faker) VehicleRecord(stateName string) (types.VehicleRecord, error) {
	plate, err := f.LicensePlate(stateName)
	if err != nil {
		return types.VehicleRecord{}, err
	}

	make_ := providers.Vehicles.VehicleMakes[f.random(len(providers.Vehicles.VehicleMakes))]
	model := f.pick(make_.Models)
	color := providers.Vehicles.VehicleColors[f.random(len(providers.Vehicles.VehicleColors))]

	currentYear := time.Now().Year()
	year := f.random(15) + (currentYear - 15)

	return types.VehicleRecord{
		LicensePlate: plate.Plate,
		Make:         make_.Make,
		Model:        model,
		Year:         year,
		Color:        color,
	}, nil
}

func (f *Faker) Company() (types.Company, error) {
	prefix := f.pick(providers.Companies.CompanyPrefixes)
	noun := f.pick(providers.Companies.CompanyNouns)
	suffix := f.pick(providers.Companies.CompanySuffixes)
	industry := f.pick(providers.Companies.CompanyIndustries)

	name := fmt.Sprintf("%s %s %s", prefix, noun, suffix)
	rcNumber := fmt.Sprintf("RC-%07d", f.random(9000000)+1000000)

	return types.Company{
		Name:     name,
		RCNumber: rcNumber,
		Industry: industry,
	}, nil
}

// University returns a random Nigerian university.
func (f *Faker) University() types.University {
	u := providers.Universities.Universities[f.random(len(providers.Universities.Universities))]
	return types.University{
		Name:         u.Name,
		Abbreviation: u.Abbreviation,
		State:        u.State,
		Type:         u.Type,
	}
}

// LicensePlate generates a Nigerian license plate.
func (f *Faker) LicensePlate(stateName string) (types.LicensePlate, error) {
	var code string

	if stateName == "" {
		codes := make([]string, 0, len(providers.Plates.StateCodes))
		for _, v := range providers.Plates.StateCodes {
			codes = append(codes, v)
		}
		code = f.pick(codes)
	} else {
		var ok bool
		code, ok = providers.Plates.StateCodes[stateName]
		if !ok {
			return types.LicensePlate{}, errors.New(
				errors.ErrInvalidState,
				"state not found: "+stateName,
			)
		}
	}

	digits := fmt.Sprintf("%03d", f.random(999)+1)
	letters := string([]byte{
		byte('A' + f.random(26)),
		byte('A' + f.random(26)),
	})

	fullPlateNumber := fmt.Sprintf("%s-%s%s", code, digits, letters)

	return types.LicensePlate{Plate: fullPlateNumber}, nil
}
