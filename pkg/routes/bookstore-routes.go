package routes

import (
	"github.com/ascendantaditya/Go-BookManagement/pkg/controllers"
	"github.com/gorilla/mux"
)

//absolute path

var RegisterBookstoreRoutes = func(router *mux.Router) {
	//bookstore routes
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}
