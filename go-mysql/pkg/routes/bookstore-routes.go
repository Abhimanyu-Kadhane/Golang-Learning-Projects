package routes

import (
	"net/http"

	"github.com/Abhimanyu-Kadhane/go-mysql/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/book/", controllers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods(http.MethodDelete)
}
