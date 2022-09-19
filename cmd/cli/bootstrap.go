package cli

import (
	"database/sql"

	emDomain "github.com/anggiatyoga/hris-api/internal/domain/employee"
	app "github.com/anggiatyoga/hris-api/internal/platform"
	emRepo "github.com/anggiatyoga/hris-api/internal/platform/database"
)

func Bootstrap(d *sql.DB, config Config) (app.AppModule, error) {
	// Employee Modules
	em := emRepo.NewRepository(d)
	emUC := emDomain.NewEmployeeUsecase(em)

	appModule := app.AppModule{
		EmployeeModules: emUC,
	}

	return appModule, nil
}
