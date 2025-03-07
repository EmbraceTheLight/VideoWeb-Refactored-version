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

func newUserFavorite(db *gorm.DB, opts ...gen.DOOption) userFavorite {
	_userFavorite := userFavorite{}

	_userFavorite.userFavoriteDo.UseDB(db, opts...)
	_userFavorite.userFavoriteDo.UseModel(&model.UserFavorite{})

	tableName := _userFavorite.userFavoriteDo.TableName()
	_userFavorite.ALL = field.NewAsterisk(tableName)
	_userFavorite.FavoriteID = field.NewInt64(tableName, "favorite_id")
	_userFavorite.UserID = field.NewInt64(tableName, "user_id")
	_userFavorite.FavoriteName = field.NewString(tableName, "favorite_name")
	_userFavorite.Description = field.NewString(tableName, "description")
	_userFavorite.IsPrivate = field.NewInt64(tableName, "is_private")

	_userFavorite.fillFieldMap()

	return _userFavorite
}

// userFavorite 用户收藏夹表
type userFavorite struct {
	userFavoriteDo

	ALL          field.Asterisk
	FavoriteID   field.Int64
	UserID       field.Int64
	FavoriteName field.String
	Description  field.String
	IsPrivate    field.Int64 // 表示该收藏夹是否私密，1表示公开，-1表示私密

	fieldMap map[string]field.Expr
}

func (u userFavorite) Table(newTableName string) *userFavorite {
	u.userFavoriteDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userFavorite) As(alias string) *userFavorite {
	u.userFavoriteDo.DO = *(u.userFavoriteDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userFavorite) updateTableName(table string) *userFavorite {
	u.ALL = field.NewAsterisk(table)
	u.FavoriteID = field.NewInt64(table, "favorite_id")
	u.UserID = field.NewInt64(table, "user_id")
	u.FavoriteName = field.NewString(table, "favorite_name")
	u.Description = field.NewString(table, "description")
	u.IsPrivate = field.NewInt64(table, "is_private")

	u.fillFieldMap()

	return u
}

func (u *userFavorite) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userFavorite) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 5)
	u.fieldMap["favorite_id"] = u.FavoriteID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["favorite_name"] = u.FavoriteName
	u.fieldMap["description"] = u.Description
	u.fieldMap["is_private"] = u.IsPrivate
}

func (u userFavorite) clone(db *gorm.DB) userFavorite {
	u.userFavoriteDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userFavorite) replaceDB(db *gorm.DB) userFavorite {
	u.userFavoriteDo.ReplaceDB(db)
	return u
}

type userFavoriteDo struct{ gen.DO }

type IUserFavoriteDo interface {
	gen.SubQuery
	Debug() IUserFavoriteDo
	WithContext(ctx context.Context) IUserFavoriteDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserFavoriteDo
	WriteDB() IUserFavoriteDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserFavoriteDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserFavoriteDo
	Not(conds ...gen.Condition) IUserFavoriteDo
	Or(conds ...gen.Condition) IUserFavoriteDo
	Select(conds ...field.Expr) IUserFavoriteDo
	Where(conds ...gen.Condition) IUserFavoriteDo
	Order(conds ...field.Expr) IUserFavoriteDo
	Distinct(cols ...field.Expr) IUserFavoriteDo
	Omit(cols ...field.Expr) IUserFavoriteDo
	Join(table schema.Tabler, on ...field.Expr) IUserFavoriteDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserFavoriteDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserFavoriteDo
	Group(cols ...field.Expr) IUserFavoriteDo
	Having(conds ...gen.Condition) IUserFavoriteDo
	Limit(limit int) IUserFavoriteDo
	Offset(offset int) IUserFavoriteDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserFavoriteDo
	Unscoped() IUserFavoriteDo
	Create(values ...*model.UserFavorite) error
	CreateInBatches(values []*model.UserFavorite, batchSize int) error
	Save(values ...*model.UserFavorite) error
	First() (*model.UserFavorite, error)
	Take() (*model.UserFavorite, error)
	Last() (*model.UserFavorite, error)
	Find() ([]*model.UserFavorite, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserFavorite, err error)
	FindInBatches(result *[]*model.UserFavorite, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserFavorite) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserFavoriteDo
	Assign(attrs ...field.AssignExpr) IUserFavoriteDo
	Joins(fields ...field.RelationField) IUserFavoriteDo
	Preload(fields ...field.RelationField) IUserFavoriteDo
	FirstOrInit() (*model.UserFavorite, error)
	FirstOrCreate() (*model.UserFavorite, error)
	FindByPage(offset int, limit int) (result []*model.UserFavorite, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserFavoriteDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userFavoriteDo) Debug() IUserFavoriteDo {
	return u.withDO(u.DO.Debug())
}

func (u userFavoriteDo) WithContext(ctx context.Context) IUserFavoriteDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userFavoriteDo) ReadDB() IUserFavoriteDo {
	return u.Clauses(dbresolver.Read)
}

func (u userFavoriteDo) WriteDB() IUserFavoriteDo {
	return u.Clauses(dbresolver.Write)
}

func (u userFavoriteDo) Session(config *gorm.Session) IUserFavoriteDo {
	return u.withDO(u.DO.Session(config))
}

func (u userFavoriteDo) Clauses(conds ...clause.Expression) IUserFavoriteDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userFavoriteDo) Returning(value interface{}, columns ...string) IUserFavoriteDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userFavoriteDo) Not(conds ...gen.Condition) IUserFavoriteDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userFavoriteDo) Or(conds ...gen.Condition) IUserFavoriteDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userFavoriteDo) Select(conds ...field.Expr) IUserFavoriteDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userFavoriteDo) Where(conds ...gen.Condition) IUserFavoriteDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userFavoriteDo) Order(conds ...field.Expr) IUserFavoriteDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userFavoriteDo) Distinct(cols ...field.Expr) IUserFavoriteDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userFavoriteDo) Omit(cols ...field.Expr) IUserFavoriteDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userFavoriteDo) Join(table schema.Tabler, on ...field.Expr) IUserFavoriteDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userFavoriteDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserFavoriteDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userFavoriteDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserFavoriteDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userFavoriteDo) Group(cols ...field.Expr) IUserFavoriteDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userFavoriteDo) Having(conds ...gen.Condition) IUserFavoriteDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userFavoriteDo) Limit(limit int) IUserFavoriteDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userFavoriteDo) Offset(offset int) IUserFavoriteDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userFavoriteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserFavoriteDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userFavoriteDo) Unscoped() IUserFavoriteDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userFavoriteDo) Create(values ...*model.UserFavorite) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userFavoriteDo) CreateInBatches(values []*model.UserFavorite, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userFavoriteDo) Save(values ...*model.UserFavorite) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userFavoriteDo) First() (*model.UserFavorite, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorite), nil
	}
}

func (u userFavoriteDo) Take() (*model.UserFavorite, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorite), nil
	}
}

func (u userFavoriteDo) Last() (*model.UserFavorite, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorite), nil
	}
}

func (u userFavoriteDo) Find() ([]*model.UserFavorite, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserFavorite), err
}

func (u userFavoriteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserFavorite, err error) {
	buf := make([]*model.UserFavorite, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userFavoriteDo) FindInBatches(result *[]*model.UserFavorite, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userFavoriteDo) Attrs(attrs ...field.AssignExpr) IUserFavoriteDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userFavoriteDo) Assign(attrs ...field.AssignExpr) IUserFavoriteDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userFavoriteDo) Joins(fields ...field.RelationField) IUserFavoriteDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userFavoriteDo) Preload(fields ...field.RelationField) IUserFavoriteDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userFavoriteDo) FirstOrInit() (*model.UserFavorite, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorite), nil
	}
}

func (u userFavoriteDo) FirstOrCreate() (*model.UserFavorite, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorite), nil
	}
}

func (u userFavoriteDo) FindByPage(offset int, limit int) (result []*model.UserFavorite, count int64, err error) {
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

func (u userFavoriteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userFavoriteDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userFavoriteDo) Delete(models ...*model.UserFavorite) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userFavoriteDo) withDO(do gen.Dao) *userFavoriteDo {
	u.DO = *do.(*gen.DO)
	return u
}
