package auth

import (
	"context"
	"encoding/json"
	"metrole/src/config"
	authpb "metrole/src/core/auth/gRPC"
	"os"
	"path/filepath"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type User struct {
	Token string `json:"token"`
}

type AuthService struct {
	user       *User
	configPath string
}

func NewAuthService() *AuthService {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	executablePath := filepath.Dir(ex)
	configPath := filepath.Join(executablePath, "auth_config.json")

	service := &AuthService{
		configPath: configPath,
	}

	service.user = service.LoadLocally()

	return service
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

	s.user = &User{Token: resp.Token}

	SaveLocallyErr := s.SaveLocally()
	if SaveLocallyErr != nil {
		return SaveLocallyErr
	}

	return nil
}

func (s *AuthService) IsAuthenticated() bool {

	if s.user == nil {
		return false
	}

	return s.user.Token != ""
}

func (s *AuthService) LoadLocally() *User {
	data, err := os.ReadFile(s.configPath)
	if err != nil {
		return nil
	}
	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil
	}
	return &user
}

func (s *AuthService) SaveLocally() error {
	data, err := json.Marshal(s.user)
	if err != nil {
		return err
	}

	return os.WriteFile(s.configPath, data, 0600)
}
