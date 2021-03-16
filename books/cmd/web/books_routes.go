package main

func (a *app) routes() {

	a.router.POST("/books", a.Add)
	a.router.GET("/books/:book_id", a.Get)
}
