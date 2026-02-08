package db

type EmployeeDetails struct {
	ID         string  `json: "id"`
	Name       string  `json: "name"`
	Department string  `json: "department"`
	Address    Address `json: "address"`
}

type Address struct {
	Street string `json: "street"`
	City   string `json: "city"`
	State  string `json: "state"`
}

var Employees map[string]EmployeeDetails = map[string]EmployeeDetails{
	"add120": {ID: "user-01", Name: "Akama Idion", Department: "Finance", Address: Address{Street: "1 Jude street", City: "Hamona", State: "Riela"}},
	"add132": {ID: "user-02", Name: "Abek John", Department: "Logistics", Address: Address{Street: "1 Kogo street", City: "Tamo", State: "Riela"}},
	"add143": {ID: "user-03", Name: "Rebecca Kijoremi", Department: "Administration", Address: Address{Street: "1 Dramo street", City: "Lokongoma", State: "Kogi"}},
	"add159": {ID: "user-04", Name: "Philemon Oti", Department: "Production", Address: Address{Street: "1 Julani street", City: "Tiv", State: "Benue"}},
	"add127": {ID: "user-05", Name: "Kasandra Onu", Department: "Secretariat", Address: Address{Street: "1 kapsula street", City: "Huga", State: "Dresa"}},
}
