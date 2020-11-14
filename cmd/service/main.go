package main

import (
	"net"

	"github.com/Muha113/task-repo/pkg/apiproto"
	"github.com/Muha113/task-repo/pkg/grpc/taskservice"
	"github.com/Muha113/task-repo/pkg/repo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	repo, err := repo.NewRepository("./configs/repo_config_debug.json")
	if err != nil {
		logrus.Fatal(err)
	}

	s := grpc.NewServer()
	srv := &taskservice.TaskServiceServer{
		Repo: repo,
	}
	apiproto.RegisterTaskRepoServer(s, srv)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		logrus.Fatal(err)
	}

	if err = s.Serve(l); err != nil {
		logrus.Fatal(err)
	}
}
