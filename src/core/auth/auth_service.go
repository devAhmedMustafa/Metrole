package auth

import (
	"context"
	"metrole/src/config"
	authpb "metrole/src/core/auth/gRPC"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthService struct {
	token string
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(id_token string) error {

	conn, err := grpc.NewClient(config.AppSettings.GRPC_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	defer conn.Close()
	client := authpb.NewAuthGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	resp, err := client.GoogleLogin(ctx, &authpb.GoogleLoginRequest{IdToken: id_token})
	if err != nil {
		return err
	}

	s.token = resp.Token

	return nil
}
