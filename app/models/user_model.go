package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username" validate:"required,len=15"`
	Password []byte             `bson:"password" validate:"required"`
	Email    string             `bson:"email" validate:"required"`
	Role     string             `bson:"role" validate:"required"`
}

type UserResponse struct {
	ID       primitive.ObjectID `json:"id,omitempty" validate:"required"`
	Username string             `json:"username" validate:"required"`
	Email    string             `json:"email" validate:"required"`
	Role     string             `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,len=15"`
	Password string `json:"password" validate:"required"`
}

type SignUpRequest struct {
	Username string `json:"username" validate:"required,len=15"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
}
