package auth

import (
	"context"
	"errors"
	"ticket-engine/domain/model"
	"ticket-engine/infrastructure/security"
	"ticket-engine/usecase/dto"
)

func (i *authInteractor) Register(ctx context.Context, input dto.RegisterInput) (*dto.AuthResponse, error) {

	// 1. Verificar si el usuario ya existe
	existingUser, err := i.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email is already registered")
	}

	// 2. Hashear la contraseña
	hashedPassword, err := security.HashPassword(input.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// 3. Crear el modelo de dominio
	newUser := &model.User{
		Email:    input.Email,
		Password: hashedPassword,
	}

	// 4. Guardar en BD
	if err := i.userRepository.Create(ctx, newUser); err != nil {
		return nil, err
	}

	// 5. Generar token JWT
	token, err := security.GenerateToken(newUser.ID, newUser.Email, "")
	if err != nil {
		return nil, errors.New("failed to generate auth token")
	}

	return &dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:    newUser.ID,
			Email: newUser.Email,
		},
	}, nil
}

func (i *authInteractor) Login(ctx context.Context, input dto.LoginInput) (*dto.AuthResponse, error) {
	// 1. Buscar usuario por email
	user, err := i.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	// 2. Validar contraseña
	if !security.CheckPasswordHash(input.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	// 3. Generar token JWT
	token, err := security.GenerateToken(user.ID, user.Email, "")
	if err != nil {
		return nil, errors.New("failed to generate auth token")
	}

	return &dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:    user.ID,
			Email: user.Email,
		},
	}, nil
}
