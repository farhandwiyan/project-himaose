package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	InternalID int64 `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID uuid.UUID `json:"public_id" db:"public_id"` 
	Username string	`json:"username" db:"username"`
	Password string	`json:"password" db:"password"`
	Name string `json:"name" db:"name"`
	Role string	`json:"role" db:"role"`
	CreatedAt time.Time	`json:"created_at" db:"created_at"`
}

type UserResponse struct {
	PublicID   	uuid.UUID 		`json:"public_id"`
	Username string	`json:"username" db:"username"`
	Name 		string 			`json:"name"`
	Role 		string 			`json:"role" `
	CreatedAt 	time.Time 		`json:"created_at"`
}