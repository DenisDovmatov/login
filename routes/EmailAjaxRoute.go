package routes

import (
	"encoding/json"
	"fmt"
	"loginB/db"
	"net/http"
)


// Проверка почты через ajax
func HaveEmailHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		email := r.FormValue("email")
		fmt.Println("Проверка почты в базе ", email)

		have, _ := db.HaveDBUser(email)

		if have {
			fmt.Println("Почта занята")

		}else {
			fmt.Println("Почта свободна")
		}
		// Создаем json для отправки клиенту
		a, err := json.Marshal(have)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(a)
	}
}
