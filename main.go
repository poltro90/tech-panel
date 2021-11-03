package main

import (
	"database/sql"

	"net/http"

	"github.com/poltro90/tech-panel/controller"
	"github.com/poltro90/tech-panel/handler"
	"github.com/poltro90/tech-panel/store/mysql"
)

func main() {
	// Setup DB connection
	// Here we want to source environment variables and fetch the connection string
	db := sql.DB{}

	// With a DB connection, we can init the user store, responsible for retrieving data
	userStore := mysql.NewMySQLStore(db)

	// With a store, we can init the controller, implementing our business logic
	userController := controller.NewUserController(userStore)

	// With a controller, we can int the handler, taking care of serving service requests
	userHandler := handler.NewUserHandler(userController)

	// Eventually, we serve the endpoint and listed to the proper HTTP port
	http.HandleFunc("/company/:companyID/users", userHandler.GetUsers)
	http.ListenAndServe(":8090", nil)
}
