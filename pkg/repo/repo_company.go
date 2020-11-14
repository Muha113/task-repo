package repo

import (
	"context"
	"fmt"

	"github.com/Muha113/task-repo/pkg/apiproto"
	"github.com/Muha113/task-repo/pkg/utils"
)

var LegalForm = []string{"OOO", "ZAO", "IP"}

func (r *Repository) CreateNewCompany(company *apiproto.Company) (*apiproto.Company, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.isExist(CollectionCompany, "id", company.Id) ||
		!utils.Contains(LegalForm, company.LegalForm) {
		return nil, fmt.Errorf("%s", ErrorInvalidInput)
	}

	_, err := r.repo.Collection(CollectionCompany).InsertOne(context.Background(), company)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (r *Repository) UpdateCompany(company *apiproto.Company) error {
	return nil
}

func (r *Repository) FindCompany(id int64) (*apiproto.Company, error) {
	return &apiproto.Company{}, nil
}

func (r *Repository) UpdateCompanyWithFormData(company *apiproto.OptionalFields) error {
	return nil
}

func (r *Repository) DeleteCompany(id int64) error {
	return nil
}

func (r *Repository) GetCompanyEmployeesList(id int64) ([]*apiproto.Employee, error) {
	return nil, nil
}
