// package models

// import "gorm.io/gorm"

// type User struct {
// 	gorm.Model
// 	ID         uint   `gorm:"primaryKey"`
// 	Name       string `json:"name"`
// 	Email      string `json:"email" gorm:"unique"`
// 	Password   string `json:"-"`
// 	ProfilePic string `json:"profile_pic"`
// }
package models

import (
    "time"
    "github.com/google/uuid"
)

type User struct {
    ID        uuid.UUID `json:"id" db:"id"`
    Email     string    `json:"email" db:"email"`
    Username  string    `json:"username" db:"username"`
    Password  string    `json:"-" db:"password"`
    FirstName string    `json:"first_name" db:"first_name"`
    LastName  string    `json:"last_name" db:"last_name"`
    Role      UserRole  `json:"role" db:"role"`
    Active    bool      `json:"active" db:"active"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type UserRole string

const (
    RoleAdmin    UserRole = "admin"
    RoleUser     UserRole = "user"
    RoleGuest    UserRole = "guest"
)

// Value objects
type Email struct {
    Value string
}

func (e Email) Validate() error {
    // Email validation logic
    return nil
}

type Password struct {
    Hash string
}

func (p Password) Verify(plain string) bool {
    // Password verification logic
    return true
}