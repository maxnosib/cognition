package handlers

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/maxnosib/cognition/internal/config"
	"github.com/maxnosib/cognition/internal/entity"
	"github.com/maxnosib/cognition/internal/usecase"
	"github.com/maxnosib/cognition/utils"
)

type myHandlerFunc func(ctx context.Context, request *http.Request, us usecase.Model, userID int) utils.Answer
type myHandlerFuncAuth func(ctx context.Context, request *http.Request, us usecase.Model) (entity.User, error)

func New(ctx context.Context, cfg config.Server, us usecase.Model) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/note/{id}", writeResult(ctx, us, GetNote))
	r.Get("/notes", writeResult(ctx, us, GetNotes))
	r.Post("/note", writeResult(ctx, us, AddNote))
	r.Patch("/note", writeResult(ctx, us, UpdateNote))

	r.Get("/tags", writeResult(ctx, us, GetTags))
	r.Post("/tag", writeResult(ctx, us, AddTag))
	r.Delete("/tag/{id}", writeResult(ctx, us, DeleteTag))

	r.Get("/categories", writeResult(ctx, us, GetCategories))

	r.Post("/auth", writeResultAuth(ctx, us, Auth))
	r.Post("/register", writeResultAuth(ctx, us, Register))

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "./static"))
	FileServer(r, "/", filesDir)

	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatal("listen server fatal err: %w", err)
	}
}

// делаем обертку чтоб писать ответы централизованно.
func writeResult(ctx context.Context, us usecase.Model, next myHandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := utils.GetUserID(r)
		if err != nil {
			utils.WriteJSON(w, utils.Answer{Err: err})
			return
		}

		utils.WriteJSON(w, next(ctx, r, us, userID))
	})
}

// делаем обертку чтоб писать ответы централизованно.
func writeResultAuth(ctx context.Context, us usecase.Model, next myHandlerFuncAuth) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var answer utils.Answer
		usr, err := next(ctx, r, us)
		if err != nil {
			answer.Err = err
			utils.WriteJSON(w, answer)
			return
		}

		utils.AddCookie(w, "id", strconv.Itoa(usr.ID))
		answer.Data = usr
		utils.WriteJSON(w, answer)
	})
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
