package taskservice

import (
	"context"

	"github.com/Muha113/task-repo/pkg/apiproto"
	"github.com/Muha113/task-repo/pkg/repo"
)

type TaskServiceServer struct {
	repo *repo.Repository
}

func (tss *TaskServiceServer) AddNewEmployee(ctx context.Context, e *apiproto.Employee) (*apiproto.ResponseEmployee, error) {
	err := tss.repo.CreateNewEmployee(e)
	return &apiproto.ResponseEmployee{
		Status:   0,
		Employee: nil,
	}, err
}

func (tss *TaskServiceServer) UpdateEmployee(ctx context.Context, e *apiproto.Employee) (*apiproto.ResponseEmployee, error) {
	err := tss.repo.UpdateEmployee(e)
	return &apiproto.ResponseEmployee{
		Status:   0,
		Employee: nil,
	}, err
}

func (tss *TaskServiceServer) FindEmployee(ctx context.Context, e *apiproto.Employee) (*apiproto.ResponseEmployee, error) {
	empl, err := tss.repo.FindEmployee(e.Id)
	return &apiproto.ResponseEmployee{
		Status:   0,
		Employee: []*apiproto.Employee{empl},
	}, err
}

func (tss *TaskServiceServer) UpdateEmployeeWithFormData(ctx context.Context, e *apiproto.Employee) (*apiproto.ResponseEmployee, error) {
	err := tss.repo.UpdateEmployeeWithFormData(e)
	return &apiproto.ResponseEmployee{
		Status:   0,
		Employee: nil,
	}, err
}

func (tss *TaskServiceServer) DeleteEmployee(ctx context.Context, e *apiproto.Employee) (*apiproto.ResponseEmployee, error) {
	err := tss.repo.DeleteEmployee(e.Id)
	return &apiproto.ResponseEmployee{
		Status:   0,
		Employee: nil,
	}, err
}

func (tss *TaskServiceServer) AddNewCompany(ctx context.Context, e *apiproto.Company) (*apiproto.ResponseCompany, error) {
	err := tss.repo.CreateNewCompany(e)
	return &apiproto.ResponseCompany{
		Status:  0,
		Company: nil,
	}, err
}

func (tss *TaskServiceServer) UpdateCompany(ctx context.Context, e *apiproto.Company) (*apiproto.ResponseCompany, error) {
	err := tss.repo.UpdateCompany(e)
	return &apiproto.ResponseCompany{
		Status:  0,
		Company: nil,
	}, err
}

func (tss *TaskServiceServer) FindCompany(ctx context.Context, e *apiproto.Company) (*apiproto.ResponseCompany, error) {
	comp, err := tss.repo.FindCompany(e.Id)
	return &apiproto.ResponseCompany{
		Status:  0,
		Company: []*apiproto.Company{comp},
	}, err
}

func (tss *TaskServiceServer) UpdateCompanyWithFormData(ctx context.Context, e *apiproto.Company) (*apiproto.ResponseCompany, error) {
	err := tss.repo.UpdateCompanyWithFormData(e)
	return &apiproto.ResponseCompany{
		Status:  0,
		Company: nil,
	}, err
}

func (tss *TaskServiceServer) DeleteCompany(ctx context.Context, e *apiproto.Company) (*apiproto.ResponseCompany, error) {
	err := tss.repo.DeleteCompany(e.Id)
	return &apiproto.ResponseCompany{
		Status:  0,
		Company: nil,
	}, err
}

func (tss *TaskServiceServer) GetCompanyEmployeesList(ctx context.Context, e *apiproto.Company) (*apiproto.ResponseEmployee, error) {
	empls, err := tss.repo.GetCompanyEmployeesList(e.Id)
	return &apiproto.ResponseEmployee{
		Status:   0,
		Employee: empls,
	}, err
}
