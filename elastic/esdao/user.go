package esdao

import (
	"apiserver/elastic/esmodel"
	"context"
	"encoding/json"
	"fmt"
	"github.com/chenglongcl/log"
	"github.com/olivere/elastic/v7"
	"reflect"
	"strconv"
	"time"
)

const (
	author       = "chenglongcl"
	project      = "apiserver_user"
	esRetryLimit = 3 //bulk 错误重试机制
)

type UserES struct {
	index  string
	client *elastic.Client
}

func NewUserES(client *elastic.Client) *UserES {
	index := fmt.Sprintf("%s_%s", author, project)
	userEs := &UserES{
		client: client,
		index:  index,
	}
	return userEs
}

func (es *UserES) BatchAdd(ctx context.Context, user []*esmodel.UserEs) error {
	var err error
	for i := 0; i < esRetryLimit; i++ {
		if err = es.batchAdd(ctx, user); err != nil {
			log.Error("batch add failed ", err)
			continue
		}
		return err
	}
	return err
}

func (es *UserES) batchAdd(ctx context.Context, user []*esmodel.UserEs) error {
	req := es.client.Bulk().Index(es.index)
	for _, u := range user {
		u.CreatedTime = time.Now().Format("2006-01-02 15:04:05")
		u.UpdatedTime = time.Now().Format("2006-01-02 15:04:05")
		doc := elastic.NewBulkIndexRequest().Id(strconv.FormatUint(u.ID, 10)).Doc(u)
		req.Add(doc)
	}
	if req.NumberOfActions() < 0 {
		return nil
	}
	res, err := req.Do(ctx)
	if err != nil {
		return err
	}
	// 任何子请求失败，该 `errors` 标志被设置为 `true` ，并且在相应的请求报告出错误明细
	// 所以如果没有出错，说明全部成功了，直接返回即可
	if !res.Errors {
		return nil
	}
	for _, it := range res.Failed() {
		if it.Error == nil {
			continue
		}
		return &elastic.Error{
			Status:  it.Status,
			Details: it.Error,
		}
	}
	return nil
}

func (es *UserES) BatchUpdate(ctx context.Context, user []*esmodel.UserEs) error {
	var err error
	for i := 0; i < esRetryLimit; i++ {
		if err = es.batchUpdate(ctx, user); err != nil {
			continue
		}
		return err
	}
	return err
}

func (es *UserES) batchUpdate(ctx context.Context, user []*esmodel.UserEs) error {
	req := es.client.Bulk().Index(es.index)
	for _, u := range user {
		u.UpdatedTime = time.Now().Format("2006-01-02 15:04:05")
		doc := elastic.NewBulkUpdateRequest().Id(strconv.FormatUint(u.ID, 10)).Doc(u)
		req.Add(doc)
	}

	if req.NumberOfActions() < 0 {
		return nil
	}
	res, err := req.Do(ctx)
	if err != nil {
		return err
	}
	// 任何子请求失败，该 `errors` 标志被设置为 `true` ，并且在相应的请求报告出错误明细
	// 所以如果没有出错，说明全部成功了，直接返回即可
	if !res.Errors {
		return nil
	}
	for _, it := range res.Failed() {
		if it.Error == nil {
			continue
		}
		return &elastic.Error{
			Status:  it.Status,
			Details: it.Error,
		}
	}

	return nil
}

// 根据id 批量获取
func (es *UserES) MGet(ctx context.Context, IDS []uint64) ([]*esmodel.UserEs, error) {
	userES := make([]*esmodel.UserEs, 0, len(IDS))
	idStr := make([]string, 0, len(IDS))
	for _, id := range IDS {
		idStr = append(idStr, strconv.FormatUint(id, 10))
	}
	resp, err := es.client.Search(es.index).Query(
		elastic.NewIdsQuery().Ids(idStr...)).Size(len(IDS)).Do(ctx)

	if err != nil {
		return nil, err
	}

	if resp.TotalHits() == 0 {
		return nil, nil
	}
	for _, e := range resp.Each(reflect.TypeOf(&esmodel.UserEs{})) {
		us := e.(*esmodel.UserEs)
		userES = append(userES, us)
	}
	return userES, nil
}

func (es *UserES) BatchDel(ctx context.Context, user []*esmodel.UserEs) error {
	var err error
	for i := 0; i < esRetryLimit; i++ {
		if err = es.batchDel(ctx, user); err != nil {
			continue
		}
		return err
	}
	return err
}

func (es *UserES) batchDel(ctx context.Context, user []*esmodel.UserEs) error {
	req := es.client.Bulk().Index(es.index)
	for _, u := range user {
		doc := elastic.NewBulkDeleteRequest().Id(strconv.FormatUint(u.ID, 10))
		req.Add(doc)
	}

	if req.NumberOfActions() < 0 {
		return nil
	}

	res, err := req.Do(ctx)
	if err != nil {
		return err
	}
	// 任何子请求失败，该 `errors` 标志被设置为 `true` ，并且在相应的请求报告出错误明细
	// 所以如果没有出错，说明全部成功了，直接返回即可
	if !res.Errors {
		return nil
	}
	for _, it := range res.Failed() {
		if it.Error == nil {
			continue
		}
		return &elastic.Error{
			Status:  it.Status,
			Details: it.Error,
		}
	}

	return nil
}

//根据id 单个获取
func (es *UserES) Get(ctx context.Context, id uint64) (*esmodel.UserEs, error) {
	user := &esmodel.UserEs{}
	resp, err := es.client.Get().Index(es.index).Id(strconv.FormatUint(id, 10)).Do(ctx)
	if err != nil && !elastic.IsNotFound(err) {
		return nil, err
	}
	if elastic.IsNotFound(err) {
		return nil, nil
	}
	source, err := resp.Source.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(source, &user); err != nil {
		return nil, err
	}
	return user, nil
}

func (es *UserES) Search(ctx context.Context, filter *esmodel.EsSearch) ([]*esmodel.UserEs, uint64, error) {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(filter.MustQuery...)
	boolQuery.MustNot(filter.MustNotQuery...)
	boolQuery.Should(filter.ShouldQuery...)
	boolQuery.Filter(filter.Filters...)

	// 当should不为空时，保证至少匹配should中的一项
	if len(filter.MustQuery) == 0 && len(filter.MustNotQuery) == 0 && len(filter.ShouldQuery) > 0 {
		boolQuery.MinimumShouldMatch("1")
	}

	service := es.client.Search().Index(es.index).Query(boolQuery).SortBy(filter.Sorters...).From(filter.From).Size(filter.Size)
	resp, err := service.Do(ctx)
	if err != nil {
		return nil, 0, err
	}
	userES := make([]*esmodel.UserEs, 0)
	for _, e := range resp.Each(reflect.TypeOf(&esmodel.UserEs{})) {
		us := e.(*esmodel.UserEs)
		userES = append(userES, us)
	}
	return userES, uint64(resp.TotalHits()), nil
}
