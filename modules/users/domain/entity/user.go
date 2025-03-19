package entity

import (
	"time"
)

// User represents a user entity
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for User
func (*User) TableName() string {
	return "users"
}

// NewUser creates a new user
func NewUser(name, email, password string) *User {
	now := time.Now()
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
