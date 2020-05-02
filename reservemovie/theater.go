package reservemovie

import "fmt"

type Theater struct {
	movieList []Movie
}

func (t *Theater) AddMovie(movie *Movie) {
	t.movieList = append(t.movieList, *movie)
}

func (t Theater) ShowMovieList() {
	for _, movie := range t.movieList {
		fmt.Println(movie)
	}
}

func (t *Theater) Reserve(movie Movie, audience Audience) bool {
	return true
}
