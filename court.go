package main

type Court struct {
	Id           int
	Address      string
	Municipality string

	PhoneNumber  string
	OpenTimes    string
	Description  string
	PaymentTypes []string
}

// func (c *Court) String() string{
//
// }

type Courts []Court
