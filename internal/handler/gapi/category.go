package gapi

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	"MamangRust/echobloggrpc/internal/pb"
	"MamangRust/echobloggrpc/internal/service"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type categoryHandleGrpc struct {
	pb.UnimplementedCategoryServiceServer
	category service.CategoryService
}

func NewCategoryHandleGrpc(category service.CategoryService) *categoryHandleGrpc {
	return &categoryHandleGrpc{category: category}
}

func (s *categoryHandleGrpc) GetCategories(ctx context.Context, empty *emptypb.Empty) (*pb.CategoriesResponse, error) {
	categories, err := s.category.FindAll()
	if err != nil {
		return nil, err
	}

	var pbCategories []*pb.Category

	for _, category := range categories {
		createdAtProto := timestamppb.New(category.CreatedAt.Time)

		var updatedAtProto *timestamppb.Timestamp
		if category.UpdatedAt.Valid {
			updatedAtProto = timestamppb.New(category.UpdatedAt.Time)
		}

		pbCategory := &pb.Category{
			Id:        int32(category.ID),
			Name:      category.Name,
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
		}
		pbCategories = append(pbCategories, pbCategory)
	}

	return &pb.CategoriesResponse{Categories: pbCategories}, nil
}

func (s *categoryHandleGrpc) GetCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.CategoryResponse, error) {
	category, err := s.category.FindByID(int(req.Id))
	if err != nil {
		return nil, err
	}
	createdAtProto := timestamppb.New(category.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp
	if category.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(category.UpdatedAt.Time)
	}

	return &pb.CategoryResponse{
		Category: &pb.Category{
			Id:        int32(category.ID),
			Name:      category.Name,
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
		},
	}, nil

}

func (s *categoryHandleGrpc) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	newCategory := &requests.CreateCategoryRequest{
		Name: req.Name,
	}

	category, err := s.category.Create(newCategory)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	createdAtProto := timestamppb.New(category.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp

	if category.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(category.UpdatedAt.Time)
	}

	return &pb.CategoryResponse{
		Category: &pb.Category{
			Id:        int32(category.ID),
			Name:      category.Name,
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
		},
	}, nil
}

func (s *categoryHandleGrpc) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := s.category.Update(&requests.UpdateCategoryRequest{
		ID:   int(req.Id),
		Name: req.Name,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	createdAtProto := timestamppb.New(category.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp

	if category.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(category.UpdatedAt.Time)
	}

	return &pb.CategoryResponse{
		Category: &pb.Category{
			Id:        int32(category.ID),
			Name:      category.Name,
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
		},
	}, nil
}

func (s *categoryHandleGrpc) DeleteCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.DeleteCategoryResponse, error) {
	err := s.category.Delete(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.DeleteCategoryResponse{
		Success: true,
	}, nil
}
