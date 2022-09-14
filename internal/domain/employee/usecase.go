package employee

type EmployeeUsecase struct {
	repo EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo EmployeeRepository) EmployeeUsecase {
	return EmployeeUsecase{
		repo: employeeRepo,
	}
}

func (u EmployeeUsecase) Create(em Employee) error {
	return u.repo.Create(em)
}

func (u EmployeeUsecase) Read() ([]Employee, error) {
	return u.repo.Read()
}

func (u EmployeeUsecase) Update(em Employee) error {
	return u.repo.Update(em)
}

func (u EmployeeUsecase) Delete(em Employee) error {
	return u.repo.Delete(em)
}
