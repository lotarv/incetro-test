package main

// import (
// 	"github.com/go-chi/chi/v5"
// 	"log"
// 	"net/http"
// 	"reactor-game/backend/db"
// )

// func main() {
// 	database, err := db.NewDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer database.Close()

// 	router := chi.NewRouter()
// 	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Welcome to Reactor Game API"))
// 	})

// 	log.Println("Server starting on :8080...")
// 	http.ListenAndServe(":8080", router)
// }

import (
	"github.com/go-chi/chi/v5"
	"github.com/pressly/goose/v3"
	"log"
	"net/http"
	"reactor-game/backend/db"
	"reactor-game/backend/handlers"
)

func main() {
	database, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	err = goose.SetDialect("postgres")
	if err != nil {
		log.Fatalf("failed to set dialect: %v", err)
	}
	err = goose.Up(database.DB, "migrations")
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	router := chi.NewRouter()
	router.Get("/bonuses", handlers.GetBonuses(database))
	// router.Post("/bonuses/start", handlers.StartFarming(database))
	// router.Post("bonuses/claim", handlers.ClaimRewards(database))

	// router.Get("reactors", handlers.GetReactors(database))
	// router.Post("reactors/buy", handlers.BuyReactor(database))
	// router.Post("reactores/use", handlers.UseReactor(database))

	// router.Get("/top", handlers.GetLeaderboard(database))

	log.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", router)
}
