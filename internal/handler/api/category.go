package api

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	"MamangRust/echobloggrpc/internal/domain/response"
	"MamangRust/echobloggrpc/internal/pb"
	"MamangRust/echobloggrpc/pkg/auth"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

type categoryHandle struct {
	client pb.CategoryServiceClient
	token  auth.TokenManager
}

func NewHandlerCategory(client pb.CategoryServiceClient, router *echo.Echo, token auth.TokenManager) *categoryHandle {
	categoryHandler := &categoryHandle{
		client: client,
		token:  token,
	}

	routerCategory := router.Group("/api/category")

	routerCategory.GET("/hello", categoryHandler.handleHello)
	routerCategory.GET("/test-token", categoryHandler.handleTestToken)
	routerCategory.GET("/", categoryHandler.handleGetCategories)
	routerCategory.GET("/:id", categoryHandler.handleGetCategory)
	routerCategory.POST("/create", categoryHandler.handleCreateCategory)
	routerCategory.PUT("/update/:id", categoryHandler.handleUpdateCategory)
	routerCategory.DELETE("/delete/:id", categoryHandler.handleDeleteCategory)

	return categoryHandler
}

func (h *categoryHandle) handleHello(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return echo.ErrUnauthorized
	}

	tokenString := strings.Split(authHeader, " ")[1]

	claims, err := h.token.ValidateToken(tokenString)
	if err != nil {
		return err
	}

	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func (h *categoryHandle) handleTestToken(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return errors.New("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return errors.New("failed to cast claims as jwt.MapClaims")
	}
	return c.JSON(http.StatusOK, claims)
}

func (h *categoryHandle) handleGetCategories(c echo.Context) error {

	res, err := h.client.GetCategories(c.Request().Context(), &emptypb.Empty{})

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

func (h *categoryHandle) handleGetCategory(c echo.Context) error {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	res, err := h.client.GetCategory(c.Request().Context(), &pb.CategoryRequest{
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

func (h *categoryHandle) handleCreateCategory(c echo.Context) error {
	var body requests.CreateCategoryRequest

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

	data := &pb.CreateCategoryRequest{
		Name: body.Name,
	}

	res, err := h.client.CreateCategory(c.Request().Context(), data)

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

func (h *categoryHandle) handleUpdateCategory(c echo.Context) error {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	var body requests.UpdateCategoryRequest

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

	data := &pb.UpdateCategoryRequest{
		Id:   int32(body.ID),
		Name: body.Name,
	}

	res, err := h.client.UpdateCategory(c.Request().Context(), data)

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

func (h *categoryHandle) handleDeleteCategory(c echo.Context) error {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	res, err := h.client.DeleteCategory(c.Request().Context(), &pb.CategoryRequest{
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
