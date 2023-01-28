// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package apiserverquery

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q           = new(Query)
	TbArticle   *tbArticle
	TbUser      *tbUser
	TbUserToken *tbUserToken
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	TbArticle = &Q.TbArticle
	TbUser = &Q.TbUser
	TbUserToken = &Q.TbUserToken
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:          db,
		TbArticle:   newTbArticle(db, opts...),
		TbUser:      newTbUser(db, opts...),
		TbUserToken: newTbUserToken(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	TbArticle   tbArticle
	TbUser      tbUser
	TbUserToken tbUserToken
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:          db,
		TbArticle:   q.TbArticle.clone(db),
		TbUser:      q.TbUser.clone(db),
		TbUserToken: q.TbUserToken.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:          db,
		TbArticle:   q.TbArticle.replaceDB(db),
		TbUser:      q.TbUser.replaceDB(db),
		TbUserToken: q.TbUserToken.replaceDB(db),
	}
}

type queryCtx struct {
	TbArticle   ITbArticleDo
	TbUser      ITbUserDo
	TbUserToken ITbUserTokenDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		TbArticle:   q.TbArticle.WithContext(ctx),
		TbUser:      q.TbUser.WithContext(ctx),
		TbUserToken: q.TbUserToken.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
