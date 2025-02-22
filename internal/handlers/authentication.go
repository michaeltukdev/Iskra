package handlers

import (
	"fmt"
	"iskra/centralized/internal/database/models"
	"iskra/centralized/internal/middlewares"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handlers) Register(c echo.Context) error {
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

	existingEmailUser, err := models.GetUserByEmail(user.Email, h.DB)
	if err != nil {
		fmt.Printf("Failed to check if user exists by email: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	if existingEmailUser != nil {
		return echo.NewHTTPError(http.StatusConflict, "Email already in use")
	}

	existingUsernameUser, err := models.GetUserByUsername(user.Username, h.DB)
	if err != nil {
		fmt.Printf("Failed to check if user exists by username: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	if existingUsernameUser != nil {
		return echo.NewHTTPError(http.StatusConflict, "Username already in use")
	}

	_, err = models.CreateUser(user, h.DB)
	if err != nil {
		fmt.Printf("Failed to create user: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "User created successfully")
}

func (h *Handlers) Login(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		fmt.Printf("Failed to bind user: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := user.ValidateLogin()
	if err != nil {
		fmt.Printf("Failed to validate user: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	validUser, err := models.GetUserByEmail(user.Email, h.DB)
	if err != nil {
		fmt.Printf("Failed to get user: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if validUser == nil {
		fmt.Println("User not found")
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(validUser.Password), []byte(user.Password))
	if err != nil {
		fmt.Printf("Failed to compare passwords: %v\n", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	claims := &middlewares.JWTCustomClaims{
		UserID:   validUser.ID,
		Username: validUser.Username,
		Email:    validUser.Email,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(h.JWTSecret))
	if err != nil {
		fmt.Printf("Failed to sign token: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.HttpOnly = true
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "Logged in successfully")
}

func (h *Handlers) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.MaxAge = 0

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "Logged out successfully")
}

func (h *Handlers) Me(c echo.Context) error {
	if user := c.Get("user"); user != nil {
		return c.JSON(http.StatusOK, user)
	}

	return c.JSON(http.StatusOK, "No user in context")
}
