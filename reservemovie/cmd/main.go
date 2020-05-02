package main

import (
	"fmt"
	"time"
)

func main() {
	theater := &Theater{}

	theater.AddMovie(
		Movie{
			"Titanic",
			time.Hour,
			Money{10},
		})
	theater.AddMovie(
		Movie{
			"Alien",
			2 * time.Hour,
			Money{15},
		})

	theater.ShowMovieList()

	audience := Audience{
		Count{10},
		Money{1000},
	}

	reserveMovie := Movie{
		"Alien",
		2 * time.Hour,
		Money{15},
	}

	ret := theater.Reserve(reserveMovie, audience)
	if ret {
		fmt.Println("reserve success")
	} else {
		fmt.Println("reserve fail")
	}
}
