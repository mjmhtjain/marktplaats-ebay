package models

type Creditor struct {
	Name     string  `csv:"Name"`
	Address  string  `csv:"Address"`
	Postcode string  `csv:"Postcode"`
	Phone    string  `csv:"Phone"`
	Credit   float64 `csv:"Credit"`
	Birthday string  `csv:"Birthday"`
}
