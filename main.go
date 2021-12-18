package main

import (
	"net/http"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	// dependensiInjection
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)

	// manual
	// db := app.NewDB()
	// validate := validator.New()
	// categoryRepository := repository.NewCategoryRepository()
	// categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// categoryController := controller.NewCategoryController(categoryService)
	// router := app.NewRouter(categoryController)
	// authMiddleware := middleware.NewAuthMiddleware(router)

	// server := NewServer(authMiddleware)

	// err := server.ListenAndServe()
	// helper.PanicIfError(err)
}
