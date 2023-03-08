package util

import "apiserver/pkg/constvar"

type PageSetting struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
	Page   uint64 `json:"page"`
}
type Page struct {
	CurrentPage uint64      `json:"currentPage"`
	PageSize    uint64      `json:"pageSize"`
	TotalPage   uint64      `json:"totalPage"`
	TotalCount  uint64      `json:"totalCount"`
	FirstPage   bool        `json:"firstPage"`
	LastPage    bool        `json:"lastPage"`
	List        interface{} `json:"list"`
}

func (s *PageSetting) Setting(page, limit uint64) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	if page == 0 {
		page = 1
	}
	s.Limit = limit
	s.Page = page
	s.Offset = (page - 1) * s.Limit
}

func PageUtil(count uint64, currentPage uint64, pageSize uint64, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{CurrentPage: currentPage, PageSize: pageSize, TotalPage: tp, TotalCount: count,
		FirstPage: currentPage == 1, LastPage: currentPage == tp, List: list}
}

func MysqlPagination(ps PageSetting) (offset, limit int) {
	if ps.Offset == 0 && ps.Limit == 0 {
		return 0, -1
	}
	return int(ps.Offset), int(ps.Limit)
}
