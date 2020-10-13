package portin

import "github.com/fernandoocampo/validator/pkg/domain"

// EmployeeService defines behavior for employee business logic.
type EmployeeService interface {
	// Create creates a new employee.
	Create(newEmployee domain.Employee) error
}
