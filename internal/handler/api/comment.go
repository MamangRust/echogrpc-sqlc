package api

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	"MamangRust/echobloggrpc/internal/domain/response"
	"MamangRust/echobloggrpc/internal/pb"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

type commentHandle struct {
	client pb.CommentServiceClient
}

func NewHandlerComment(client pb.CommentServiceClient, router *echo.Echo) *commentHandle {
	commentHandler := &commentHandle{
		client: client,
	}

	routerComment := router.Group("/api/comment")

	routerComment.GET("/hello", commentHandler.handleHello)
	routerComment.GET("/", commentHandler.handleGetComments)
	routerComment.GET("/:id", commentHandler.handleGetComment)
	routerComment.POST("/create", commentHandler.handleCreateComment)
	routerComment.PUT("/update/:id", commentHandler.handleUpdateComment)
	routerComment.DELETE("/delete/:id", commentHandler.handleDeleteComment)

	return commentHandler
}

func (h *commentHandle) handleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

func (h *commentHandle) handleGetComments(c echo.Context) error {
	res, err := h.client.GetComments(c.Request().Context(), &emptypb.Empty{})

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

func (h *commentHandle) handleGetComment(c echo.Context) error {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	res, err := h.client.GetComment(c.Request().Context(), &pb.CommentRequest{
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

func (h *commentHandle) handleCreateComment(c echo.Context) error {
	var body requests.CreateCommentRequest

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

	data := &pb.CreateCommentRequest{
		IdPostComment:   int32(body.IdPostComment),
		UserNameComment: body.Username,
		Comment:         body.Comment,
	}

	res, err := h.client.CreateComment(c.Request().Context(), data)

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

func (h *commentHandle) handleUpdateComment(c echo.Context) error {

	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	var body requests.UpdateCommentRequest

	body.Id = idInt

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

	data := &pb.UpdateCommentRequest{
		Id:      int32(body.Id),
		Comment: body.Comment,
	}

	res, err := h.client.UpdateComment(c.Request().Context(), data)

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

func (h *commentHandle) handleDeleteComment(c echo.Context) error {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	res, err := h.client.DeleteComment(c.Request().Context(), &pb.CommentRequest{
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
