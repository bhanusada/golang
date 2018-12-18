package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/articles", func(r chi.Router) {
		r.With(paginate).Get("/", listArticles)
		r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate)

		// r.Post("/", createArticle)
		// r.Get("/search", searchArticles)

		r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug)

		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(ArticleCtx)
			r.Get("/", getArticle)
			// r.Put("/", updateArticle)
			// r.Delete("/", deleteArticle)
		})
	})

	r.Mount("/admin", adminRouter())

	http.ListenAndServe(":8080", r)
}

func ArticleCtx(next http.Handler) http.Handler {
	return http.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		articleID := chi.URLParam(r, "articleID")
		article, err := dbGetArticle(articleID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "article", article)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	article, ok := ctx.Value("article").(*Article)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(fmt.Sprintf("title: %s", article.Title)))
}

func adminRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(AdminOnly)
	r.Get("/", adminIndex)
	r.Get("/accounts", adminListAccounts)
	return r
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		perm, ok := ctx.Value("acl.permission").(YourPermissionType)
		if !ok || !perm.IsAdmin() {
			http.Error(w, httpStatusText(403), 403)
			return
		}
		next.ServeHTTP(w, r)
	})
}
