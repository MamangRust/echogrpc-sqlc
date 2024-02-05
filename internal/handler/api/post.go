package api

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	"MamangRust/echobloggrpc/internal/domain/response"
	"MamangRust/echobloggrpc/internal/pb"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

type postHandle struct {
	client pb.PostServiceClient
}

func NewHandlerPost(client pb.PostServiceClient, router *echo.Echo) *postHandle {
	postHandler := &postHandle{
		client: client,
	}

	routerPost := router.Group("/api/post")

	routerPost.GET("/hello", postHandler.handleHello)
	routerPost.GET("/", postHandler.handleGetPosts)
	routerPost.GET("/:id", postHandler.handleGetPost)
	routerPost.POST("/create", postHandler.handleCreatePost)
	routerPost.PUT("/update/:id", postHandler.handleUpdatePost)
	routerPost.DELETE("/delete/:id", postHandler.handleDeletePost)

	return postHandler
}

func (h *postHandle) handleHello(c echo.Context) error {
	return c.JSON(200, "Hello World")
}

func (h *postHandle) handleGetPosts(c echo.Context) error {
	empty := &emptypb.Empty{}

	res, err := h.client.GetPosts(c.Request().Context(), empty)

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

func (h *postHandle) handleGetPost(c echo.Context) error {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	res, err := h.client.GetPost(c.Request().Context(), &pb.GetPostRequest{
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

func (h *postHandle) GetPostRelations(c echo.Context) error {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	res, err := h.client.GetPostRelation(c.Request().Context(), &pb.GetPostRelationRequest{
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

func (h *postHandle) handleCreatePost(c echo.Context) error {

	var body requests.CreatePostRequest

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

	data := &pb.CreatePostRequest{
		Title:      body.Title,
		Slug:       body.Slug,
		Img:        body.Img,
		Body:       body.Body,
		UserId:     int32(body.UserID),
		CategoryId: int32(body.CategoryID),
		UserName:   body.UserName,
	}

	res, err := h.client.CreatePost(c.Request().Context(), data)

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

func (h *postHandle) handleUpdatePost(c echo.Context) error {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	var body requests.UpdatePostRequest

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

	data := &pb.UpdatePostRequest{
		Id:         int32(body.ID),
		Title:      body.Title,
		Slug:       body.Slug,
		Img:        body.Img,
		Body:       body.Body,
		UserId:     int32(body.UserID),
		CategoryId: int32(body.CategoryID),
		UserName:   body.UserName,
	}

	res, err := h.client.UpdatePost(c.Request().Context(), data)

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

func (h *postHandle) handleDeletePost(c echo.Context) error {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	res, err := h.client.DeletePost(c.Request().Context(), &pb.DeletePostRequest{
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
