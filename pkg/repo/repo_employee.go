package repo

import "github.com/Muha113/task-repo/pkg/models"

func (r *Repository) CreateNewEmployee(employee *models.Employee) error {
	return nil
}

func (r *Repository) UpdateEmployee(employee *models.Employee) error {
	return nil
}

func (r *Repository) FindEmployee(id int64) (Repository, error) {
	return Repository{}, nil
}

func (r *Repository) UpdateEmployeeWithFormData(employee *models.Employee) error {
	return nil
}

func (r *Repository) DeleteEmployee(id int64) error {
	return nil
}
