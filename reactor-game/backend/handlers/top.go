package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type UserInfo struct {
	Name    string `json:"username"`
	Balance int    `json:"balance"`
}

type TopResponse struct {
	Rating []UserInfo `json:"rating"`
}

func GetUsersRating(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var top_100_users []UserInfo
		err := db.Select(&top_100_users, `
		SELECT name, balance FROM users ORDER BY balance DESC`)
		if err != nil {
			http.Error(w, "Failed to fetch users rating", http.StatusInternalServerError)
			return
		}

		if len(top_100_users) > 100 {
			top_100_users = top_100_users[:100]
		}

		resp := TopResponse{
			Rating: top_100_users,
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

	}
}
