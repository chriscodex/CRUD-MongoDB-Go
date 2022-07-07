package models

import "time"

// User: data of user
type User struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Users: List of users
type Users []User
