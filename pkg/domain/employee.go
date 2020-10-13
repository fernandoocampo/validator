package domain

import (
	"errors"

	"github.com/adrg/postcode"
	"github.com/goware/emailx"
)

// Employee contains important data for employees.
type Employee struct {
	UserName string
	Email    string
	Address  Address
}

// Address contains data related to the particulars of the place
// where the someone lives
type Address struct {
	Street string
	City   string
	Zip    string
}

// ValidateInABasicWay validates an employee in a basic way
func (e Employee) ValidateInABasicWay() error {
	if e.UserName == "" {
		return errors.New("username is mandatory")
	}
	emailErr := emailx.Validate(e.Email)
	if emailErr != nil {
		return errors.New("given Email is not valid")
	}
	addressErr := e.Address.ValidateInABasicWay()
	if addressErr != nil {
		return addressErr
	}
	return nil
}

// ValidateInABasicWay validates an address in a basic way
func (a Address) ValidateInABasicWay() error {
	if a.Street == "" {
		return errors.New("street is mandatory")
	}
	if a.City == "" {
		return errors.New("city is mandatory")
	}

	if a.Zip == "" {
		return nil
	}

	err := postcode.Validate(a.Zip)
	if err != nil {
		return errors.New("zip code is not a valid")
	}
	return nil
}
