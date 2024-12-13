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

func newUserFollow(db *gorm.DB, opts ...gen.DOOption) userFollow {
	_userFollow := userFollow{}

	_userFollow.userFollowDo.UseDB(db, opts...)
	_userFollow.userFollowDo.UseModel(&model.UserFollow{})

	tableName := _userFollow.userFollowDo.TableName()
	_userFollow.ALL = field.NewAsterisk(tableName)
	_userFollow.FollowlistID = field.NewInt64(tableName, "followlist_id")
	_userFollow.FollowUserID = field.NewInt64(tableName, "follow_user_id")
	_userFollow.UserID = field.NewInt64(tableName, "user_id")
	_userFollow.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userFollow.DeletedAt = field.NewField(tableName, "deleted_at")
	_userFollow.CreatedAt = field.NewTime(tableName, "created_at")

	_userFollow.fillFieldMap()

	return _userFollow
}

// userFollow 存放用户关注的用户的id
type userFollow struct {
	userFollowDo

	ALL          field.Asterisk
	FollowlistID field.Int64
	FollowUserID field.Int64 // 被关注的用户id
	UserID       field.Int64
	UpdatedAt    field.Time
	DeletedAt    field.Field
	CreatedAt    field.Time

	fieldMap map[string]field.Expr
}

func (u userFollow) Table(newTableName string) *userFollow {
	u.userFollowDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userFollow) As(alias string) *userFollow {
	u.userFollowDo.DO = *(u.userFollowDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userFollow) updateTableName(table string) *userFollow {
	u.ALL = field.NewAsterisk(table)
	u.FollowlistID = field.NewInt64(table, "followlist_id")
	u.FollowUserID = field.NewInt64(table, "follow_user_id")
	u.UserID = field.NewInt64(table, "user_id")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")
	u.CreatedAt = field.NewTime(table, "created_at")

	u.fillFieldMap()

	return u
}

func (u *userFollow) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userFollow) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["followlist_id"] = u.FollowlistID
	u.fieldMap["follow_user_id"] = u.FollowUserID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
	u.fieldMap["created_at"] = u.CreatedAt
}

func (u userFollow) clone(db *gorm.DB) userFollow {
	u.userFollowDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userFollow) replaceDB(db *gorm.DB) userFollow {
	u.userFollowDo.ReplaceDB(db)
	return u
}

type userFollowDo struct{ gen.DO }

type IUserFollowDo interface {
	gen.SubQuery
	Debug() IUserFollowDo
	WithContext(ctx context.Context) IUserFollowDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserFollowDo
	WriteDB() IUserFollowDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserFollowDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserFollowDo
	Not(conds ...gen.Condition) IUserFollowDo
	Or(conds ...gen.Condition) IUserFollowDo
	Select(conds ...field.Expr) IUserFollowDo
	Where(conds ...gen.Condition) IUserFollowDo
	Order(conds ...field.Expr) IUserFollowDo
	Distinct(cols ...field.Expr) IUserFollowDo
	Omit(cols ...field.Expr) IUserFollowDo
	Join(table schema.Tabler, on ...field.Expr) IUserFollowDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserFollowDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserFollowDo
	Group(cols ...field.Expr) IUserFollowDo
	Having(conds ...gen.Condition) IUserFollowDo
	Limit(limit int) IUserFollowDo
	Offset(offset int) IUserFollowDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserFollowDo
	Unscoped() IUserFollowDo
	Create(values ...*model.UserFollow) error
	CreateInBatches(values []*model.UserFollow, batchSize int) error
	Save(values ...*model.UserFollow) error
	First() (*model.UserFollow, error)
	Take() (*model.UserFollow, error)
	Last() (*model.UserFollow, error)
	Find() ([]*model.UserFollow, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserFollow, err error)
	FindInBatches(result *[]*model.UserFollow, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserFollow) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserFollowDo
	Assign(attrs ...field.AssignExpr) IUserFollowDo
	Joins(fields ...field.RelationField) IUserFollowDo
	Preload(fields ...field.RelationField) IUserFollowDo
	FirstOrInit() (*model.UserFollow, error)
	FirstOrCreate() (*model.UserFollow, error)
	FindByPage(offset int, limit int) (result []*model.UserFollow, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserFollowDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userFollowDo) Debug() IUserFollowDo {
	return u.withDO(u.DO.Debug())
}

func (u userFollowDo) WithContext(ctx context.Context) IUserFollowDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userFollowDo) ReadDB() IUserFollowDo {
	return u.Clauses(dbresolver.Read)
}

func (u userFollowDo) WriteDB() IUserFollowDo {
	return u.Clauses(dbresolver.Write)
}

func (u userFollowDo) Session(config *gorm.Session) IUserFollowDo {
	return u.withDO(u.DO.Session(config))
}

func (u userFollowDo) Clauses(conds ...clause.Expression) IUserFollowDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userFollowDo) Returning(value interface{}, columns ...string) IUserFollowDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userFollowDo) Not(conds ...gen.Condition) IUserFollowDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userFollowDo) Or(conds ...gen.Condition) IUserFollowDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userFollowDo) Select(conds ...field.Expr) IUserFollowDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userFollowDo) Where(conds ...gen.Condition) IUserFollowDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userFollowDo) Order(conds ...field.Expr) IUserFollowDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userFollowDo) Distinct(cols ...field.Expr) IUserFollowDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userFollowDo) Omit(cols ...field.Expr) IUserFollowDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userFollowDo) Join(table schema.Tabler, on ...field.Expr) IUserFollowDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userFollowDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserFollowDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userFollowDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserFollowDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userFollowDo) Group(cols ...field.Expr) IUserFollowDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userFollowDo) Having(conds ...gen.Condition) IUserFollowDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userFollowDo) Limit(limit int) IUserFollowDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userFollowDo) Offset(offset int) IUserFollowDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userFollowDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserFollowDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userFollowDo) Unscoped() IUserFollowDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userFollowDo) Create(values ...*model.UserFollow) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userFollowDo) CreateInBatches(values []*model.UserFollow, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userFollowDo) Save(values ...*model.UserFollow) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userFollowDo) First() (*model.UserFollow, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFollow), nil
	}
}

func (u userFollowDo) Take() (*model.UserFollow, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFollow), nil
	}
}

func (u userFollowDo) Last() (*model.UserFollow, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFollow), nil
	}
}

func (u userFollowDo) Find() ([]*model.UserFollow, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserFollow), err
}

func (u userFollowDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserFollow, err error) {
	buf := make([]*model.UserFollow, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userFollowDo) FindInBatches(result *[]*model.UserFollow, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userFollowDo) Attrs(attrs ...field.AssignExpr) IUserFollowDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userFollowDo) Assign(attrs ...field.AssignExpr) IUserFollowDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userFollowDo) Joins(fields ...field.RelationField) IUserFollowDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userFollowDo) Preload(fields ...field.RelationField) IUserFollowDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userFollowDo) FirstOrInit() (*model.UserFollow, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFollow), nil
	}
}

func (u userFollowDo) FirstOrCreate() (*model.UserFollow, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFollow), nil
	}
}

func (u userFollowDo) FindByPage(offset int, limit int) (result []*model.UserFollow, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userFollowDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userFollowDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userFollowDo) Delete(models ...*model.UserFollow) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userFollowDo) withDO(do gen.Dao) *userFollowDo {
	u.DO = *do.(*gen.DO)
	return u
}
