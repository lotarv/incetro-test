package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"reactor-game/backend/models"
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
		var user models.User
		err := db.Get(&user, "SELECT * FROM users WHERE id=$1", 1)
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
