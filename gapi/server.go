package gapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/habibiazmi123/simplebank/db/sqlc"
	"github.com/habibiazmi123/simplebank/pb"
	"github.com/habibiazmi123/simplebank/token"
	"github.com/habibiazmi123/simplebank/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}