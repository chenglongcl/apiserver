// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package apiserverquery

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"apiserver/dal/apiserverdb/apiservermodel"
)

func newTbArticle(db *gorm.DB, opts ...gen.DOOption) tbArticle {
	_tbArticle := tbArticle{}

	_tbArticle.tbArticleDo.UseDB(db, opts...)
	_tbArticle.tbArticleDo.UseModel(&apiservermodel.TbArticle{})

	tableName := _tbArticle.tbArticleDo.TableName()
	_tbArticle.ALL = field.NewAsterisk(tableName)
	_tbArticle.ID = field.NewUint64(tableName, "id")
	_tbArticle.UID = field.NewUint64(tableName, "uid")
	_tbArticle.CateID = field.NewUint64(tableName, "cate_id")
	_tbArticle.Title = field.NewString(tableName, "title")
	_tbArticle.Content = field.NewString(tableName, "content")
	_tbArticle.Images = field.NewString(tableName, "images")
	_tbArticle.CreatedAt = field.NewTime(tableName, "created_at")
	_tbArticle.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tbArticle.DeletedAt = field.NewField(tableName, "deleted_at")

	_tbArticle.fillFieldMap()

	return _tbArticle
}

type tbArticle struct {
	tbArticleDo tbArticleDo

	ALL       field.Asterisk
	ID        field.Uint64
	UID       field.Uint64
	CateID    field.Uint64
	Title     field.String
	Content   field.String
	Images    field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (t tbArticle) Table(newTableName string) *tbArticle {
	t.tbArticleDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tbArticle) As(alias string) *tbArticle {
	t.tbArticleDo.DO = *(t.tbArticleDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tbArticle) updateTableName(table string) *tbArticle {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint64(table, "id")
	t.UID = field.NewUint64(table, "uid")
	t.CateID = field.NewUint64(table, "cate_id")
	t.Title = field.NewString(table, "title")
	t.Content = field.NewString(table, "content")
	t.Images = field.NewString(table, "images")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *tbArticle) WithContext(ctx context.Context) ITbArticleDo {
	return t.tbArticleDo.WithContext(ctx)
}

func (t tbArticle) TableName() string { return t.tbArticleDo.TableName() }

func (t tbArticle) Alias() string { return t.tbArticleDo.Alias() }

func (t *tbArticle) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tbArticle) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 10)
	t.fieldMap["id"] = t.ID
	t.fieldMap["uid"] = t.UID
	t.fieldMap["cate_id"] = t.CateID
	t.fieldMap["title"] = t.Title
	t.fieldMap["content"] = t.Content
	t.fieldMap["images"] = t.Images
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt

}

func (t tbArticle) clone(db *gorm.DB) tbArticle {
	t.tbArticleDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tbArticle) replaceDB(db *gorm.DB) tbArticle {
	t.tbArticleDo.ReplaceDB(db)
	return t
}

type tbArticleDo struct{ gen.DO }

type ITbArticleDo interface {
	gen.SubQuery
	Debug() ITbArticleDo
	WithContext(ctx context.Context) ITbArticleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITbArticleDo
	WriteDB() ITbArticleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITbArticleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITbArticleDo
	Not(conds ...gen.Condition) ITbArticleDo
	Or(conds ...gen.Condition) ITbArticleDo
	Select(conds ...field.Expr) ITbArticleDo
	Where(conds ...gen.Condition) ITbArticleDo
	Order(conds ...field.Expr) ITbArticleDo
	Distinct(cols ...field.Expr) ITbArticleDo
	Omit(cols ...field.Expr) ITbArticleDo
	Join(table schema.Tabler, on ...field.Expr) ITbArticleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITbArticleDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITbArticleDo
	Group(cols ...field.Expr) ITbArticleDo
	Having(conds ...gen.Condition) ITbArticleDo
	Limit(limit int) ITbArticleDo
	Offset(offset int) ITbArticleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITbArticleDo
	Unscoped() ITbArticleDo
	Create(values ...*apiservermodel.TbArticle) error
	CreateInBatches(values []*apiservermodel.TbArticle, batchSize int) error
	Save(values ...*apiservermodel.TbArticle) error
	First() (*apiservermodel.TbArticle, error)
	Take() (*apiservermodel.TbArticle, error)
	Last() (*apiservermodel.TbArticle, error)
	Find() ([]*apiservermodel.TbArticle, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*apiservermodel.TbArticle, err error)
	FindInBatches(result *[]*apiservermodel.TbArticle, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*apiservermodel.TbArticle) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITbArticleDo
	Assign(attrs ...field.AssignExpr) ITbArticleDo
	Joins(fields ...field.RelationField) ITbArticleDo
	Preload(fields ...field.RelationField) ITbArticleDo
	FirstOrInit() (*apiservermodel.TbArticle, error)
	FirstOrCreate() (*apiservermodel.TbArticle, error)
	FindByPage(offset int, limit int) (result []*apiservermodel.TbArticle, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITbArticleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tbArticleDo) Debug() ITbArticleDo {
	return t.withDO(t.DO.Debug())
}

func (t tbArticleDo) WithContext(ctx context.Context) ITbArticleDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tbArticleDo) ReadDB() ITbArticleDo {
	return t.Clauses(dbresolver.Read)
}

func (t tbArticleDo) WriteDB() ITbArticleDo {
	return t.Clauses(dbresolver.Write)
}

func (t tbArticleDo) Session(config *gorm.Session) ITbArticleDo {
	return t.withDO(t.DO.Session(config))
}

func (t tbArticleDo) Clauses(conds ...clause.Expression) ITbArticleDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tbArticleDo) Returning(value interface{}, columns ...string) ITbArticleDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tbArticleDo) Not(conds ...gen.Condition) ITbArticleDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tbArticleDo) Or(conds ...gen.Condition) ITbArticleDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tbArticleDo) Select(conds ...field.Expr) ITbArticleDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tbArticleDo) Where(conds ...gen.Condition) ITbArticleDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tbArticleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITbArticleDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tbArticleDo) Order(conds ...field.Expr) ITbArticleDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tbArticleDo) Distinct(cols ...field.Expr) ITbArticleDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tbArticleDo) Omit(cols ...field.Expr) ITbArticleDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tbArticleDo) Join(table schema.Tabler, on ...field.Expr) ITbArticleDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tbArticleDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITbArticleDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tbArticleDo) RightJoin(table schema.Tabler, on ...field.Expr) ITbArticleDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tbArticleDo) Group(cols ...field.Expr) ITbArticleDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tbArticleDo) Having(conds ...gen.Condition) ITbArticleDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tbArticleDo) Limit(limit int) ITbArticleDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tbArticleDo) Offset(offset int) ITbArticleDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tbArticleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITbArticleDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tbArticleDo) Unscoped() ITbArticleDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tbArticleDo) Create(values ...*apiservermodel.TbArticle) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tbArticleDo) CreateInBatches(values []*apiservermodel.TbArticle, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tbArticleDo) Save(values ...*apiservermodel.TbArticle) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tbArticleDo) First() (*apiservermodel.TbArticle, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*apiservermodel.TbArticle), nil
	}
}

func (t tbArticleDo) Take() (*apiservermodel.TbArticle, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*apiservermodel.TbArticle), nil
	}
}

func (t tbArticleDo) Last() (*apiservermodel.TbArticle, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*apiservermodel.TbArticle), nil
	}
}

func (t tbArticleDo) Find() ([]*apiservermodel.TbArticle, error) {
	result, err := t.DO.Find()
	return result.([]*apiservermodel.TbArticle), err
}

func (t tbArticleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*apiservermodel.TbArticle, err error) {
	buf := make([]*apiservermodel.TbArticle, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tbArticleDo) FindInBatches(result *[]*apiservermodel.TbArticle, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tbArticleDo) Attrs(attrs ...field.AssignExpr) ITbArticleDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tbArticleDo) Assign(attrs ...field.AssignExpr) ITbArticleDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tbArticleDo) Joins(fields ...field.RelationField) ITbArticleDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tbArticleDo) Preload(fields ...field.RelationField) ITbArticleDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tbArticleDo) FirstOrInit() (*apiservermodel.TbArticle, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*apiservermodel.TbArticle), nil
	}
}

func (t tbArticleDo) FirstOrCreate() (*apiservermodel.TbArticle, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*apiservermodel.TbArticle), nil
	}
}

func (t tbArticleDo) FindByPage(offset int, limit int) (result []*apiservermodel.TbArticle, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tbArticleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tbArticleDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tbArticleDo) Delete(models ...*apiservermodel.TbArticle) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tbArticleDo) withDO(do gen.Dao) *tbArticleDo {
	t.DO = *do.(*gen.DO)
	return t
}