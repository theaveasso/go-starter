package application

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/theaveasso/go-starter/internal/handler"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func setupRoutes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	return router
}

func make(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiError, ok := err.(handler.APIError); ok {
				handler.WriteJSON(
					w,
					apiError.StatusCode,
					handler.Envelope{"error": apiError},
					nil,
				)
        return
			}

			handler.WriteJSON(
				w,
				http.StatusInternalServerError,
				handler.Envelope{
          "error":
					"server encountered a problem and could not process your request",
				},
        nil,
			)
			slog.Error("http handler", "error", err.Error(), "path", r.URL.Path)
      return
		}
	}
}
