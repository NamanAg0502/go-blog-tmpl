package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/namanag0502/go-blog/pkg/handlers"
	"github.com/namanag0502/go-blog/pkg/middleware"
)

func Routes() *chi.Mux {
	mux := chi.NewMux()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	h := middleware.AuthHandler{
		Username: os.Getenv("AUTH_USERNAME"),
		Password: os.Getenv("AUTH_PASSWORD"),
	}

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	mux.Get("/", handlers.Home)
	mux.Get("/view/{id}", handlers.ArticleView)
	mux.Get("/dashboard", h.Authenticate(handlers.Dashboard))
	mux.Get("/new", handlers.ArticleCreateForm)
	mux.Post("/create", h.Authenticate(handlers.ArticleCreate))
	mux.Get("/edit/{id}", h.Authenticate(handlers.ArticleUpdateForm))
	mux.Post("/update/{id}", h.Authenticate(handlers.ArticleUpdate))
	mux.Post("/delete/{id}", h.Authenticate(handlers.ArticleDelete))

	return mux
}
