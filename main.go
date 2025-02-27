package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Deadline  time.Time `json:"deadline,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

var db *sql.DB

func main() {
	// Database connection string
	connStr := "postgres://renstask:your_password@localhost/renstask?sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")

	// Create tasks table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			completed BOOLEAN DEFAULT FALSE,
			deadline TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Set up router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/tasks", getTasks).Methods("GET")
	r.HandleFunc("/tasks", createTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	// Start server
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Handler functions will go here
func getTasks(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Get tasks endpoint"}`))
}

func createTask(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Create task endpoint"}`))
}

func getTask(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Get single task endpoint"}`))
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Update task endpoint"}`))
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Delete task endpoint"}`))
}
