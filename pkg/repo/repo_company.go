package repo

import "github.com/Muha113/task-repo/pkg/models"

func (r *Repository) CreateCompany(company *models.Company) error {
	return nil
}

func (r *Repository) UpdateCompany(company *models.Company) error {
	return nil
}

func (r *Repository) FindCompany(id int64) (models.Company, error) {
	return models.Company{}, nil
}

func (r *Repository) UpdateCompanyWithFormData(company *models.Company) error {
	return nil
}

func (r *Repository) DeleteCompany(id int64) error {
	return nil
}

func (r *Repository) GetCompanyEmployeesList(id int64) ([]models.Employee, error) {
	return nil, nil
}
