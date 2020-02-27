package models

import "time"

// Banner describes structure of site banner
type Banner struct {
	ID          *int    `json:"ID,omitempty"`
	Header      *string `json:"header,omitempty"`
	Description *string `json:"description,omitempty"`
}

// User describes structure of site user
type User struct {
	ID       *int       `json:"ID,omitempty"`
	Phone    *int       `json:"phone,omitempty"`
	Password *string    `json:"password,omitempty"`
	Name     *string    `json:"name,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
	Email    *string    `json:"email,omitempty"`
	IsAdmin  *bool      `json:"isAdmin,omitempty"`
}
