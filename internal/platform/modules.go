package platform

import "github.com/anggiatyoga/hris-api/internal/domain/employee"

type AppModule struct {
	EmployeeModules employee.EmployeeUsecase
}
