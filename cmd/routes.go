package main

func (a *app) routes() {
	a.router.GET("/ping", ping.Ping)

	a.router.POST("/users", users.Create)
	a.router.GET("/users/:user_id", users.Get)
	a.router.PUT("/users/:user_id", users.Update)
	a.router.PATCH("/users/:user_id", users.Update)
	a.router.DELETE("/users/:user_id", users.Delete)
	a.router.GET("/internal/users/search", users.Search)
	a.router.POST("/users/login", users.Login)
}
