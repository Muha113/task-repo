package repo

import (
	"github.com/Muha113/task-repo/pkg/apiproto"
)

func (r *Repository) CreateNewCompany(company *apiproto.Company) error {
	return nil
}

func (r *Repository) UpdateCompany(company *apiproto.Company) error {
	return nil
}

func (r *Repository) FindCompany(id int64) (*apiproto.Company, error) {
	return &apiproto.Company{}, nil
}

func (r *Repository) UpdateCompanyWithFormData(company *apiproto.Company) error {
	return nil
}

func (r *Repository) DeleteCompany(id int64) error {
	return nil
}

func (r *Repository) GetCompanyEmployeesList(id int64) ([]*apiproto.Employee, error) {
	return nil, nil
}
