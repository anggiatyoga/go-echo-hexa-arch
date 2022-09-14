package employee

type EmployeeRepository interface {
	Create(em Employee) error
	Read() ([]Employee, error)
	Update(employee Employee) error
	Delete(em Employee) error
}
