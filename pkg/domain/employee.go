package domain

import (
	"errors"

	"github.com/adrg/postcode"
	"github.com/goware/emailx"
)

// EmployeeValidationError contains employee validation errors
type EmployeeValidationError struct {
	Validations []error
}

func (e *EmployeeValidationError) Error() string {
	if len(e.Validations) == 1 {
		return e.Validations[0].Error()
	}
	return "employee doesn't contain valid data"
}

// AddressValidationError contains address validation errors
type AddressValidationError struct {
	Validations []error
}

func (a *AddressValidationError) Error() string {
	if len(a.Validations) == 1 {
		return a.Validations[0].Error()
	}
	return "address doesn't contain valid data"
}

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

// ValidateInAFullBasicWay validates an employee in a basic way
// returning all the possible errors.
func (e Employee) ValidateInAFullBasicWay() error {
	var result []error
	if e.UserName == "" {
		result = append(result, errors.New("username is mandatory"))
	}
	emailErr := emailx.Validate(e.Email)
	if emailErr != nil {
		result = append(result, errors.New("given Email is not valid"))
	}
	addressErr := e.Address.ValidateInAFullBasicWay()
	if addressErr != nil {
		if err, ok := addressErr.(*AddressValidationError); ok {
			result = append(result, err.Validations...)
		} else {
			result = append(result, addressErr)
		}
	}
	if len(result) > 0 {
		return &EmployeeValidationError{
			Validations: result,
		}
	}
	return nil
}

// ValidateInAFullBasicWay validates an address in a basic way
// returning all the possible errors.
func (a Address) ValidateInAFullBasicWay() error {
	var result []error
	if a.Street == "" {
		result = append(result, errors.New("street is mandatory"))
	}
	if a.City == "" {
		result = append(result, errors.New("city is mandatory"))
	}

	if a.Zip != "" {
		err := postcode.Validate(a.Zip)
		if err != nil {
			result = append(result, errors.New("zip code is not a valid"))
		}
	}

	if len(result) > 0 {
		return &AddressValidationError{
			Validations: result,
		}
	}
	return nil
}
