package handlers

import "github.com/uptrace/bun"

type Handlers struct {
	DB        *bun.DB
	JWTSecret string
}
