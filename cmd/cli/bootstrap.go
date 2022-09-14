package cli

import (
	emDomain "github.com/anggiatyoga/hris-api/internal/domain/employee"
	app "github.com/anggiatyoga/hris-api/internal/platform"
	emRepo "github.com/anggiatyoga/hris-api/internal/platform/database"
	"github.com/go-pg/pg"
)

func Bootstrap(d *pg.DB, config Config) (app.AppModule, error) {
	// Employee Modules
	em := emRepo.NewRepository(d)
	emUC := emDomain.NewEmployeeUsecase(em)

	appModule := app.AppModule{
		EmployeeModules: emUC,
	}

	return appModule, nil
}
