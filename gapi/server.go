package gapi

import (
	"fmt"

	db "github.com/alim7007/go_bank_k8s/db/sqlc"
	"github.com/alim7007/go_bank_k8s/pb"
	"github.com/alim7007/go_bank_k8s/token"
	"github.com/alim7007/go_bank_k8s/util"
	"github.com/alim7007/go_bank_k8s/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedOlimBankServer
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
