package user

import (
	"encoding/json"
	"net/http"
)

func ListAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{Id: 1, Email: "teste@teste.com", Nome: "Teste Teste"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
