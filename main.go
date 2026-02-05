package main

import (
	"context"
	"log"
	"net"

	// Import du code généré (grâce au replace dans go.mod)
	pb "github.com/TheJobMateCompany/jobmate-proto/gen/go/proto/auth/v1"
	"google.golang.org/grpc"
)

// server implémente l'interface générée par Protobuf
type server struct {
	pb.UnimplementedAuthServiceServer
}

// Implémentation dummy de Login pour tester
func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	log.Printf("Tentative de login pour : %v", req.Email)
	return &pb.AuthResponse{
		AccessToken: "dummy_token_123",
		UserId:      "user_0000",
	}, nil
}

// Implémentation dummy de Register
func (s *server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	return &pb.AuthResponse{AccessToken: "new_token"}, nil
}

// Implémentation dummy de ValidateToken
func (s *server) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	return &pb.ValidateTokenResponse{IsValid: true, UserId: "user_0000"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	
	// Enregistrement du service
	pb.RegisterAuthServiceServer(s, &server{})

	log.Printf("Auth Service listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}