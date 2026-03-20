package providers

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed addresses.json
var addressesJSON []byte

//go:embed banks.json
var banksJSON []byte

//go:embed companies.json
var companiesJSON []byte

//go:embed emails.json
var emailsJSON []byte

//go:embed geo.json
var geoJSON []byte

//go:embed jobs.json
var jobsJSON []byte

//go:embed locations.json
var locationsJSON []byte

//go:embed medical.json
var medicalJSON []byte

//go:embed names.json
var namesJSON []byte

//go:embed numbers.json
var numbersJSON []byte

//go:embed plates.json
var platesJSON []byte

//go:embed salaries.json
var salariesJSON []byte

//go:embed titles.json
var titlesJSON []byte

//go:embed universities.json
var universitiesJSON []byte

//go:embed vehicles.json
var vehiclesJSON []byte

type AddressesData struct {
	StreetTypes    []string            `json:"types"`
	StreetPrefixes []string            `json:"prefixes"`
	Locale         []string            `json:"locale"`
	Places         map[string][]string `json:"places"`
}

type BanksData struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CompaniesData struct {
	CompanySuffixes   []string `json:"companySuffixes"`
	CompanyPrefixes   []string `json:"companyPrefixes"`
	CompanyNouns      []string `json:"companyNouns"`
	CompanyIndustries []string `json:"industries"`
}

type EmailsData struct {
	Extensions []string `json:"extensions"`
}

type GeoData struct {
	RegionMap map[string][]Region `json:"regionMap"`
	StateLgas map[string][]string `json:"stateLgas"`
}

type Region struct {
	Language []string `json:"language"`
	States   []string `json:"states"`
}

type JobsData struct {
	Positions []string `json:"positions"`
	Degrees   []string `json:"degrees"`
	Courses   []string `json:"courses"`
}

type LocationsData struct {
	States []string `json:"states"`
	Lgas   []string `json:"lgas"`
}

type MedicalData struct {
	BloodGroups []string `json:"bloodGroups"`
	Genotypes   []string `json:"genotypes"`
}

type NamesData struct {
	Yoruba map[string][]string `json:"yoruba"`
	Igbo   map[string][]string `json:"igbo"`
	Hausa  map[string][]string `json:"hausa"`
}

type NumbersData struct {
	Networks []string            `json:"networks"`
	Prefix   map[string][]string `json:"prefix"`
}

type PlatesData struct {
	StateCodes map[string]string `json:"stateCodes"`
}

type SalariesData struct {
	SalaryBands map[string][]SalaryBand `json:"salaryBands"`
}

type SalaryBand struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type TitlesData struct {
	Male   string `json:"male"`
	Female string `json:"female"`
}

type University struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	State        string `json:"state"`
	Type         string `json:"type"`
}

type UniversitiesData struct {
	Universities []University `json:"universities"`
}

type VehichlesData struct {
	VehicleMakes  []VehicleMake `json:"vehicleMakes"`
	VehicleColors []string      `json:"vehicleColors"`
}

type VehicleMake struct {
	Make   string `json:"make"`
	Models []string
}

var (
	Addresses    AddressesData
	Banks        BanksData
	Companies    CompaniesData
	Emails       EmailsData
	Geo          GeoData
	Jobs         JobsData
	Locations    LocationsData
	Medical      MedicalData
	Names        NamesData
	Numbers      NumbersData
	Plates       PlatesData
	Salaries     SalariesData
	Titles       TitlesData
	Universities UniversitiesData
	Vehicles     VehichlesData
)

func init() {
	mustUnmarshal(addressesJSON, &Addresses, "addresses.json")
	mustUnmarshal(banksJSON, &Banks, "banks.json")
	mustUnmarshal(companiesJSON, &Companies, "companies.json")
	mustUnmarshal(emailsJSON, &Emails, "emails.json")
	mustUnmarshal(geoJSON, &Geo, "geo.json")
	mustUnmarshal(jobsJSON, &Jobs, "jobs.json")
	mustUnmarshal(locationsJSON, &Locations, "locations.json")
	mustUnmarshal(medicalJSON, &Medical, "medical.json")
	mustUnmarshal(namesJSON, &Names, "names.json")
	mustUnmarshal(numbersJSON, &Numbers, "numbers.json")
	mustUnmarshal(platesJSON, &Plates, "plates.json")
	mustUnmarshal(salariesJSON, &Salaries, "salaries.json")
	mustUnmarshal(titlesJSON, &Titles, "titles.json")
	mustUnmarshal(universitiesJSON, &Universities, "universities.json")
	mustUnmarshal(vehiclesJSON, &Vehicles, "vehicles.json")
}

func mustUnmarshal(b []byte, v any, filename string) {
	if err := json.Unmarshal(b, v); err != nil {
		log.Fatalf("naija-faker: failed to load %s: %v", filename, err)
	}
}
