package employee

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/anggiatyoga/hris-api/internal/domain/employee"
)

type Repository struct {
	DB *sql.DB
}

const (
	tabelEmployee = "pegawai"
)

func NewRepository(d *sql.DB) Repository {
	return Repository{
		DB: d,
	}
}

func (r *Repository) Close() {
	r.DB.Close()
}

func (r Repository) Create(employee employee.Employee) error {
	// err := r.DB.Insert(employee)
	q := "INSERT INTO employee (name, address, phone) VALUES ($1, $2, $3)"

	s, err := r.DB.Exec(q, employee.Name, employee.Address, employee.Phone)

	if err != nil {
		fmt.Printf("Failed Create info: %s", err.Error())
		return err
	}

	count, err := s.RowsAffected()
	if err != nil {
		fmt.Printf("Failed Create query info: %s", err.Error())
		return err
	}

	if count == 0 {
		fmt.Printf("Failed Create 0 row affected")
		return errors.New("0 row affectedd")
	}

	return nil
}

func (r Repository) Read() (result []employee.Employee, err error) {
	// var result []employee.Employee
	// err := r.DB.Model(&result).Select()

	var obj employee.Employee
	q := "SELECT id, name, address, phone FROM employee"

	rows, err := r.DB.Query(q)

	if err != nil {
		fmt.Printf("Failed read data info: %s", err.Error())
		return
	}

	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.Name, &obj.Address, &obj.Phone)
		if err != nil {
			return
		}
		result = append(result, obj)
	}
	defer rows.Close()

	return result, nil
}

func (r Repository) Update(employee employee.Employee) error {
	q := "UPDATE employee SET name=$1, address=$2, phone=$3 WHERE id=$4"

	s, err := r.DB.Exec(q, employee.Name, employee.Address, employee.Phone, employee.ID)
	if err != nil {
		fmt.Printf("Failed update data info: %s", err.Error())
		return err
	}

	count, err := s.RowsAffected()
	if err != nil || count == 0 {
		fmt.Printf("Failed Update query 0 affected")
		return errors.New("Update query 0 affected")
	}

	return nil
}

func (r Repository) Delete(employee employee.Employee) error {
	q := "DELETE FROM employee WHERE id=$1"

	s, err := r.DB.Exec(q, employee.ID)
	if err != nil {
		fmt.Printf("Failed delete data info: %s", err.Error())
		return err
	}

	count, err := s.RowsAffected()
	if err != nil || count == 0 {
		fmt.Printf("Failed Delete query 0 affected")
		return errors.New("Delete query 0 affected")
	}

	return nil
}
