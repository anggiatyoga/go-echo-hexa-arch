package request

import "github.com/anggiatyoga/hris-api/internal/domain/employee"

type EmployeeParam struct {
	ID      int    `json:"id" validate:"required,number"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (p EmployeeParam) ToEntity() employee.Employee {
	return employee.Employee{
		ID:      p.ID,
		Name:    p.Name,
		Address: p.Address,
		Phone:   p.Phone,
	}
}
