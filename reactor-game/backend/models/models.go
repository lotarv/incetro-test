package models

import "time"

type User struct {
	ID            int        `db:"id" json:"id"`
	Name          string     `db:"name" json:"name"`
	Balance       int        `db:"balance" json:"balance"`
	ActiveReactor int        `db:"active_reactor" json:"active_reactor"`
	FarmStatus    string     `db:"farm_status" json:"farm_status"`
	FarmStartTime *time.Time `db:"farm_start_time" json:"farm_start_time"` //Указатель, т.к может быть nil
	FarmProgress  int        `db:"farm_progress" json:"farm_progress"`
}

type Reactor struct {
	ID             int `db:"id" json:"id"`
	FarmTime       int `db:"farm_time" json:"farm_time"` // в секундах
	TokensPerCycle int `db:"tokens_per_cycle" json:"tokens_per_cycle"`
	Price          int `db:"price" json:"price"`
}
