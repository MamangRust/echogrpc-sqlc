package api

import (
	"MamangRust/echobloggrpc/internal/pb"
	"MamangRust/echobloggrpc/pkg/auth"
	"fmt"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

// type handler struct {
// 	auth     pb.AuthServiceClient
// 	category pb.CategoryServiceClient
// 	post     pb.PostServiceClient
// 	comment  pb.CommentServiceClient
// 	user     pb.UserServiceClient
// 	token    auth.TokenManager
// }

func NewHandler(conn *grpc.ClientConn, token auth.TokenManager, e *echo.Echo) {
	if token == nil {
		fmt.Println("Token is not initialized")
	}

	clientAuth := pb.NewAuthServiceClient(conn)
	clientCategory := pb.NewCategoryServiceClient(conn)
	clientPost := pb.NewPostServiceClient(conn)
	clientComment := pb.NewCommentServiceClient(conn)
	clientUser := pb.NewUserServiceClient(conn)

	NewHandlerAuth(clientAuth, e)
	NewHandlerCategory(clientCategory, e, token)
	NewHandlerPost(clientPost, e)
	NewHandlerComment(clientComment, e)
	NewHandlerUser(clientUser, e)

}
