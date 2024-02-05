package api

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	"MamangRust/echobloggrpc/internal/domain/response"
	"MamangRust/echobloggrpc/internal/pb"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userHandle struct {
	pb.UnimplementedUserServiceServer
	user pb.UserServiceClient
}

func NewHandlerUser(client pb.UserServiceClient, router *echo.Echo) *userHandle {
	userHandler := &userHandle{
		user: client,
	}

	routerUser := router.Group("/api/user")

	routerUser.GET("/hello", userHandler.handleHello)

	routerUser.GET("/", userHandler.handleGetUsers)
	routerUser.GET("/:id", userHandler.handleGetUser)
	routerUser.POST("/create", userHandler.handleCreateUser)
	routerUser.PUT("/update/:id", userHandler.handleUpdateUser)
	routerUser.DELETE("/delete/:id", userHandler.handleDeleteUser)

	return userHandler

}

func (h *userHandle) handleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

func (h *userHandle) handleGetUsers(c echo.Context) error {
	res, err := h.user.GetUsers(c.Request().Context(), &emptypb.Empty{})

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *userHandle) handleGetUser(c echo.Context) error {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	res, err := h.user.GetUser(c.Request().Context(), &pb.UserRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *userHandle) handleCreateUser(c echo.Context) error {

	var body requests.CreateUserRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request Validate: " + err.Error(),
			Data:       nil,
		})
	}

	data := &pb.CreateUserRequest{
		Firstname:       body.FirstName,
		Lastname:        body.LastName,
		Email:           body.Email,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
	}

	res, err := h.user.CreateUser(c.Request().Context(), data)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *userHandle) handleUpdateUser(c echo.Context) error {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	var body requests.UpdateUserRequest

	body.ID = idInt

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request Validate: " + err.Error(),
			Data:       nil,
		})
	}

	data := &pb.UpdateUserRequest{
		Id:              int32(body.ID),
		Firstname:       body.FirstName,
		Lastname:        body.LastName,
		Email:           body.Email,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
	}

	res, err := h.user.UpdateUser(c.Request().Context(), data)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *userHandle) handleDeleteUser(c echo.Context) error {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	res, err := h.user.DeleteUser(c.Request().Context(), &pb.UserRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}
