package common

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//nolint:gochecknoglobals
var Router = chi.NewRouter()

//nolint:gochecknoinits
func init() {
	Router.Use(middleware.RequestID)
	Router.Use(middleware.RealIP)
	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)
	Router.Use(middleware.PathRewrite("/index.php", ""))
}

func Serve(port int) error {
	server := &http.Server{
		Handler:      Router,
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
