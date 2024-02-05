package api

import (
	"MamangRust/echobloggrpc/internal/pb"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type handler struct {
	auth     pb.AuthServiceClient
	category pb.CategoryServiceClient
	post     pb.PostServiceClient
	comment  pb.CommentServiceClient
	user     pb.UserServiceClient
}

func NewHandler(conn *grpc.ClientConn) *handler {
	clientAuth := pb.NewAuthServiceClient(conn)
	clientCategory := pb.NewCategoryServiceClient(conn)
	clientPost := pb.NewPostServiceClient(conn)
	clientComment := pb.NewCommentServiceClient(conn)
	clientUser := pb.NewUserServiceClient(conn)

	return &handler{
		auth:     clientAuth,
		category: clientCategory,
		post:     clientPost,
		comment:  clientComment,
		user:     clientUser,
	}
}

func (h *handler) Init(e *echo.Echo) {
	NewHandlerAuth(h.auth, e)
	NewHandlerCategory(h.category, e)
	NewHandlerPost(h.post, e)
	NewHandlerComment(h.comment, e)
	NewHandlerUser(h.user, e)

}
