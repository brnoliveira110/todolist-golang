package mai

import "time"

type Todo struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"creation_date"`
	StartedAt  time.Time `json:"start_date,omitempty"`
	FinishedAt time.Time `json:"end_date,omitempty"`
}

func main() {

}
