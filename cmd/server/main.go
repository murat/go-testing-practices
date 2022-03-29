package main

import (
	"log"
	"net/http"
	"os"

	"github.com/murat/testing-practices/internal/user"
)

type Handler struct{}
type UserHandler struct {
	Repository *user.Repository
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.Handle("/ping", &Handler{})

	r, _ := user.New("users.db")
	mux.Handle("/users", &UserHandler{r})

	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("pong"))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	records, err := h.Repository.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// bytes, _ := json.Marshal(records)

	w.WriteHeader(http.StatusOK)
	w.Write(records.([]byte))
}
