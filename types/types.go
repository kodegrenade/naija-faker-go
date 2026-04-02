package types

type Person struct {
	Title     string
	FirstName string
	LastName  string
	FullName  string
	Email     string
	Phone     string
	Address   string
}

type DetailedPerson struct {
	Person
	State         string
	LGA           string
	DateOfBirth   DateOfBirth
	MaritalStatus string
	BloodGroup    string
	Genotype      string
	Salary        Salary
	NextOfKin     NextOfKin
	Education     EducationRecord
	Work          WorkRecord
	Vehicle       VehicleRecord
}

type ConsistentPerson struct {
	Person
	State string
	LGA   string
}

type BankAccount struct {
	BankName      string
	BankCode      string
	AccountNumber string
}

type Salary struct {
	Amount    int
	Currency  string
	Level     string
	Frequency string
}

type DateOfBirth struct {
	Date string
	Age  int
}

type NextOfKin struct {
	FullName     string
	Relationship string
	Phone        string
	Address      string
}

type EducationRecord struct {
	University     string
	Abbreviation   string
	Degree         string
	Course         string
	GraduationYear int
}

type WorkRecord struct {
	Company   string
	Position  string
	Industry  string
	StartYear int
}

type VehicleRecord struct {
	LicensePlate string
	Make         string
	Model        string
	Year         int
	Color        string
}

type Company struct {
	Name     string
	RCNumber string
	Industry string
}

type University struct {
	Name         string
	Abbreviation string
	State        string
	Type         string
}

type LicensePlate struct {
	Plate string
}
