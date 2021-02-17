package main

func (a *app) routes() {

	a.router.POST("/users", Create)
	a.router.GET("/users/:user_id", Get)
	a.router.PUT("/users/:user_id", Update)
	a.router.PATCH("/users/:user_id", Update)
	a.router.DELETE("/users/:user_id", Delete)
	a.router.GET("/internal/users/search", Search)
	a.router.POST("/users/login", Login)
}
