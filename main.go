package mai

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Todo struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"creation_date"`
	StartedAt  time.Time `json:"start_date,omitempty"`
	FinishedAt time.Time `json:"end_date,omitempty"`
}

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"mensagem": "pong"})
	})

	fmt.Println("Servidor rodando na porta 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor", err)
	}
}
