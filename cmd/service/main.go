package main

import (
	"net"

	"github.com/Muha113/task-repo/pkg/apiproto"
	"github.com/Muha113/task-repo/pkg/grpc/taskservice"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	// repo, err := repo.NewRepository("./configs/repo_config_debug.json")
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	// e := &models.Employee{
	// 	ID:         1,
	// 	Name:       "lox",
	// 	SecondName: "pidr",
	// 	Surname:    "sasi",
	// 	HireDate:   time.Now(),
	// 	Position:   "eblan",
	// 	CompanyID:  1,
	// }

	// err = repo.CreateNewEmployee(e)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	s := grpc.NewServer()
	srv := &taskservice.TaskServiceServer{}
	apiproto.RegisterTaskRepoServer(s, srv)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		logrus.Fatal(err)
	}

	if err = s.Serve(l); err != nil {
		logrus.Fatal(err)
	}
}
