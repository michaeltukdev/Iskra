package handlers

import (
	"fmt"
	"iskra/centralized/internal/database/models"
	"iskra/centralized/internal/middlewares"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		fmt.Printf("Failed to bind user: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := user.Validate()
	if err != nil {
		fmt.Printf("Failed to validate user: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	existingEmailUser, err := models.GetUserByEmail(user.Email)
	if err != nil {
		fmt.Printf("Failed to check if user exists by email: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	if existingEmailUser != nil {
		return echo.NewHTTPError(http.StatusConflict, "Email already in use")
	}

	existingUsernameUser, err := models.GetUserByUsername(user.Username)
	if err != nil {
		fmt.Printf("Failed to check if user exists by username: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	if existingUsernameUser != nil {
		return echo.NewHTTPError(http.StatusConflict, "Username already in use")
	}

	_, err = models.CreateUser(user)
	if err != nil {
		fmt.Printf("Failed to create user: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "User created successfully")
}

func Login(c echo.Context) error {
	var user *models.User

	if err := c.Bind(&user); err != nil {
		fmt.Printf("Failed to bind user: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	validator := validator.New()
	if err := validator.StructPartial(user, "Email", "Password"); err != nil {
		fmt.Printf("Failed to validate user: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	validUser, err := models.GetUserByEmail(user.Email)
	if err != nil {
		fmt.Printf("Failed to get user: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(validUser.Password), []byte(user.Password))
	if err != nil {
		fmt.Printf("Failed to compare passwords: %v\n", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	claims := &middlewares.JWTCustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	if err != nil {
		fmt.Printf("Failed to sign token: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JWTCustomClaims)
	name := claims.Username

	return c.String(http.StatusOK, "Welcome "+name+"!")
}
