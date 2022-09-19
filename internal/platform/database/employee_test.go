package employee

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/anggiatyoga/hris-api/internal/domain/employee"
	"github.com/stretchr/testify/assert"
)

var e = &employee.Employee{
	ID:      1,
	Name:    "Budi",
	Address: "Jakarta",
	Phone:   "08123",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestEmployee(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	t.Run("read success", func(t *testing.T) {
		query := "SELECT id, name, address, phone FROM employee"
		rows := sqlmock.NewRows([]string{"id", "name", "address", "phone"}).
			AddRow(e.ID, e.Name, e.Address, e.Phone)

		mock.ExpectQuery(query).WillReturnRows(rows)
		employee, err := repo.Read()
		assert.NotEmpty(t, employee)
		assert.NoError(t, err)
		assert.Len(t, employee, 1)
	})

	t.Run("create success", func(t *testing.T) {
		query := "INSERT INTO employee \\(name, address, phone\\) VALUES \\(\\$1, \\$2, \\$3\\)"

		mock.ExpectExec(query).WithArgs(e.Name, e.Address, e.Phone).WillReturnResult(sqlmock.NewResult(1, 1))
		err := repo.Create(*e)
		assert.NoError(t, err)
	})

	t.Run("create failed", func(t *testing.T) {
		query := "INSERT INTO employee \\(name, address, phone\\) VALUES \\(\\$1, \\$2, \\$3\\)"

		mock.ExpectExec(query).WithArgs(e.Name, e.Address, e.Phone).WillReturnResult(sqlmock.NewResult(0, 0))
		err := repo.Create(*e)
		assert.Error(t, err)
	})

	t.Run("update success", func(t *testing.T) {
		query := "UPDATE employee SET name=\\$1, address=\\$2, phone=\\$3 WHERE id=\\$4"

		mock.ExpectExec(query).WithArgs(e.Name, e.Address, e.Phone, e.ID).WillReturnResult(sqlmock.NewResult(0, 1))
		err := repo.Update(*e)
		assert.NoError(t, err)
	})

	t.Run("update failed", func(t *testing.T) {
		query := "UPDATE employee SET name=\\$1, address=\\$2, phone=\\$3 WHERE id=\\$4"

		mock.ExpectExec(query).WithArgs(e.Name, e.Address, e.Phone, e.ID).WillReturnResult(sqlmock.NewResult(0, 0))
		err := repo.Update(*e)
		assert.Error(t, err)
	})

	t.Run("delete success", func(t *testing.T) {
		query := "DELETE FROM employee WHERE id=\\$1"

		mock.ExpectExec(query).WithArgs(e.ID).WillReturnResult(sqlmock.NewResult(0, 1))
		err := repo.Delete(*e)
		assert.NoError(t, err)
	})

	t.Run("delete failed", func(t *testing.T) {
		query := "DELETE FROM employee WHERE id=\\$1"

		mock.ExpectExec(query).WithArgs(e.ID).WillReturnResult(sqlmock.NewResult(0, 0))
		err := repo.Delete(*e)
		assert.Error(t, err)
	})
}
