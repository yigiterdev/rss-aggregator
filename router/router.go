package V1Router

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/yigiterdev/rss-aggregator/handlers"
	"github.com/yigiterdev/rss-aggregator/internal/database"
)

func GetRouter() *chi.Mux {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	apiCfg := handlers.ApiConfig{
		DB: dbQueries,
	}

	V1Router := chi.NewRouter()

	V1Router.Post("/users", apiCfg.HandlerCreateUser)
	V1Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUser))
	V1Router.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
	V1Router.Get("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeeds))
	V1Router.Post("/feed-follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	V1Router.Get("/feed-follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedFollows))
	V1Router.Delete("/feed-follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollow))

	V1Router.Get("/healthz", handlers.HandlerReadiness)
	V1Router.Get("/err", handlers.HandleErr)

	return V1Router
}
