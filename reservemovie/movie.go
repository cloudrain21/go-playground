package reservemovie

import "time"

type Movie struct {
	name        string
	runningTime time.Duration
	price       Money
}

func NewMovie(name string, runningTime time.Duration, price Money) *Movie {
	return &Movie{name, runningTime, price}
}