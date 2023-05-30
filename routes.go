package main

func routes(r *Repo) {
	r.app.Get("/", r.Index)
	r.app.Post("/create", r.Create)
}
