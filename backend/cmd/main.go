package main

import (
	"log"
	"net/http"

	deliveryHTTP "golangbackend/delivery/http"
	"golangbackend/delivery/middleware"
	"golangbackend/infrastructure"
	"golangbackend/repository"
	"golangbackend/usecase"

	"github.com/gorilla/mux"
)

func main() {
	db, err := infrastructure.NewSQLiteDB("./users.db")
	if err != nil {
		log.Fatal("cannot open database:", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUseCase(userRepo)
	userHandler := deliveryHTTP.NewUserHandler(userUC)

	r := mux.NewRouter()

	r.Use(corsMiddleware)

	r.HandleFunc("/api/register", userHandler.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/login", userHandler.Login).Methods("POST", "OPTIONS")

	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTMiddleware)
	protected.HandleFunc("/profile", userHandler.Profile).Methods("GET", "OPTIONS")

	log.Println("Backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
