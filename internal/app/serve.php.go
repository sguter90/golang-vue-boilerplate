package app

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func (app *App) Serve(params []string) error {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(app.appCtx)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", ping)
	r.Get("/ping", ping)

	port := strconv.Itoa(app.Config.ServePort)
	err := http.ListenAndServe(":"+port, r)

	return err
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (app *App) appCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", app.Db)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
