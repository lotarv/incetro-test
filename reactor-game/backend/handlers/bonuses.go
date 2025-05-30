package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reactor-game/backend/models"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type BonusResponse struct {
	Balance       int            `json:"balance"`
	ActiveReactor models.Reactor `json:"active_reactor"`
	FarmStatus    string         `json:"farm_status"`
	Progress      int            `json:"progress"`
	TimeLeft      int            `json:"time_left"` //в секундах
}

func GetBonuses(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//1. Находим пользователя
		userIDStr := r.URL.Query().Get("userID")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}
		log.Print("Trying to get Bonuses for user with id=", userID)
		var user models.User
		err = db.Get(&user, "SELECT * FROM users WHERE id=$1", userID)
		if err != nil {
			http.Error(w, "user not found", http.StatusNotFound)
			log.Fatal(err)
			return
		}

		//1.5 Проверяем, есть ли активный реактор
		if user.ActiveReactor == 0 {
			http.Error(w, "no active reactor set", http.StatusBadRequest)
		}

		//2. Находим реактор
		var active_reactor models.Reactor
		err = db.Get(&active_reactor, "SELECT * FROM reactors WHERE id=$1", user.ActiveReactor)
		if err != nil {
			http.Error(w, "reactor not found", http.StatusInternalServerError)
			return
		}

		//3. Формируем ответ

		resp := BonusResponse{
			Balance:       user.Balance,
			ActiveReactor: active_reactor,
			FarmStatus:    user.FarmStatus,
			Progress:      user.FarmProgress,
			TimeLeft:      0,
		}

		//4. Обработка состояния фарминга
		if user.FarmStatus == "farming" && user.FarmStartTime != nil {
			elapsed := int(time.Since(*user.FarmStartTime).Seconds())
			total := active_reactor.FarmTime
			resp.Progress = min((elapsed*100)/total, 100)
			if elapsed >= active_reactor.FarmTime {
				resp.FarmStatus = "claim"
				resp.TimeLeft = 0
			} else {
				resp.TimeLeft = active_reactor.FarmTime - elapsed
			}
		}

		//5. отправка JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func StartFarming(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		userIDStr := r.URL.Query().Get("userID")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}

		//Находим пользователя
		err = db.Get(&user, "SELECT * FROM users WHERE id=$1", userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		//Проверяем, не фармит ли он уже
		if user.FarmStatus != "start" {
			http.Error(w, "User is already farming", http.StatusBadRequest)
			return
		}

		//Обновляем статус и время начала фарма

		_, err = db.Exec(`
		UPDATE users
		SET farm_status=$1, farm_start_time=$2
		WHERE id = $3`, "farming", time.Now(), userID)
		if err != nil {
			http.Error(w, "Failed to start farming", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Farming started"))

	}
}

func ClaimBonuses(db *sqlx.DB) http.HandlerFunc { // /bonuses/claim
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		userIDStr := r.URL.Query().Get("userID")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}

		err = db.Get(&user, "SELECT * FROM users WHERE id = $1", userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		//Находим активный реактор

		var reactor models.Reactor

		err = db.Get(&reactor, "SELECT * FROM reactors WHERE id = $1", user.ActiveReactor)
		if err != nil {
			http.Error(w, fmt.Sprintf("Reactor not found: %v", err), http.StatusNotFound)
			return
		}

		//Проверяем, может ли он заклеймить поинты

		if user.FarmStatus != "farming" || user.FarmStartTime.IsZero() {
			http.Error(w, "Farming not started of already claimed", http.StatusBadRequest)
			return
		}

		elapsed := int(time.Since(*user.FarmStartTime).Seconds())
		if elapsed < reactor.FarmTime {
			http.Error(w, "Farming not yet complete", http.StatusBadRequest)
			return
		}

		//Начисляем токены и сбрасываем статус
		newBalance := user.Balance + reactor.TokensPerCycle
		_, err = db.Exec(`
		UPDATE users
		SET balance=$1, farm_status = $2, farm_start_time=NULL
		WHERE id = $3`, newBalance, "start", userID)

		if err != nil {
			http.Error(w, "Failed to claim bonuses", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Bonuses claimed"))
	}
}
