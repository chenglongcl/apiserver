package movieservice

import (
	"apiserver/mgomodel"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"sync"
	"time"
)

type Movie struct {
	ID          bson.ObjectId
	MovieName   string
	Description string
	Thumb       string
	ReleaseTime time.Time
	BoxOffice   float64
}

func (a *Movie) Add() *errno.Errno {
	data := map[string]interface{}{
		"movie_name":   a.MovieName,
		"description":  a.Description,
		"thumb":        a.Thumb,
		"release_time": a.ReleaseTime,
		"box_office":   a.BoxOffice,
	}
	if err := mgomodel.AddMovie(data); err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *Movie) Edit() *errno.Errno {
	data := bson.M{
		"movie_name":   a.MovieName,
		"description":  a.Description,
		"thumb":        a.Thumb,
		"release_time": a.ReleaseTime,
		"box_office":   a.BoxOffice,
	}
	if err := mgomodel.EditMovieByID(a.ID, data); err != nil {
		if err == mgo.ErrNotFound {
			return errno.ErrRecordNotFound
		}
		return errno.ErrDatabase
	}
	return nil
}

func (a *Movie) Get() (*mgomodel.Movie, *errno.Errno) {
	movie, err := mgomodel.GetMovie(a.ID)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return movie, nil
}

func (a *Movie) GetList(ps util.PageSetting) ([]*mgomodel.MovieInfo, uint64, *errno.Errno) {
	w := make(map[string]interface{}, 0)
	if a.MovieName != "" {
		w["movie_name"] = bson.RegEx{Pattern: a.MovieName}
	}
	movies, count, err := mgomodel.GetMovieList(w, ps.Offset, ps.Limit)
	if err != nil {
		return nil, count, errno.ErrDatabase
	}
	var ids []string
	for _, movie := range movies {
		ids = append(ids, movie.ID.Hex())
	}
	info := make([]*mgomodel.MovieInfo, 0)
	wg := sync.WaitGroup{}
	movieList := mgomodel.MovieList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[string]*mgomodel.MovieInfo, len(movies)),
	}
	finished := make(chan bool, 1)
	for _, movie := range movies {
		wg.Add(1)
		go func(movie *mgomodel.Movie) {
			defer wg.Done()
			movieList.Lock.Lock()
			defer movieList.Lock.Unlock()
			movieList.IdMap[movie.ID.Hex()] = &mgomodel.MovieInfo{
				ID:          movie.ID.Hex(),
				MovieName:   movie.MovieName,
				Description: movie.Description,
				Thumb:       movie.Thumb,
				ReleaseTime: movie.ReleaseTime.Local().Format("2006-01-02 15:04:05"),
				BoxOffice:   movie.BoxOffice,
			}
		}(movie)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	}

	for _, id := range ids {
		info = append(info, movieList.IdMap[id])
	}
	return info, count, nil
}

func (a *Movie) Delete() *errno.Errno {
	if err := mgomodel.DeleteMovie(a.ID); err != nil {
		if err == mgo.ErrNotFound {
			return errno.ErrRecordNotFound
		}
		return errno.ErrDatabase
	}
	return nil
}
