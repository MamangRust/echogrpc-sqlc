package app

import (
	"MamangRust/echobloggrpc/internal/handler/gapi"
	"MamangRust/echobloggrpc/internal/pb"
	"MamangRust/echobloggrpc/internal/repository"
	"MamangRust/echobloggrpc/internal/service"
	"MamangRust/echobloggrpc/pkg/auth"
	"MamangRust/echobloggrpc/pkg/database/postgres"
	db "MamangRust/echobloggrpc/pkg/database/postgres/schema"
	"MamangRust/echobloggrpc/pkg/dotenv"
	"MamangRust/echobloggrpc/pkg/hash"
	"MamangRust/echobloggrpc/pkg/logger"
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

func RunServer() {
	logger, err := logger.NewLogger()

	if err != nil {
		logger.Fatal("Failed to create logger", zap.Error(err))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(err))
	}

	err = dotenv.Viper()

	if err != nil {
		logger.Fatal("Failed to load .env file", zap.Error(err))
	}

	token, err := auth.NewManager(viper.GetString("SECRET_KEY"))

	if err != nil {
		logger.Fatal("Failed to create token manager", zap.Error(err))
	}

	conn, err := postgres.NewClient(*logger)

	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	hash := hash.NewHashingPassword()

	DB := db.New(conn)

	repository := repository.NewRepositories(DB, context.Background())

	service, err := service.NewServices(service.Deps{
		Repository: repository,
		Hash:       hash,
		Logger:     logger,
		Token:      token,
	})

	if err != nil {
		logger.Fatal("Failed to create service", zap.Error(err))
	}

	handlerAuth := gapi.NewAuthHandlerGrpc(service.Auth)
	handlerUser := gapi.NewUserHandleGrpc(service.User)
	handlerCategory := gapi.NewCategoryHandleGrpc(service.Category)
	handlerPost := gapi.NewPostHandleGrpc(service.Post)
	handlerComment := gapi.NewCommentHandler(service.Comment)

	s := grpc.NewServer()

	pb.RegisterAuthServiceServer(s, handlerAuth)
	pb.RegisterUserServiceServer(s, handlerUser)
	pb.RegisterCategoryServiceServer(s, handlerCategory)
	pb.RegisterPostServiceServer(s, handlerPost)
	pb.RegisterCommentServiceServer(s, handlerComment)

	logger.Info(fmt.Sprintf("Server running on port %d", *port))

	if err := s.Serve(lis); err != nil {
		logger.Fatal("Failed to serve", zap.Error(err))
	}
}
