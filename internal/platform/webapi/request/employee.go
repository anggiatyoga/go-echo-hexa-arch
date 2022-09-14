package request

import "github.com/anggiatyoga/hris-api/internal/domain/employee"

type EmployeeParam struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (p EmployeeParam) ToEntity() employee.Employee {
	return employee.Employee{
		Name:    p.Name,
		Address: p.Address,
		Phone:   p.Phone,
	}
}
