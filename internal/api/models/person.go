package models

// Person represents a person in our database
type Person struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}
