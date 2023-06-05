package common

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//nolint:gochecknoglobals
var router = chi.NewRouter()

//nolint:gochecknoinits
func init() {
	router.Use(middleware.Logger)
}

func Serve(port int) error {
	server := &http.Server{
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf(":%d", port),
	}

	fmt.Printf("Serving on *:%d\n", port)
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("unable to start http server: %w", err)
	}

	return nil
}
