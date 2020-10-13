package application

import (
	"github.com/fernandoocampo/validator/pkg/domain"
	"github.com/fernandoocampo/validator/pkg/portin"
)

// Employee implements behavior for employee applicatio logic
type Employee struct {
}

// NewEmployeeService creates a new employee service
func NewEmployeeService() portin.EmployeeService {
	return &Employee{}
}

// Create creates a new employee and store it in a database.
func (e *Employee) Create(newEmployee domain.Employee) error {
	return nil
}
