package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"reactor-game/backend/models"
	"sort"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type TelegramAuthRequest struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Hash      string `json:"hash"`
	AuthDate  int64  `json:auth_date`
}

type AuthResponse struct {
	UserID int `json:"user_id"`
}

func AuthenticateTelegram(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req TelegramAuthRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		//Проверяем, что данные от Telegram настоящие
		if !verifyTelegramHash(req, "8023897789:AAGpFMh6tpGQr6Wnc5_ICz8DYyYZpB0qYIc") {
			http.Error(w, "Invalid Telegram hash", http.StatusUnauthorized)
			return
		}

		var user models.User
		err = db.Get(&user, "SELECT * FROM users WHERE telegram_id=$1", req.ID)
		if err != nil && err.Error() != "sql: no rows in result set" {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		//Если пользователя нет, создаем нового
		if err != nil {
			err = db.Get(&user, `
			INSERT INTO users(name, balance, active_reactor, farm_status, telegram_id)
			VALUES ($1,$2,$3,$4,$5)
			RETURNING *`, req.FirstName, 50, 1, "start", req.ID)
			if err != nil {
				http.Error(w, "Failed to create new user", http.StatusInternalServerError)
				return
			}

			//Даем пользователю новый реактор

			_, err = db.Exec(`
			INSERT INTO user_reactors (user_id, reactor_id)
			VALUES ($1, $2)`, user.ID, 1)

			if err != nil {
				http.Error(w, "Failed to assign initial reactor", http.StatusInternalServerError)
				return
			}
		}

		//Отправляем ID пользователя обратно
		resp := AuthResponse{UserID: user.ID}
		w.Header().Set("Content-Type:", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

// Функция проверки подписи Telegram
func verifyTelegramHash(req TelegramAuthRequest, botToken string) bool {
	//Сборка данных в строку для проверки
	dataCheckString := []string{
		"auth_date=" + strconv.FormatInt(req.AuthDate, 10),
		"first_name=" + req.FirstName,
		"id=" + strconv.FormatInt(req.ID, 10),
	}

	sort.Strings(dataCheckString)
	checkString := strings.Join(dataCheckString, "\n")

	secretKey := sha256.Sum256([]byte(botToken))
	h := hmac.New(sha256.New, secretKey[:])
	h.Write([]byte(checkString))
	expectedHash := hex.EncodeToString(h.Sum(nil))

	return expectedHash == req.Hash
}
