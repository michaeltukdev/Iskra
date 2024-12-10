package models

import (
	"context"
	"iskra/centralized/internal/database"
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int       `bun:",pk,autoincrement" json:"id,omitempty"`
	Email     string    `bun:",unique,notnull" json:"email" validate:"required,email"`
	Username  string    `bun:",unique" json:"username" validate:"required,min=3,max=50,alphanum"`
	Password  string    `bun:",notnull" json:"password" validate:"required,min=8"`
	CreatedAt time.Time `bun:",default:current_timestamp" json:"created_at,omitempty"`
	UpdatedAt time.Time `bun:",default:current_timestamp" json:"updated_at,omitempty"`
}

func UserExists(email, username string) (bool, error) {
	var user User
	err := database.DB.NewSelect().Model(&user).Where("email = ? OR username = ?", email, username).Scan(context.Background())

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func CreateUser(user User) (*User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashed)

	_, err = database.DB.NewInsert().Model(&user).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return &user, nil
}
