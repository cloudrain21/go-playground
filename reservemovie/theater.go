package reservemovie

import "fmt"

type Theater struct {
	MovieList []Movie
}

func (t *Theater) AddMovie(movie Movie) {
	t.MovieList = append(t.MovieList, *movie)
}

func (t Theater) ShowMovieList() {
	for _, movie := range t.MovieList {
		fmt.Println(movie)
	}
}

func (t *Theater) Reserve(movie Movie, audience Audience) bool {
	return true
}
