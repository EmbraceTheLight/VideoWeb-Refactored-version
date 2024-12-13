// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"vw_user/internal/data/dal/model"
)

func newFollowList(db *gorm.DB, opts ...gen.DOOption) followList {
	_followList := followList{}

	_followList.followListDo.UseDB(db, opts...)
	_followList.followListDo.UseModel(&model.FollowList{})

	tableName := _followList.followListDo.TableName()
	_followList.ALL = field.NewAsterisk(tableName)
	_followList.ListID = field.NewInt64(tableName, "list_id")
	_followList.ListName = field.NewString(tableName, "list_name")
	_followList.UserID = field.NewInt64(tableName, "user_id")

	_followList.fillFieldMap()

	return _followList
}

// followList 存放用户关注列表元信息
type followList struct {
	followListDo

	ALL      field.Asterisk
	ListID   field.Int64
	ListName field.String
	UserID   field.Int64

	fieldMap map[string]field.Expr
}

func (f followList) Table(newTableName string) *followList {
	f.followListDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f followList) As(alias string) *followList {
	f.followListDo.DO = *(f.followListDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *followList) updateTableName(table string) *followList {
	f.ALL = field.NewAsterisk(table)
	f.ListID = field.NewInt64(table, "list_id")
	f.ListName = field.NewString(table, "list_name")
	f.UserID = field.NewInt64(table, "user_id")

	f.fillFieldMap()

	return f
}

func (f *followList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *followList) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 3)
	f.fieldMap["list_id"] = f.ListID
	f.fieldMap["list_name"] = f.ListName
	f.fieldMap["user_id"] = f.UserID
}

func (f followList) clone(db *gorm.DB) followList {
	f.followListDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f followList) replaceDB(db *gorm.DB) followList {
	f.followListDo.ReplaceDB(db)
	return f
}

type followListDo struct{ gen.DO }

type IFollowListDo interface {
	gen.SubQuery
	Debug() IFollowListDo
	WithContext(ctx context.Context) IFollowListDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFollowListDo
	WriteDB() IFollowListDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFollowListDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFollowListDo
	Not(conds ...gen.Condition) IFollowListDo
	Or(conds ...gen.Condition) IFollowListDo
	Select(conds ...field.Expr) IFollowListDo
	Where(conds ...gen.Condition) IFollowListDo
	Order(conds ...field.Expr) IFollowListDo
	Distinct(cols ...field.Expr) IFollowListDo
	Omit(cols ...field.Expr) IFollowListDo
	Join(table schema.Tabler, on ...field.Expr) IFollowListDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFollowListDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFollowListDo
	Group(cols ...field.Expr) IFollowListDo
	Having(conds ...gen.Condition) IFollowListDo
	Limit(limit int) IFollowListDo
	Offset(offset int) IFollowListDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFollowListDo
	Unscoped() IFollowListDo
	Create(values ...*model.FollowList) error
	CreateInBatches(values []*model.FollowList, batchSize int) error
	Save(values ...*model.FollowList) error
	First() (*model.FollowList, error)
	Take() (*model.FollowList, error)
	Last() (*model.FollowList, error)
	Find() ([]*model.FollowList, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FollowList, err error)
	FindInBatches(result *[]*model.FollowList, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FollowList) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFollowListDo
	Assign(attrs ...field.AssignExpr) IFollowListDo
	Joins(fields ...field.RelationField) IFollowListDo
	Preload(fields ...field.RelationField) IFollowListDo
	FirstOrInit() (*model.FollowList, error)
	FirstOrCreate() (*model.FollowList, error)
	FindByPage(offset int, limit int) (result []*model.FollowList, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFollowListDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f followListDo) Debug() IFollowListDo {
	return f.withDO(f.DO.Debug())
}

func (f followListDo) WithContext(ctx context.Context) IFollowListDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f followListDo) ReadDB() IFollowListDo {
	return f.Clauses(dbresolver.Read)
}

func (f followListDo) WriteDB() IFollowListDo {
	return f.Clauses(dbresolver.Write)
}

func (f followListDo) Session(config *gorm.Session) IFollowListDo {
	return f.withDO(f.DO.Session(config))
}

func (f followListDo) Clauses(conds ...clause.Expression) IFollowListDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f followListDo) Returning(value interface{}, columns ...string) IFollowListDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f followListDo) Not(conds ...gen.Condition) IFollowListDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f followListDo) Or(conds ...gen.Condition) IFollowListDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f followListDo) Select(conds ...field.Expr) IFollowListDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f followListDo) Where(conds ...gen.Condition) IFollowListDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f followListDo) Order(conds ...field.Expr) IFollowListDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f followListDo) Distinct(cols ...field.Expr) IFollowListDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f followListDo) Omit(cols ...field.Expr) IFollowListDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f followListDo) Join(table schema.Tabler, on ...field.Expr) IFollowListDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f followListDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFollowListDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f followListDo) RightJoin(table schema.Tabler, on ...field.Expr) IFollowListDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f followListDo) Group(cols ...field.Expr) IFollowListDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f followListDo) Having(conds ...gen.Condition) IFollowListDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f followListDo) Limit(limit int) IFollowListDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f followListDo) Offset(offset int) IFollowListDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f followListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFollowListDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f followListDo) Unscoped() IFollowListDo {
	return f.withDO(f.DO.Unscoped())
}

func (f followListDo) Create(values ...*model.FollowList) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f followListDo) CreateInBatches(values []*model.FollowList, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f followListDo) Save(values ...*model.FollowList) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f followListDo) First() (*model.FollowList, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowList), nil
	}
}

func (f followListDo) Take() (*model.FollowList, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowList), nil
	}
}

func (f followListDo) Last() (*model.FollowList, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowList), nil
	}
}

func (f followListDo) Find() ([]*model.FollowList, error) {
	result, err := f.DO.Find()
	return result.([]*model.FollowList), err
}

func (f followListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FollowList, err error) {
	buf := make([]*model.FollowList, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f followListDo) FindInBatches(result *[]*model.FollowList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f followListDo) Attrs(attrs ...field.AssignExpr) IFollowListDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f followListDo) Assign(attrs ...field.AssignExpr) IFollowListDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f followListDo) Joins(fields ...field.RelationField) IFollowListDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f followListDo) Preload(fields ...field.RelationField) IFollowListDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f followListDo) FirstOrInit() (*model.FollowList, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowList), nil
	}
}

func (f followListDo) FirstOrCreate() (*model.FollowList, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowList), nil
	}
}

func (f followListDo) FindByPage(offset int, limit int) (result []*model.FollowList, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f followListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f followListDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f followListDo) Delete(models ...*model.FollowList) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *followListDo) withDO(do gen.Dao) *followListDo {
	f.DO = *do.(*gen.DO)
	return f
}
