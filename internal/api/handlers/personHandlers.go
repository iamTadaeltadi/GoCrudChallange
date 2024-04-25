package handlers
import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"myproject/internal/api/models"
	"myproject/internal/storage"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p models.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the input
	if p.Name == "" || p.Age <= 0 || !isValidHobbies(p.Hobbies) {
		http.Error(w, "name, age, and hobbies are required fields; hobbies should be an array of strings", http.StatusBadRequest)
		return
	}

	
	p.ID = uuid.New().String()

	storage.AddPerson(p)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	p, err := storage.GetPerson(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	persons := storage.GetAllPersons()

	if len(persons) == 0 {
		customMessage := "No persons found"
		http.Error(w, customMessage, http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var updatedPerson models.Person
	err := json.NewDecoder(r.Body).Decode(&updatedPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the input
	if updatedPerson.Name == "" || updatedPerson.Age <= 0 || !isValidHobbies(updatedPerson.Hobbies) {
		http.Error(w, "name, age, and hobbies are required fields; hobbies should be an array of strings or empty array", http.StatusBadRequest)
		return
	}

	updatedPerson.ID = id

	p, err := storage.UpdatePerson(updatedPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := storage.DeletePerson(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Not Found", http.StatusNotFound)
}

// isValidHobbies checks if the hobbies field is a valid array of strings
func isValidHobbies(hobbies []string) bool {
	if hobbies == nil {
		return false
	}
	for _, hobby := range hobbies {
		if hobby == "" {
			return false
		}
	}
	return true
}