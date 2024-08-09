package main

import (
	"database/sql"
	// "fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"github.com/AdarshShukla1001/first-go-server/internal/database"

	_ "github.com/lib/pq"
	
)


func respondWithError(w http.ResponseWriter,code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:")
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w,code,errResponse{
		Error: msg,
	})
}


type apiConfig struct {
	DB *database.Queries
}




func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}


	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)


	apiCfg := apiConfig{
		DB: dbQueries,
	}


	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	// On get request
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err",handlerErr)

	// Users Route
	v1Router.Post("/users",apiCfg.handlerCreateUser)
	v1Router.Get("/users",apiCfg.middlewareAuth(apiCfg.handlerGetUser))


	// Feeds
	v1Router.Post("/feeds",apiCfg.middlewareAuth(apiCfg.handlerFeedCreate))
	v1Router.Get("/feeds",apiCfg.handlerGetFeeds)

	
	// feedfollows
	v1Router.Post("/feed-follows",apiCfg.middlewareAuth(apiCfg.handlerFeedFollowCreate))


	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	// err := srv.ListenAndServe()

	// if err !
	
	log.Printf("Serving on port: %s\n", portString)
	log.Fatal(srv.ListenAndServe())

}
