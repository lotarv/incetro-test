package handlers

import (
	"encoding/json"
	"net/http"
	"reactor-game/backend/models"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func BuyReactor(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := 1

		//Взять id реактора из запроса
		reactorID := chi.URLParam(r, "id")

		//Взять реактор
		var reactor models.Reactor
		err := db.Get(&reactor, "SELECT * FROM reactors WHERE id=$1", reactorID)
		if err != nil {
			http.Error(w, "Reactor not found", http.StatusNotFound)
			return
		}
		//Взять текущего пользователя
		var user models.User
		err = db.Get(&user, "SELECT * FROM users WHERE id = $1", userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		//Проверить, не куплен ли уже реактор
		var count int
		err = db.Get(&count, `
		SELECT COUNT(*) FROM user_reactors WHERE
		user_id=$1 AND reactor_id=$2`, userID, reactorID)
		if err != nil {
			http.Error(w, "Failed to check ownership", http.StatusInternalServerError)
			return
		}
		if count > 0 {
			http.Error(w, "Reactor already owned", http.StatusBadRequest)
			return
		}

		//Проверить баланс пользователя
		if user.Balance < reactor.Price {
			http.Error(w, "Not enough tokens to buy this reactor", http.StatusBadRequest)
			return
		}
		//Если баланс позволяет, добавить реактор в таблицу user_reactors
		tx, err := db.Beginx()
		if err != nil {
			http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
			return
		}
		_, err = tx.Exec(`
		UPDATE users
		SET balance = balance - $1
		WHERE id = $2`, reactor.Price, userID)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to update balance", http.StatusInternalServerError)
			return
		}

		_, err = tx.Exec(`
		INSERT INTO user_reactors (user_id, reactor_id)
		VALUES ($1, $2)`, userID, reactorID)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to buy reactor", http.StatusInternalServerError)
			return
		}

		err = tx.Commit()
		if err != nil {
			http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Reactor bought"))
	}
}

func UseReactor(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Взять пользователя
		var user models.User
		userID := 1
		err := db.Get(&user, "SELECT * FROM users WHERE id = $1", userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		//Взять айди реактора
		reactor_id_request := chi.URLParam(r, "id")
		reactorID, err := strconv.Atoi(reactor_id_request)
		if err != nil {
			http.Error(w, "Invalid reactor id", http.StatusBadRequest)
			return
		}

		//Если уже установлен как активный, не тратим ресурсы
		if user.ActiveReactor == reactorID {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Reactor was already active"))
			return
		}
		//Проверить, не запущен ли фарм
		if user.FarmStatus == "farming" {
			http.Error(w, "Cannot change reactor while farming", http.StatusBadRequest)
			return
		}
		//Проверить, что пользователь владеет реактором
		var count int
		err = db.Get(&count, `
		SELECT COUNT(*) FROM user_reactors
		WHERE user_id=$1 AND reactor_id=$2`, userID, reactorID)
		if err != nil {
			http.Error(w, "Failed to check ownership", http.StatusInternalServerError)
			return
		}
		if count == 0 {
			http.Error(w, "Reactor not owned", http.StatusBadRequest)
			return
		}
		//Если владеет, установить активный реактор на этот айди
		_, err = db.Exec(`
		UPDATE users
		SET active_reactor = $1
		WHERE id = $2`, reactorID, userID)

		if err != nil {
			http.Error(w, "Failed to set active reactor", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Reactor set as active"))
	}
}
