package main

func routes(r *Repo) {
	r.app.Get("/", r.Index)
	r.app.Post("/create", r.Create)
	users := r.app.Group("/users")
	users.Post("/sign-up", r.CreateUser)
	users.Post("/sign-in", r.LoginUser)
}
