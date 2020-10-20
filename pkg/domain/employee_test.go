package domain_test

import (
	"fmt"
	"testing"

	"github.com/fernandoocampo/validator/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestEmployeeDataBasicValidator(t *testing.T) {
	// GIVEN
	testcases := map[string]struct {
		givenEmployee domain.Employee
		want          error
	}{
		"valid_employee": {
			givenEmployee: domain.Employee{
				UserName: "anyname",
				Email:    "anyname@gmail.com",
				Address: domain.Address{
					Street: "quinta",
					City:   "Cali",
					Zip:    "760001",
				},
			},
			want: nil,
		},
		"empty_username": {
			givenEmployee: domain.Employee{
				Email: "anyname@gmail.com",
				Address: domain.Address{
					Street: "quinta",
					City:   "Cali",
					Zip:    "760001",
				},
			},
			want: fmt.Errorf("username is mandatory"),
		},
		"empty_email": {
			givenEmployee: domain.Employee{
				UserName: "anyname",
				Address: domain.Address{
					Street: "quinta",
					City:   "Cali",
					Zip:    "760001",
				},
			},
			want: fmt.Errorf("given Email is not valid"),
		},
		"empty_street": {
			givenEmployee: domain.Employee{
				UserName: "anyname",
				Email:    "anyname@gmail.com",
				Address: domain.Address{
					City: "Cali",
					Zip:  "760001",
				},
			},
			want: fmt.Errorf("street is mandatory"),
		},
		"empty_city": {
			givenEmployee: domain.Employee{
				UserName: "anyname",
				Email:    "anyname@gmail.com",
				Address: domain.Address{
					Street: "quinta",
					Zip:    "760001",
				},
			},
			want: fmt.Errorf("city is mandatory"),
		},
		"invalid_zip": {
			givenEmployee: domain.Employee{
				UserName: "anyname",
				Email:    "anyname@gmail.com",
				Address: domain.Address{
					Street: "quinta",
					City:   "Cali",
					Zip:    "234d",
				},
			},
			want: fmt.Errorf("zip code is not a valid"),
		},
	}

	for testname, v := range testcases {
		err := v.givenEmployee.ValidateInABasicWay()
		assert.Equal(t, v.want, err, testname)
	}
}

func TestValidateInAFullBasicWay(t *testing.T) {
	// GIVEN
	testcases := map[string]struct {
		givenEmployee domain.Employee
		want          []error
	}{
		"valid_employee": {
			givenEmployee: domain.Employee{
				UserName: "anyname",
				Email:    "anyname@gmail.com",
				Address: domain.Address{
					Street: "quinta",
					City:   "Cali",
					Zip:    "760001",
				},
			},
			want: nil,
		},
		"empty_username": {
			givenEmployee: domain.Employee{
				Email: "anyname@gmail.com",
				Address: domain.Address{
					Street: "quinta",
					City:   "Cali",
					Zip:    "760001",
				},
			},
			want: []error{fmt.Errorf("username is mandatory")},
		},
		"empty_email": {
			givenEmployee: domain.Employee{
				UserName: "anyname",
				Address: domain.Address{
					Street: "quinta",
					City:   "Cali",
					Zip:    "760001",
				},
			},
			want: []error{fmt.Errorf("given Email is not valid")},
		},
		"empty_street": {
			givenEmployee: domain.Employee{
				UserName: "anyname",
				Email:    "anyname@gmail.com",
				Address: domain.Address{
					City: "Cali",
					Zip:  "760001",
				},
			},
			want: []error{fmt.Errorf("street is mandatory")},
		},
		"empty_city": {
			givenEmployee: domain.Employee{
				UserName: "anyname",
				Email:    "anyname@gmail.com",
				Address: domain.Address{
					Street: "quinta",
					Zip:    "760001",
				},
			},
			want: []error{fmt.Errorf("city is mandatory")},
		},
		"invalid_zip": {
			givenEmployee: domain.Employee{
				UserName: "anyname",
				Email:    "anyname@gmail.com",
				Address: domain.Address{
					Street: "quinta",
					City:   "Cali",
					Zip:    "234d",
				},
			},
			want: []error{fmt.Errorf("zip code is not a valid")},
		},
	}

	for testname, v := range testcases {
		got := v.givenEmployee.ValidateInAFullBasicWay()
		if v.want == nil && got == nil {
			continue
		}
		err, ok := got.(*domain.EmployeeValidationError)
		if !ok {
			t.Errorf("%q: expected EmployeeValidationError but got other: %s", testname, err)
			t.FailNow()
		}

		assert.Equal(t, v.want, err.Validations, testname)
	}
}
