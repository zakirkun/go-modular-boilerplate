package service

import (
	"context"
	"errors"
	"go-modular-boilerplate/modules/users/domain/entity"
	"go-modular-boilerplate/modules/users/domain/repository"
)

// Errors
var (
	ErrUserNotFound     = errors.New("user not found")
	ErrEmailAlreadyUsed = errors.New("email already in use")
)

// UserService handles user domain logic
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// GetAllUsers gets all users
func (s *UserService) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	return s.userRepo.FindAll(ctx)
}

// GetUserByID gets a user by ID
func (s *UserService) GetUserByID(ctx context.Context, id uint) (*entity.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	// existingUser, err := s.userRepo.FindByEmail(ctx, user.Email)
	// if err != nil && err != repository.ERR_RECORD_NOT_FOUND {
	// 	return err
	// }
	// if existingUser != nil {
	// 	return ErrEmailAlreadyUsed
	// }

	return s.userRepo.Create(ctx, user)
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	existingUser, err := s.userRepo.FindByID(ctx, user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return ErrUserNotFound
	}

	return s.userRepo.Update(ctx, user)
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	existingUser, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return ErrUserNotFound
	}

	return s.userRepo.Delete(ctx, id)
}
