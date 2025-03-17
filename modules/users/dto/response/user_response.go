// internal/modules/user/interfaces/dto/response/user_response.go

package response

import (
	"go-modular-boilerplate/modules/users/domain/entity"
	"time"
)

// UserResponse represents a user response
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FromEntity converts a user entity to a user response
func FromEntity(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// FromEntities converts a slice of user entities to a slice of user responses
func FromEntities(users []*entity.User) []*UserResponse {
	userResponses := make([]*UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = FromEntity(user)
	}
	return userResponses
}
