package handlers

import (
	"encoding/json"
	"net/http"
	"reactor-game/backend/models"

	"github.com/jmoiron/sqlx"
)

type UserResponse struct {
	ID            int              `json:"id"`
	Name          string           `json:"name"`
	Balance       int              `json:"balance"`
	ActiveReactor int              `json:"active_reactor"`
	Reactors      []models.Reactor `json:"reactors"`
}

func GetUser(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User //id, balance,active_reactor
		userID := 1
		err := db.Get(&user, "SELECT * FROM users WHERE id=$1", userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		var reactors []models.Reactor
		err = db.Select(&reactors, `
		SELECT reactors.* FROM user_reactors
		JOIN reactors ON user_reactors.reactor_id = reactors.id
		WHERE user_reactors.user_id = $1`, user.ID)

		if err != nil {
			http.Error(w, "Failed to load user reactors", http.StatusInternalServerError)
			return
		}

		resp := UserResponse{
			ID:            user.ID,
			Name:          user.Name,
			Balance:       user.Balance,
			ActiveReactor: user.ActiveReactor,
			Reactors:      reactors,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

	}
}
