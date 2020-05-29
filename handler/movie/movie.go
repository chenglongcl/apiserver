package movie

type CreateRequest struct {
	MovieName   string  `json:"movieName" binding:"required"`
	Description string  `json:"description"`
	Thumb       string  `json:"thumb"`
	ReleaseTime string  `json:"releaseTime"`
	BoxOffice   float64 `json:"boxOffice"`
}

type UpdateRequest struct {
	ID          string  `json:"id" binding:"len=24"`
	MovieName   string  `json:"movieName" binding:"required"`
	Description string  `json:"description"`
	Thumb       string  `json:"thumb"`
	ReleaseTime string  `json:"releaseTime"`
	BoxOffice   float64 `json:"boxOffice"`
}

type GetRequest struct {
	ID string `form:"id" binding:"len=24"`
}

type GetResponse struct {
	ID          string  `json:"movieID"`
	MovieName   string  `json:"movieName"`
	Description string  `json:"description"`
	Thumb       string  `json:"thumb"`
	ReleaseTime string  `json:"releaseTime"`
	BoxOffice   float64 `json:"boxOffice"`
}

type ListRequest struct {
	Page  uint64 `form:"page"`
	Limit uint64 `form:"limit"`
}

type DeleteRequest struct {
	ID string `form:"id" binding:"len=24"`
}
