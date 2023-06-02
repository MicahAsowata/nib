package main

import (
	"github.com/MicahAsowata/nib/config"
	jwtware "github.com/gofiber/contrib/jwt"
)

func routes(r *Repo) {
	r.app.Get("/", r.Index)
	r.app.Post("/create", r.Create)
	users := r.app.Group("/users")
	users.Post("/sign-up", r.CreateUser)
	users.Post("/sign-in", r.LoginUser)
	r.app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.Secret)},
	}))
	r.app.Get("/protected", r.Protected)
}
