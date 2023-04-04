package routes

import (
	"github.com/gorilla/mux"
	// "github.com/thomaslievre/go-bookstore/pkg/controllers"
	// "github.com/thomaslievre/go-bookstore/pkg/middlewares"
	"api/pkg/controllers"
	"api/pkg/middlewares"
	"net/http"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	//router.HandleFunc("/book/", controllers.GetBook).Methods("GET")

	router.Handle("/users/user/", middlewares.AuthMiddleware(http.HandlerFunc(controllers.CreateUser())).Methods("POST"))
	// router.Handle("/book/", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetBook))).Methods("GET")
	//router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	//router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	//router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	//router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
