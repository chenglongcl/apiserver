package esmodel

import "github.com/olivere/elastic/v7"

type UserEs struct {
	ID          uint64 `json:"id" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Mobile      string `json:"mobile" binding:"required"`
	Age         uint64 `json:"age" binding:"required"`
	Sex         int64  `json:"sex"`
	CreatedTime string `json:"createdTime,omitempty"`
	UpdatedTime string `json:"updatedTime,omitempty"`
}

type SearchUserRequest struct {
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Page     int    `json:"page" binding:"min=1"`
	Limit    int    `json:"limit" binding:"min=1"`
}

func (r *SearchUserRequest) ToFilter() *EsSearch {
	search := &EsSearch{}
	if len(r.Username) != 0 {
		search.ShouldQuery = append(search.ShouldQuery, elastic.NewMatchQuery("username", r.Username))
	}
	if len(r.Mobile) != 0 {
		search.ShouldQuery = append(search.ShouldQuery, elastic.NewTermsQuery("mobile", r.Mobile))
	}
	if search.Sorters == nil {
		search.Sorters = append(search.Sorters, elastic.NewScoreSort().Desc(),
			elastic.NewFieldSort("createdTime").Desc())
	}
	search.From, search.Size = search.SetPaging(r.Page, r.Limit)
	return search
}
