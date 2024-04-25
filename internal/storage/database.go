package storage

import (
	"fmt"
	"myproject/internal/api/models"
	"sync"

	"github.com/google/uuid"
)

type DB struct {
	m    map[string]models.Person
	lock sync.RWMutex
}

var db = NewDB()

func NewDB() *DB {
	return &DB{
		m: make(map[string]models.Person),
	}
}

func GenerateID() string {
	return uuid.New().String()
}

func AddPerson(p models.Person) {
	db.lock.Lock()
	defer db.lock.Unlock()

	db.m[p.ID] = p
}

func GetPerson(id string) (models.Person, error) {
	db.lock.RLock()
	defer db.lock.RUnlock()

	p, ok := db.m[id]
	if !ok {
		return models.Person{}, fmt.Errorf("PERSON NOT FOUND")
	}

	return p, nil
}

func GetAllPersons() []models.Person {
	db.lock.RLock()
	defer db.lock.RUnlock()

	var persons []models.Person
	for _, p := range db.m {
		persons = append(persons, p)
	}

	return persons
}

func UpdatePerson(updatedPerson models.Person) (models.Person, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	p, ok := db.m[updatedPerson.ID]
	if !ok {
		return models.Person{}, fmt.Errorf("PERSON NOT FOUND")
	}

	p.Name = updatedPerson.Name
	p.Age = updatedPerson.Age
	p.Hobbies = updatedPerson.Hobbies

	db.m[updatedPerson.ID] = p

	return p, nil
}

func DeletePerson(id string) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	_, ok := db.m[id]
	if !ok {
		return fmt.Errorf("PERSON NOT FOUND")
	}

	delete(db.m, id)
	return nil
}
