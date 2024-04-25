package handlers

import (
	"bytes"
	"encoding/json"
	"myproject/internal/api/models"
	"myproject/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreatePerson(t *testing.T) {
	validPerson := models.Person{Name: "Alice", Age: 30, Hobbies: []string{"Reading", "Swimming"}}
	payload, err := json.Marshal(validPerson)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/person", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePerson)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var createdPerson models.Person
	err = json.Unmarshal(rr.Body.Bytes(), &createdPerson)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, validPerson.Name, createdPerson.Name)
	assert.Equal(t, validPerson.Age, createdPerson.Age)
	assert.Equal(t, validPerson.Hobbies, createdPerson.Hobbies)

	// Test invalid input
	invalidPerson := models.Person{Name: "", Age: 0, Hobbies: []string{""}}
	payload, err = json.Marshal(invalidPerson)
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("POST", "/person", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(CreatePerson)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetPerson(t *testing.T) {
	id := uuid.New().String()
	storage.AddPerson(models.Person{ID: id, Name: "Alice", Age: 30, Hobbies: []string{"Reading", "Swimming"}})

	req, err := http.NewRequest("GET", "/person/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var person models.Person
	err = json.Unmarshal(rr.Body.Bytes(), &person)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Alice", person.Name)

	// Test not found
	req, err = http.NewRequest("GET", "/person/"+uuid.New().String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}



func TestUpdatePerson(t *testing.T) {
	id := uuid.New().String()
	storage.AddPerson(models.Person{ID: id, Name: "Alice", Age: 30, Hobbies: []string{"Reading", "Swimming"}})

	updatedPerson := models.Person{Name: "Alice Updated", Age: 35, Hobbies: []string{"Reading", "Swimming", "Cycling"}}
	payload, err := json.Marshal(updatedPerson)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/person/"+id, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/person/{id}", UpdatePerson).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var person models.Person
	err = json.Unmarshal(rr.Body.Bytes(), &person)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, updatedPerson.Name, person.Name)
	assert.Equal(t, updatedPerson.Age, person.Age)
	assert.Equal(t, updatedPerson.Hobbies, person.Hobbies)

	// Test invalid input
	invalidPerson := models.Person{Name: "", Age: 0, Hobbies: []string{""}}
	payload, err = json.Marshal(invalidPerson)
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("PUT", "/person/"+id, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	// Test not found
	req, err = http.NewRequest("PUT", "/person/"+uuid.New().String(), bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeletePerson(t *testing.T) {
	id := uuid.New().String()
	storage.AddPerson(models.Person{ID: id, Name: "Alice", Age: 30, Hobbies: []string{"Reading", "Swimming"}})

	req, err := http.NewRequest("DELETE", "/person/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)

	// Test not found
	req, err = http.NewRequest("DELETE", "/person/"+uuid.New().String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestNotFoundHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NotFoundHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestIsValidHobbies(t *testing.T) {
	validHobbies := []string{"Reading", "Swimming"}
	assert.True(t, isValidHobbies(validHobbies))

	invalidHobbies := []string{"", "Swimming"}
	assert.False(t, isValidHobbies(invalidHobbies))

	invalidHobbies = nil
	assert.False(t, isValidHobbies(invalidHobbies))
}
