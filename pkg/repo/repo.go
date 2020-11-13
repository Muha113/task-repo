package repo

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CollectionEmployee = "Employee"
	CollectionCompany  = "Company"
)

type Repository struct {
	mutex  *sync.RWMutex
	config *RepositoryConfig
	repo   *mongo.Database
	size   int64
}

func NewRepository(repoConfigPath string) (*Repository, error) {
	repoCfg, err := InitRepoConfig(repoConfigPath)
	if err != nil {
		return nil, err
	}

	connectionStr := fmt.Sprintf("%s%s:%s", repoCfg.ConnPrefix, repoCfg.ConnHost, repoCfg.ConnPort)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionStr))
	if err != nil {
		return nil, err
	}

	repo := client.Database(repoCfg.RepoName)

	return &Repository{
		mutex:  &sync.RWMutex{},
		config: repoCfg,
		repo:   repo,
		size:   0,
	}, nil
}
