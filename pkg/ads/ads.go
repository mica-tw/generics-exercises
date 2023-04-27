package ads

import "time"

type Ad struct {
	Title       string
	Description string
	CreatedAt   time.Time
	ID          string
	Price       int
}
