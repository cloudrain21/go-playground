package main

import (
	"fmt"
	rm "github.com/cloudrain21/go-playground/reservemovie"
	"time"
)

func main() {
	theater := &rm.Theater{}

	theater.AddMovie( rm.NewMovie( "Titanic", 1 * time.Hour, rm.Money{10}) )
	theater.AddMovie( rm.NewMovie( "Alien", 2 * time.Hour, rm.Money{15}) )

	theater.ShowMovieList()

	audience := rm.Audience{
		rm.Count{10},
		rm.Money{1000},
	}

	reserveMovie := rm.Movie{
		"Alien",
		2 * time.Hour,
		rm.Money{15},
	}

	ret := theater.Reserve(reserveMovie, audience)
	if ret {
		fmt.Println("reserve success")
	} else {
		fmt.Println("reserve fail")
	}
}
