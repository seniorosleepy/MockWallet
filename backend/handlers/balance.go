package handlers

import (
	"encoding/json"
	"mockwallet/db"
	"mockwallet/middleware"
	"mockwallet/models"
	"net/http"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int)

	rows, err := db.GetDB().Query("SELECT curr_id, amount FROM balance WHERE user_id = $1", userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var balances []models.Balance
	for rows.Next() {
		var b models.Balance
		if err := rows.Scan(&b.CurrId, &b.Amount); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		balances = append(balances, b)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balances)
}
