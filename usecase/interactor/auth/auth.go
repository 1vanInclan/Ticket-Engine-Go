package auth

import (
	"context"
	userImpl "ticket-engine/interface/repository/user"
	"ticket-engine/usecase/dto"
	userRepo "ticket-engine/usecase/repository/user"
)

type AuthInteractor interface {
	Register(ctx context.Context, input dto.RegisterInput) (*dto.AuthResponse, error)
	Login(ctx context.Context, input dto.LoginInput) (*dto.AuthResponse, error)
}

type authInteractor struct {
	userRepository userRepo.UserRepository
}

var AuthInt AuthInteractor = &authInteractor{
	userRepository: userImpl.UserRepo,
}
