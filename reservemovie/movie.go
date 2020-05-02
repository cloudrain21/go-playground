package reservemovie

import "time"

type Movie struct {
	Name        string
	RunningTime time.Duration
	Price       Money
}
