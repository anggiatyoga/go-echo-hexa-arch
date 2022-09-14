package employee

import (
	"fmt"

	"github.com/anggiatyoga/hris-api/internal/domain/employee"
	"github.com/go-pg/pg"
)

type Repository struct {
	DB *pg.DB
}

const (
	tabelEmployee = "pegawai"
)

func NewRepository(d *pg.DB) Repository {
	return Repository{
		DB: d,
	}
}

func (r Repository) Create(employee employee.Employee) error {
	err := r.DB.Insert(employee)
	return err
}

func (r Repository) Read() ([]employee.Employee, error) {
	var result []employee.Employee
	err := r.DB.Model(&result).Select()

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r Repository) Update(employee employee.Employee) error {
	_, err := r.DB.Model().Table(tabelEmployee).Set("nama=?", employee.Name).
		Set("alamat=?", employee.Address).
		Set("telepon=?", employee.Phone).
		Update()
	return err
}

func (r Repository) Delete(employee employee.Employee) error {
	res, err := r.DB.Model(&employee).Where("id=?", employee.ID).Delete()
	if err != nil {
		fmt.Printf("Failed delete data info: %s", err.Error())
	} else {
		fmt.Printf("Success delete total record: %d id: %v", res.RowsAffected(), employee.ID)
	}
	return err
}
