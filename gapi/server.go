package gapi

import (
	"fmt"

	db "github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/db/sqlc"
	"github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/pb"
	"github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/token"
	"github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/util"
	"github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
