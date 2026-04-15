package gateway

import (
	"context"
	"errors"
	"log/slog"

	"github.com/handziurdmytro/3JlArODA/auth-service/internal/crypto"
	"github.com/handziurdmytro/3JlArODA/auth-service/internal/repository"
	"github.com/handziurdmytro/3JlArODA/auth-service/pb"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	repo         *repository.Repository
	cryptoClient *crypto.Client
}

func NewServer(
	repo *repository.Repository,
	cryptoClient *crypto.Client,
) *Server {
	return &Server{
		repo:         repo,
		cryptoClient: cryptoClient,
	}
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	_, err := s.repo.GetUserByUsername(ctx, req.Username)
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "username already taken")
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, status.Error(codes.Internal, "failed to check username")
	}

	hash, err := s.cryptoClient.HashPassword(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to hash password")
	}

	user, err := s.repo.CreateUser(ctx, req.Username, hash)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	token, err := s.cryptoClient.SignJWT(user.ID.String(), user.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to sign JWT")
	}

	slog.Info("registered new user",
		slog.String("user_id", user.ID.String()),
		slog.String("username", user.Username),
	)

	return &pb.RegisterResponse{Token: token}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "database error")
	}

	isValidPassword, err := s.cryptoClient.VerifyPassword(req.Password, user.PasswordHash)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to verify password")
	}

	if !isValidPassword {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	token, err := s.cryptoClient.SignJWT(user.ID.String(), user.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate JWT")
	}

	slog.Info("user logged in successfully",
		slog.String("user_id", user.ID.String()),
		slog.String("username", user.Username),
	)

	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	resp, err := s.cryptoClient.ValidateJWT(req.Token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to validate token: %v", err)
	}

	return &pb.ValidateResponse{
		UserId:   resp.UserId,
		Username: resp.Username,
		IsValid:  resp.IsValid,
	}, nil
}
