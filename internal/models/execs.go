package models

import "database/sql"

type Exec struct {
	ID                   int            `json:"id,omitempty" db:"id,omitempty"`
	FirstName            string         `json:"first_name,omitempty" db:"first_name,omitempty"`
	LastName             string         `json:"last_name,omitempty" db:"last_name,omitempty"`
	Email                string         `json:"email,omitempty" db:"email,omitempty"`
	Username             string         `json:"username,omitempty" db:"username,omitempty"`
	Password             string         `json:"password,omitempty" db:"password,omitempty"`
	PasswordChangedAt    sql.NullString `json:"password_changed_at,omitempty" db:"password_changed_at,omitempty"`
	UserCreatedAt        sql.NullString `json:"user_created_at,omitempty" db:"user_created_at,omitempty"`
	PasswordResetToken   sql.NullString `json:"password_reset_token,omitempty" db:"password_reset_token,omitempty"`
	PasswordTokenExpires sql.NullString `json:"password_token_expires,omitempty" db:"password_token_expires,omitempty"`
	InactiveStatus       bool           `json:"inactive_status,omitempty" db:"inactive_status,omitempty"`
	Role                 string         `json:"role,omitempty" db:"role,omitempty"`
}
