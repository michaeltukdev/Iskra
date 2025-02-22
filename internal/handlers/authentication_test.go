package handlers

import (
	"bytes"
	"encoding/json"
	"iskra/centralized/internal/database"
	"iskra/centralized/internal/database/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestRegister(t *testing.T) {
	db, err := database.Init()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	h := &Handlers{
		DB:        db,
		JWTSecret: "test-secret",
	}

	e := echo.New()

	tests := []struct {
		name  string
		input models.User
		want  int
	}{
		{
			name:  "Data Submitted",
			input: models.User{},
			want:  http.StatusBadRequest,
		},
		{
			name: "Testing Validation",
			input: models.User{
				Email:    "asdasd",
				Username: "mi",
				Password: "asd",
			},
			want: http.StatusBadRequest,
		},
		{
			name: "Test Registration",
			input: models.User{
				Email:    "mikey.d.tilley@gmail.com",
				Username: "Michael",
				Password: "yourpassword123._",
			},
			want: http.StatusCreated,
		},
		{
			name: "Test Duplicate User Email",
			input: models.User{
				Email:    "mikey.d.tilley@gmail.com",
				Username: "Maozeo",
				Password: "yourpassword124._",
			},
			want: http.StatusConflict,
		},
		{
			name: "Test Duplicate Username",
			input: models.User{
				Email:    "new@gmail.com",
				Username: "Michael",
				Password: "yourpassword124._",
			},
			want: http.StatusConflict,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			payloadBytes, err := json.Marshal(test.input)
			if err != nil {
				t.Fatalf("Failed to marshal payload: %v", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(payloadBytes))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			err = h.Register(c)
			if err != nil {
				echoErr, ok := err.(*echo.HTTPError)
				if !ok {
					t.Fatalf("Handler returned an unexpected error type: %v", err)
				}

				if echoErr.Code != test.want {
					t.Errorf("Expected status code %d, got %d", test.want, echoErr.Code)
				}
			}
		})
	}
}
