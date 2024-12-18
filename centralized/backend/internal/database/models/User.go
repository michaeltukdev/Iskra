package models

import (
	"context"
	"iskra/centralized/internal/database"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int       `bun:",pk,autoincrement" json:"id,omitempty"`
	Email     string    `bun:",unique,notnull" json:"email"`
	Username  string    `bun:",unique" json:"username"`
	Password  string    `bun:",notnull" json:"password"`
	CreatedAt time.Time `bun:",default:current_timestamp" json:"created_at,omitempty"`
	UpdatedAt time.Time `bun:",default:current_timestamp" json:"updated_at,omitempty"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("Email is not valid"),
		),
		validation.Field(&u.Username,
			validation.Required.Error("Username is required"),
			validation.Length(3, 50).Error("Username must be between 3 and 50 characters"),
			is.Alphanumeric.Error("Username can only contain letters and numbers"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("Password is required"),
			validation.Length(8, 0).Error("Password must be at least 8 characters long"),
		),
	)
}

func (u User) ValidateLogin() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email,
			validation.Required.Error("Email is required"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("Password is required"),
			validation.Length(8, 0).Error("Password must be at least 8 characters long"),
		),
	)
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := database.DB.NewSelect().Model(&user).Where("email = ?", email).Scan(context.Background())

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := database.DB.NewSelect().Model(&user).Where("username = ?", username).Scan(context.Background())

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
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
