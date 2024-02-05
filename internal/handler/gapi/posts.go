package gapi

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	"MamangRust/echobloggrpc/internal/pb"
	"MamangRust/echobloggrpc/internal/service"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type postHandleGrpc struct {
	pb.UnimplementedPostServiceServer
	post service.PostService
}

func NewPostHandleGrpc(post service.PostService) *postHandleGrpc {
	return &postHandleGrpc{post: post}
}

func (s *postHandleGrpc) GetPosts(ctx context.Context, req *empty.Empty) (*pb.GetPostsResponse, error) {
	posts, err := s.post.FindAll()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pbPosts []*pb.Post

	for _, post := range posts {
		pbPosts = append(pbPosts, &pb.Post{
			Id:     post.ID,
			Title:  post.Title,
			Slug:   post.Slug,
			Img:    post.Img,
			Body:   post.Body,
			UserId: post.UserID,
		})
	}

	return &pb.GetPostsResponse{Posts: pbPosts}, nil
}

func (s *postHandleGrpc) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	post, err := s.post.FindById(int(req.Id))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Post{
		Id:     post.ID,
		Title:  post.Title,
		Slug:   post.Slug,
		Img:    post.Img,
		Body:   post.Body,
		UserId: post.UserID,
	}, nil
}

func (s *postHandleGrpc) GetPostRelation(ctx context.Context, req *pb.GetPostRelationRequest) (*pb.GetPostRelationResponse, error) {
	post, err := s.post.FindByIDRelationJoin(int(req.Id))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetPostRelationResponse{
		PostId:          post.PostID,
		Title:           post.Title,
		CommentId:       post.CommentID,
		IdPostComment:   post.IDPostComment,
		UserNameComment: post.UserNameComment,
		Comment:         post.Comment,
	}, nil
}

func (s *postHandleGrpc) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	post, err := s.post.Create(&requests.CreatePostRequest{
		Title:      req.Title,
		Slug:       req.Slug,
		Img:        req.Img,
		Body:       req.Body,
		CategoryID: int(req.CategoryId),
		UserID:     int(req.UserId),
		UserName:   req.UserName,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Post{
		Id:     post.ID,
		Title:  post.Title,
		Slug:   post.Slug,
		Img:    post.Img,
		Body:   post.Body,
		UserId: post.UserID,
	}, nil
}

func (s *postHandleGrpc) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.Post, error) {
	post, err := s.post.Update(&requests.UpdatePostRequest{
		ID:         int(req.Id),
		Title:      req.Title,
		Slug:       req.Slug,
		Img:        req.Img,
		Body:       req.Body,
		CategoryID: int(req.CategoryId),
		UserID:     int(req.UserId),
		UserName:   req.UserName,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Post{
		Id:     post.ID,
		Title:  post.Title,
		Slug:   post.Slug,
		Img:    post.Img,
		Body:   post.Body,
		UserId: post.UserID,
	}, nil
}

func (s *postHandleGrpc) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	err := s.post.Delete(int(req.Id))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeletePostResponse{
		Success: true,
	}, nil
}
