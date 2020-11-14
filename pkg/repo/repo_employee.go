package repo

import (
	"github.com/Muha113/task-repo/pkg/apiproto"
)

func (r *Repository) CreateNewEmployee(employee *apiproto.Employee) error {
	// _, err := r.repo.Collection(CollectionEmployee).InsertOne(context.Background(), employee)
	return nil
}

func (r *Repository) UpdateEmployee(employee *apiproto.Employee) error {
	return nil
}

func (r *Repository) FindEmployee(id int64) (*apiproto.Employee, error) {
	return &apiproto.Employee{}, nil
}

func (r *Repository) UpdateEmployeeWithFormData(employee *apiproto.Employee) error {
	return nil
}

func (r *Repository) DeleteEmployee(id int64) error {
	return nil
}
