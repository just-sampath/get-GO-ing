package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Todo struct {
	Title string `json:"name"`
	Body  string `json:"body"`
}

const authKey string = "test-key"

func createDBConnection() *pgxpool.Pool {
	dbUrl := "postgres://postgres:pass@localhost:5432/postgres"

	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Printf("Failed to connect to DB with an error")
		os.Exit(1)
	}

	return pool
}

func requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var providedApiKey string = r.Header.Get("Authorization")
		if providedApiKey != authKey {
			w.WriteHeader(401)
			w.Write([]byte("Wrong auth key provided!"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

type TodoService struct {
	pool *pgxpool.Pool
}

func (todo *TodoService) createTodo(w http.ResponseWriter, r *http.Request) {
	var req Todo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", 400)
		return
	}

	var newId string
	err := todo.pool.QueryRow(context.Background(), "INSERT INTO TODOS_TEST VALUES($1, $2) RETURNING title", req.Title, req.Body).Scan(&newId)
	if err != nil {
		fmt.Printf("Failed to insert new todo: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Failed to create todo"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Todo Created with " + fmt.Sprint(newId) + "!"))
}

func main() {
	var pool *pgxpool.Pool = createDBConnection()
	defer pool.Close()

	var service TodoService = TodoService{
		pool: pool,
	}

	r := chi.NewRouter()

	r.Use(requireAuth)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go!"))
	})
	r.Post("/", service.createTodo)

	fmt.Println("Server running on Port 8080")
	http.ListenAndServe(":8080", r)
}
