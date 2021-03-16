package main

func (a *app) routes() {

	a.router.POST("/users", a.Create)
	a.router.GET("/users/:user_id", a.Get)
	a.router.PUT("/users/:user_id", a.Update)
	a.router.PATCH("/users/:user_id", a.Update)
	a.router.DELETE("/users/:user_id", a.Delete)
	a.router.GET("/internal/users/search", a.Search)
	a.router.POST("/users/login", a.Login)
}
