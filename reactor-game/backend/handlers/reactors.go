package handlers

import (
	"encoding/json"
	"net/http"
	"reactor-game/backend/models"

	"github.com/jmoiron/sqlx"
)

func GetReactors(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reactors []models.Reactor
		err := db.Select(&reactors, "SELECT * FROM reactors")

		if err != nil {
			http.Error(w, "Failed to load reactors", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reactors)
	}
}
