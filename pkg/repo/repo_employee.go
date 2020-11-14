package repo

import (
	"context"
	"fmt"

	"github.com/Muha113/task-repo/pkg/apiproto"
	"github.com/Muha113/task-repo/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var CompanyPositions = []string{"developer", "manager", "quality assurance", "business analyst"}

func (r *Repository) CreateNewEmployee(employee *apiproto.Employee) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.isExist(CollectionEmployee, "id", employee.Id) ||
		!r.isExist(CollectionCompany, "id", employee.CompanyId) ||
		!utils.Contains(CompanyPositions, employee.Position) {
		return fmt.Errorf("%s", ErrorInvalidInput)
	}

	_, err := r.repo.Collection(CollectionEmployee).InsertOne(context.Background(), employee)
	return err
}

func (r *Repository) UpdateEmployee(employee *apiproto.Employee) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.isExist(CollectionEmployee, "id", employee.Id) ||
		!utils.Contains(CompanyPositions, employee.Position) {
		return fmt.Errorf("%s", ErrorInvalidInput)
	}

	filter := bson.M{"id": employee.Id}
	_, err := r.repo.Collection(CollectionEmployee).ReplaceOne(context.Background(), filter, employee)
	return err
}

func (r *Repository) FindEmployee(id int64) (*apiproto.Employee, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.isExist(CollectionEmployee, "id", id) {
		return nil, fmt.Errorf("%s", ErrorNotFound)
	}

	filter := bson.M{"id": id}
	repoResult := r.repo.Collection(CollectionEmployee).FindOne(context.Background(), filter)

	var result *apiproto.Employee
	err := repoResult.Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) UpdateEmployeeWithFormData(employee *apiproto.OptionalFields) error {
	// r.mutex.Lock()
	// defer r.mutex.Unlock()

	// mapped := map[string]interface{}{}
	// err := json.Unmarshal([]byte(employee.Data), &mapped)
	// if err != nil {
	// 	return err
	// }

	// var decoded *apiproto.Employee
	// decoded, err = json.Unmarshal([]byte(employee.Data), decoded)
	// if err != nil {
	// 	return err
	// }

	// if !r.isExist(CollectionEmployee, "id", decoded.Id) {
	// 	return fmt.Errorf("%s", ErrorInvalidInput)
	// }

	// if utils.Contains(used, "position") && !utils.Contains(CompanyPositions, decoded.Position) {
	// 	return fmt.Errorf("%s", ErrorInvalidInput)
	// }

	// used := make([]string, 0, 7)
	// for k := range mapped {
	// 	if strings.Compare(k, "id") != 0 {
	// 		used = append(used, k)
	// 	}
	// }

	// tmp := structs.Map(decoded)
	// toUpdate := map[string]interface{}

	// for _, v := range used {
	// 	toUpdate[] = tmp[v]
	// }

	// update := bson.D{{Key: "$set", Value: bson.D{}}
	// filter := bson.D{"id": decoded.Id}
	// _, err := r.repo.Collection(CollectionEmployee).UpdateOne(context.Background(), filter, employee)
	return nil
}

func (r *Repository) DeleteEmployee(id int64) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.isExist(CollectionEmployee, "id", id) {
		return fmt.Errorf("%s", ErrorNotFound)
	}

	filter := bson.M{"id": id}
	_, err := r.repo.Collection(CollectionEmployee).DeleteOne(context.Background(), filter)

	return err
}
