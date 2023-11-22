package gapi

import (
	"fmt"

	"github.com/RakshitKumar04/ChatMate/tree/master/go-server/db"
	"github.com/RakshitKumar04/ChatMate/tree/master/go-server/pb"
	"github.com/RakshitKumar04/ChatMate/tree/master/go-server/token"
	"github.com/RakshitKumar04/ChatMate/tree/master/go-server/utils"
)

type Server struct {
	pb.GrpcServerServiceServer
	config       utils.ViperConfig
	dbCollection db.MongoCollections

	tokenMaker token.Maker
}

func NewServer(config utils.ViperConfig, dbCollection db.MongoCollections) (*Server, error) {
	tokenMaker, err := token.NewJwtMaker(config.TokkenStructureKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:       config,
		dbCollection: dbCollection,
		tokenMaker:   tokenMaker,
	}

	return server, nil
}
