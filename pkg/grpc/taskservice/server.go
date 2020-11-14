package taskservice

import (
	"context"
	"fmt"

	"github.com/Muha113/task-repo/pkg/apiproto"
	"github.com/Muha113/task-repo/pkg/repo"
)

type TaskServiceServer struct {
	Repo *repo.Repository
}

func (tss *TaskServiceServer) AddNewEmployee(ctx context.Context, e *apiproto.Employee) (*apiproto.ResponseEmployee, error) {
	err := tss.Repo.CreateNewEmployee(e)

	status := 500
	code, ok := repo.StatusMapper[repo.ErrorStatus(err.Error())]
	if ok {
		status = code
	}

	return &apiproto.ResponseEmployee{
		Status:   int32(status),
		Employee: []*apiproto.Employee{},
	}, err
}

func (tss *TaskServiceServer) UpdateEmployee(ctx context.Context, e *apiproto.Employee) (*apiproto.ResponseEmployee, error) {
	err := tss.Repo.UpdateEmployee(e)

	status := 500
	code, ok := repo.StatusMapper[repo.ErrorStatus(err.Error())]
	if ok {
		status = code
	}

	return &apiproto.ResponseEmployee{
		Status:   int32(status),
		Employee: []*apiproto.Employee{},
	}, err
}

func (tss *TaskServiceServer) FindEmployee(ctx context.Context, e *apiproto.Employee) (*apiproto.ResponseEmployee, error) {
	empl, err := tss.Repo.FindEmployee(e.Id)

	status := 500
	code, ok := repo.StatusMapper[repo.ErrorStatus(err.Error())]
	if ok {
		status = code
	}

	return &apiproto.ResponseEmployee{
		Status:   int32(status),
		Employee: []*apiproto.Employee{empl},
	}, err
}

func (tss *TaskServiceServer) UpdateEmployeeWithFormData(ctx context.Context, e *apiproto.OptionalFields) (*apiproto.ResponseEmployee, error) {
	err := tss.Repo.UpdateEmployeeWithFormData(e)

	status := 500
	code, ok := repo.StatusMapper[repo.ErrorStatus(err.Error())]
	if ok {
		status = code
	}

	return &apiproto.ResponseEmployee{
		Status:   int32(status),
		Employee: []*apiproto.Employee{},
	}, err
}

func (tss *TaskServiceServer) DeleteEmployee(ctx context.Context, e *apiproto.Employee) (*apiproto.ResponseEmployee, error) {
	err := tss.Repo.DeleteEmployee(e.Id)

	status := 500
	code, ok := repo.StatusMapper[repo.ErrorStatus(err.Error())]
	if ok {
		status = code
	}

	return &apiproto.ResponseEmployee{
		Status:   int32(status),
		Employee: []*apiproto.Employee{},
	}, err
}

func (tss *TaskServiceServer) AddNewCompany(ctx context.Context, e *apiproto.Company) (*apiproto.ResponseCompany, error) {
	fmt.Println("Executing AddNewCompany() func...")

	comp, err := tss.Repo.CreateNewCompany(e)

	status := 500
	code, ok := repo.StatusMapper[repo.ErrorStatus(err.Error())]
	if ok {
		status = code
	}

	return &apiproto.ResponseCompany{
		Status:  int32(status),
		Company: []*apiproto.Company{comp},
	}, err
}

func (tss *TaskServiceServer) UpdateCompany(ctx context.Context, e *apiproto.Company) (*apiproto.ResponseCompany, error) {
	err := tss.Repo.UpdateCompany(e)

	status := 500
	code, ok := repo.StatusMapper[repo.ErrorStatus(err.Error())]
	if ok {
		status = code
	}

	return &apiproto.ResponseCompany{
		Status:  int32(status),
		Company: []*apiproto.Company{},
	}, err
}

func (tss *TaskServiceServer) FindCompany(ctx context.Context, e *apiproto.Company) (*apiproto.ResponseCompany, error) {
	comp, err := tss.Repo.FindCompany(e.Id)

	status := 500
	code, ok := repo.StatusMapper[repo.ErrorStatus(err.Error())]
	if ok {
		status = code
	}

	return &apiproto.ResponseCompany{
		Status:  int32(status),
		Company: []*apiproto.Company{comp},
	}, err
}

func (tss *TaskServiceServer) UpdateCompanyWithFormData(ctx context.Context, e *apiproto.OptionalFields) (*apiproto.ResponseCompany, error) {
	err := tss.Repo.UpdateCompanyWithFormData(e)

	status := 500
	code, ok := repo.StatusMapper[repo.ErrorStatus(err.Error())]
	if ok {
		status = code
	}

	return &apiproto.ResponseCompany{
		Status:  int32(status),
		Company: []*apiproto.Company{},
	}, err
}

func (tss *TaskServiceServer) DeleteCompany(ctx context.Context, e *apiproto.Company) (*apiproto.ResponseCompany, error) {
	err := tss.Repo.DeleteCompany(e.Id)

	status := 500
	code, ok := repo.StatusMapper[repo.ErrorStatus(err.Error())]
	if ok {
		status = code
	}

	return &apiproto.ResponseCompany{
		Status:  int32(status),
		Company: []*apiproto.Company{},
	}, err
}

func (tss *TaskServiceServer) GetCompanyEmployeesList(ctx context.Context, e *apiproto.Company) (*apiproto.ResponseEmployee, error) {
	empls, err := tss.Repo.GetCompanyEmployeesList(e.Id)

	status := 500
	code, ok := repo.StatusMapper[repo.ErrorStatus(err.Error())]
	if ok {
		status = code
	}

	return &apiproto.ResponseEmployee{
		Status:   int32(status),
		Employee: empls,
	}, err
}
