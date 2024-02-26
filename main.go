package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/skrb7f16/rssagg/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("helo")
	godotenv.Load()
	portString := os.Getenv("PORT")
	dburl := os.Getenv("DB_URL")
	if portString == "" {
		log.Fatal("Port not founds")
	}
	if dburl == "" {
		log.Fatal("DB url not found")
	}

	con, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal("Error")
	}

	apiCfg := apiConfig{
		DB: database.New(con),
	}
	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)
	v1Router.Get("/health", handlerCheck)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/user", apiCfg.handleCreateUser)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handleGetUser))
	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))
	v1Router.Get("/feed", apiCfg.handleGetFeeds)
	v1Router.Post("/feeds-follow", apiCfg.middlewareAuth(apiCfg.handleCreateFeedsFollow))
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
