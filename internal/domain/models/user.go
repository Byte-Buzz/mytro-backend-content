package models

import "github.com/google/uuid"

type UserToken struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Sub      uuid.UUID `json:"sub"`
	Jti      uuid.UUID `json:"jti"`
	Exp      int64     `json:"exp"`
	Type     string    `json:"type"`
}
