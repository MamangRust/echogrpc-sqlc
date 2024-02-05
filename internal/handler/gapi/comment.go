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

type commentHandleGrpc struct {
	pb.UnimplementedCommentServiceServer
	comment service.CommentService
}

func NewCommentHandler(comment service.CommentService) *commentHandleGrpc {
	return &commentHandleGrpc{
		comment: comment,
	}
}

func (h *commentHandleGrpc) GetComments(ctx context.Context, req *emptypb.Empty) (*pb.CommentsResponse, error) {
	res, err := h.comment.FindAll()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pbComments []*pb.Comment

	for _, comment := range res {
		pbComments = append(pbComments, &pb.Comment{
			Id:              int32(comment.ID),
			Comment:         comment.Comment,
			IdPostComment:   int32(comment.IDPostComment),
			UserNameComment: comment.UserNameComment,
			CreatedAt:       timestamppb.New(comment.CreatedAt.Time),
		})
	}

	return &pb.CommentsResponse{
		Comments: pbComments,
	}, nil
}

func (h *commentHandleGrpc) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error) {
	res, err := h.comment.Create(&requests.CreateCommentRequest{
		Comment:       req.Comment,
		IdPostComment: int(req.IdPostComment),
		Username:      req.UserNameComment,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CommentResponse{
		Comment: &pb.Comment{
			Id:              int32(res.ID),
			Comment:         res.Comment,
			IdPostComment:   int32(res.IDPostComment),
			UserNameComment: res.UserNameComment,
			CreatedAt:       timestamppb.New(res.CreatedAt.Time),
		},
	}, nil

}

func (h *commentHandleGrpc) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.CommentResponse, error) {
	res, err := h.comment.Update(&requests.UpdateCommentRequest{
		Comment:       req.Comment,
		IdPostComment: int(req.IdPostComment),
		Username:      req.UserNameComment,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CommentResponse{
		Comment: &pb.Comment{
			Id:              int32(res.ID),
			Comment:         res.Comment,
			IdPostComment:   int32(res.IDPostComment),
			UserNameComment: res.UserNameComment,
			CreatedAt:       timestamppb.New(res.CreatedAt.Time),
		},
	}, nil
}

func (h *commentHandleGrpc) DeleteComment(ctx context.Context, req *pb.CommentRequest) (*pb.DeleteCommentResponse, error) {
	err := h.comment.Delete(int(req.Id))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteCommentResponse{
		Success: true,
	}, nil
}
