package esmodel

import (
	"apiserver/pkg/constvar"
	"github.com/olivere/elastic/v7"
)

//bool query 条件
type EsSearch struct {
	MustQuery    []elastic.Query
	MustNotQuery []elastic.Query
	ShouldQuery  []elastic.Query
	Filters      []elastic.Query
	Sorters      []elastic.Sorter
	From         int //分页
	Size         int
}

func (a *EsSearch) SetPaging(page, limit int) (int, int) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	if page == 0 {
		page = 1
	}
	return (page - 1) * limit, limit
}
