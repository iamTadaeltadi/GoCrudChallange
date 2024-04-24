package routes 
import (
	"myproject/internal/api/handlers"
	"net/http"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	// CORS
	corsOptions := gorillaHandlers.CORS(
		gorillaHandlers.AllowedHeaders([]string{"Content-Type"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		gorillaHandlers.AllowedOrigins([]string{"*"}),
	)

	r.HandleFunc("/person", handlers.CreatePerson).Methods("POST")
	r.HandleFunc("/person/{id}", handlers.GetPerson).Methods("GET")
	r.HandleFunc("/person/{id}", handlers.UpdatePerson).Methods("PUT")
	r.HandleFunc("/person/{id}", handlers.DeletePerson).Methods("DELETE")
	r.HandleFunc("/person", handlers.GetAllPersons).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)

	return corsOptions(r)
}
