package reservemovie

import "time"

type Movie struct {
	name        string
	runningTime time.Duration
	price       Money
}
