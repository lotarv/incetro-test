package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reactor-game/backend/models"
	"sort"
	"strings"

	"github.com/jmoiron/sqlx"
)

type TelegramAuthRequest struct {
	InitData string `json:"initData"`
}

type AuthResponse struct {
	UserID int `json:"user_id"`
}

func AuthenticateTelegram(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("AuthenticateTelegram function is called")

		var req TelegramAuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("Error decoding JSON: %v", err)
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		log.Printf("Received initData: %s", req.InitData)

		userID, firstName, isPremium, valid := CheckTelegramAuth(req.InitData)
		if !valid {
			log.Printf("Authorization failed")
			http.Error(w, "Invalid Telegram hash", http.StatusUnauthorized)
			return
		}

		log.Printf("Authenticated user: ID=%d, FirstName=%s, IsPremium=%v", userID, firstName, isPremium)

		var user models.User
		err := db.Get(&user, "SELECT * FROM users WHERE telegram_id=$1", userID)
		if err != nil && err.Error() != "sql: no rows in result set" {
			log.Printf("Database error: %v", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		// Если пользователя нет, создаём нового
		if err != nil {
			name := firstName
			if name == "" {
				name = "User"
			}
			err = db.Get(&user, `
				INSERT INTO users(name, balance, active_reactor, farm_status, telegram_id)
				VALUES ($1,$2,$3,$4,$5)
				RETURNING *`, name, 50, 1, "start", userID)
			if err != nil {
				log.Printf("Failed to create new user: %v", err)
				http.Error(w, "Failed to create new user", http.StatusInternalServerError)
				return
			}

			// Даём пользователю начальный реактор
			_, err = db.Exec(`
				INSERT INTO user_reactors (user_id, reactor_id)
				VALUES ($1, $2)`, user.ID, 1)
			if err != nil {
				log.Printf("Failed to assign initial reactor: %v", err)
				http.Error(w, "Failed to assign initial reactor", http.StatusInternalServerError)
				return
			}
		}

		// Отправляем ID пользователя обратно
		resp := AuthResponse{UserID: user.ID}
		log.Printf("Successful authorization: %+v", resp)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func CheckTelegramAuth(initData string) (int64, string, bool, bool) {
	parsedData, _ := url.QueryUnescape(initData)
	chunks := strings.Split(parsedData, "&")
	var dataPairs [][]string
	hash := ""
	user := &struct {
		ID        int64  `json:"id"`
		FirstName string `json:"first_name"` // Изменили с Username на FirstName
		IsPremium bool   `json:"is_premium"`
	}{}

	// Filter and split the chunks
	for _, chunk := range chunks {
		if strings.HasPrefix(chunk, "user=") {
			parsedData = strings.TrimPrefix(chunk, "user=")
			if err := json.Unmarshal([]byte(parsedData), user); err != nil {
				log.Printf("Failed to unmarshal user data: %v", err)
				return 0, "", false, false
			}
		}
		if strings.HasPrefix(chunk, "hash=") {
			hash = strings.TrimPrefix(chunk, "hash=")
		} else {
			pair := strings.SplitN(chunk, "=", 2)
			dataPairs = append(dataPairs, pair)
		}
	}

	// Sort the data pairs by the key
	sort.Slice(dataPairs, func(i, j int) bool {
		return dataPairs[i][0] < dataPairs[j][0]
	})

	// Join the sorted data pairs into the initData string
	var sortedData []string
	for _, pair := range dataPairs {
		sortedData = append(sortedData, fmt.Sprintf("%s=%s", pair[0], pair[1]))
	}
	dataCheckString := strings.Join(sortedData, "\n")

	log.Printf("Data check string: %s", dataCheckString)

	// Create the secret key using HMAC and the given token
	h := hmac.New(sha256.New, []byte("WebAppData"))
	h.Write([]byte("7767502100:AAFMs9ALGfcQ1Mik1hbGr66Nb9hDQCBe-yU")) // Убедись, что токен верный
	secretKey := h.Sum(nil)

	// Create the data check using the secret key and initData
	h = hmac.New(sha256.New, secretKey)
	h.Write([]byte(dataCheckString))
	dataCheck := h.Sum(nil)

	log.Printf("Expected hash: %x", dataCheck)
	log.Printf("Received hash: %s", hash)
	log.Printf("Hash comparison: %v", fmt.Sprintf("%x", dataCheck) == hash)

	return user.ID, user.FirstName, user.IsPremium, fmt.Sprintf("%x", dataCheck) == hash
}
