package mgomodel

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/mapstructure"
	"sync"
	"time"
)

type Movie struct {
	PublicFields `bson:",inline"`
	MovieName    string    `bson:"movie_name" mapstructure:"movie_name"`
	Description  string    `bson:"description"`
	Thumb        string    `bson:"thumb"`
	ReleaseTime  time.Time `bson:"release_time" mapstructure:"release_time"`
	BoxOffice    float64   `bson:"box_office" mapstructure:"box_office"`
}

type MovieInfo struct {
	ID          string  `json:"id"`
	MovieName   string  `json:"movieName"`
	Description string  `json:"description"`
	Thumb       string  `json:"thumb"`
	ReleaseTime string  `json:"releaseTime"`
	BoxOffice   float64 `json:"boxOffice"`
}

type MovieList struct {
	Lock  *sync.Mutex
	IdMap map[string]*MovieInfo
}

func AddMovie(data map[string]interface{}) error {
	var movie Movie
	if err := mapstructure.Decode(data, &movie); err != nil {
		return err
	}
	movie.SetFieldsValue()
	return GetSession("self").DB("apiserver").Collection("movie").
		Insert(movie)
}

func EditMovieByID(id bson.ObjectId, data bson.M) error {
	data["updated_at"] = time.Now()
	return GetSession("self").DB("apiserver").Collection("movie").
		Update(bson.M{"_id": id}, bson.M{"$set": data})
}

func GetMovie(id bson.ObjectId) (*Movie, error) {
	var movie Movie
	err := GetSession("self").DB("apiserver").Collection("movie").
		FindOne(bson.M{"_id": id}, bson.M{}, &movie)
	if err != nil && err != mgo.ErrNotFound {
		return nil, err
	}
	if movie.ID == "" {
		return nil, nil
	}
	return &movie, nil
}

func GetMovieList(w map[string]interface{}, offset, limit uint64) ([]*Movie, uint64, error) {
	movies := make([]*Movie, 0)
	var count int
	count, err := GetSession("self").DB("apiserver").Collection("movie").Count(w)
	if err != nil {
		return movies, uint64(count), err
	}
	if err := GetSession("self").DB("apiserver").Collection("movie").
		FindList(int(offset), int(limit), w, bson.M{}, &movies, "-_id"); err != nil {
		return movies, uint64(count), err
	}
	return movies, uint64(count), nil
}

func DeleteMovie(id bson.ObjectId) error {
	if err := GetSession("self").DB("apiserver").Collection("movie").
		Remove(bson.M{"_id": id}); err != nil && err != mgo.ErrNotFound {
		return err
	}
	return nil
}
