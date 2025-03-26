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

	"vw_video/internal/data/dal/model"
)

func newBarrage(db *gorm.DB, opts ...gen.DOOption) barrage {
	_barrage := barrage{}

	_barrage.barrageDo.UseDB(db, opts...)
	_barrage.barrageDo.UseModel(&model.Barrage{})

	tableName := _barrage.barrageDo.TableName()
	_barrage.ALL = field.NewAsterisk(tableName)
	_barrage.CreatedAt = field.NewTime(tableName, "created_at")
	_barrage.UpdatedAt = field.NewTime(tableName, "updated_at")
	_barrage.DeletedAt = field.NewField(tableName, "deleted_at")
	_barrage.UserID = field.NewInt64(tableName, "user_id")
	_barrage.VideoID = field.NewInt64(tableName, "video_id")
	_barrage.Hour = field.NewString(tableName, "hour")
	_barrage.Minute = field.NewString(tableName, "minute")
	_barrage.Second = field.NewString(tableName, "second")
	_barrage.Content = field.NewString(tableName, "content")
	_barrage.Color = field.NewString(tableName, "color")
	_barrage.Likes = field.NewInt64(tableName, "likes")

	_barrage.fillFieldMap()

	return _barrage
}

type barrage struct {
	barrageDo

	ALL       field.Asterisk
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间
	DeletedAt field.Field  // 删除时间
	UserID    field.Int64  // 用户id
	VideoID   field.Int64  // 视频id
	Hour      field.String // 弹幕出现时间--小时
	Minute    field.String // 弹幕出现时间--分钟
	Second    field.String // 弹幕出现时间--秒
	Content   field.String // 弹幕内容
	Color     field.String // 弹幕颜色，使用十六进制表示
	Likes     field.Int64  // 弹幕获赞数

	fieldMap map[string]field.Expr
}

func (b barrage) Table(newTableName string) *barrage {
	b.barrageDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b barrage) As(alias string) *barrage {
	b.barrageDo.DO = *(b.barrageDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *barrage) updateTableName(table string) *barrage {
	b.ALL = field.NewAsterisk(table)
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")
	b.DeletedAt = field.NewField(table, "deleted_at")
	b.UserID = field.NewInt64(table, "user_id")
	b.VideoID = field.NewInt64(table, "video_id")
	b.Hour = field.NewString(table, "hour")
	b.Minute = field.NewString(table, "minute")
	b.Second = field.NewString(table, "second")
	b.Content = field.NewString(table, "content")
	b.Color = field.NewString(table, "color")
	b.Likes = field.NewInt64(table, "likes")

	b.fillFieldMap()

	return b
}

func (b *barrage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *barrage) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 11)
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
	b.fieldMap["deleted_at"] = b.DeletedAt
	b.fieldMap["user_id"] = b.UserID
	b.fieldMap["video_id"] = b.VideoID
	b.fieldMap["hour"] = b.Hour
	b.fieldMap["minute"] = b.Minute
	b.fieldMap["second"] = b.Second
	b.fieldMap["content"] = b.Content
	b.fieldMap["color"] = b.Color
	b.fieldMap["likes"] = b.Likes
}

func (b barrage) clone(db *gorm.DB) barrage {
	b.barrageDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b barrage) replaceDB(db *gorm.DB) barrage {
	b.barrageDo.ReplaceDB(db)
	return b
}

type barrageDo struct{ gen.DO }

type IBarrageDo interface {
	gen.SubQuery
	Debug() IBarrageDo
	WithContext(ctx context.Context) IBarrageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBarrageDo
	WriteDB() IBarrageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBarrageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBarrageDo
	Not(conds ...gen.Condition) IBarrageDo
	Or(conds ...gen.Condition) IBarrageDo
	Select(conds ...field.Expr) IBarrageDo
	Where(conds ...gen.Condition) IBarrageDo
	Order(conds ...field.Expr) IBarrageDo
	Distinct(cols ...field.Expr) IBarrageDo
	Omit(cols ...field.Expr) IBarrageDo
	Join(table schema.Tabler, on ...field.Expr) IBarrageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBarrageDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBarrageDo
	Group(cols ...field.Expr) IBarrageDo
	Having(conds ...gen.Condition) IBarrageDo
	Limit(limit int) IBarrageDo
	Offset(offset int) IBarrageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBarrageDo
	Unscoped() IBarrageDo
	Create(values ...*model.Barrage) error
	CreateInBatches(values []*model.Barrage, batchSize int) error
	Save(values ...*model.Barrage) error
	First() (*model.Barrage, error)
	Take() (*model.Barrage, error)
	Last() (*model.Barrage, error)
	Find() ([]*model.Barrage, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Barrage, err error)
	FindInBatches(result *[]*model.Barrage, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Barrage) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBarrageDo
	Assign(attrs ...field.AssignExpr) IBarrageDo
	Joins(fields ...field.RelationField) IBarrageDo
	Preload(fields ...field.RelationField) IBarrageDo
	FirstOrInit() (*model.Barrage, error)
	FirstOrCreate() (*model.Barrage, error)
	FindByPage(offset int, limit int) (result []*model.Barrage, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBarrageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b barrageDo) Debug() IBarrageDo {
	return b.withDO(b.DO.Debug())
}

func (b barrageDo) WithContext(ctx context.Context) IBarrageDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b barrageDo) ReadDB() IBarrageDo {
	return b.Clauses(dbresolver.Read)
}

func (b barrageDo) WriteDB() IBarrageDo {
	return b.Clauses(dbresolver.Write)
}

func (b barrageDo) Session(config *gorm.Session) IBarrageDo {
	return b.withDO(b.DO.Session(config))
}

func (b barrageDo) Clauses(conds ...clause.Expression) IBarrageDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b barrageDo) Returning(value interface{}, columns ...string) IBarrageDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b barrageDo) Not(conds ...gen.Condition) IBarrageDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b barrageDo) Or(conds ...gen.Condition) IBarrageDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b barrageDo) Select(conds ...field.Expr) IBarrageDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b barrageDo) Where(conds ...gen.Condition) IBarrageDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b barrageDo) Order(conds ...field.Expr) IBarrageDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b barrageDo) Distinct(cols ...field.Expr) IBarrageDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b barrageDo) Omit(cols ...field.Expr) IBarrageDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b barrageDo) Join(table schema.Tabler, on ...field.Expr) IBarrageDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b barrageDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBarrageDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b barrageDo) RightJoin(table schema.Tabler, on ...field.Expr) IBarrageDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b barrageDo) Group(cols ...field.Expr) IBarrageDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b barrageDo) Having(conds ...gen.Condition) IBarrageDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b barrageDo) Limit(limit int) IBarrageDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b barrageDo) Offset(offset int) IBarrageDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b barrageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBarrageDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b barrageDo) Unscoped() IBarrageDo {
	return b.withDO(b.DO.Unscoped())
}

func (b barrageDo) Create(values ...*model.Barrage) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b barrageDo) CreateInBatches(values []*model.Barrage, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b barrageDo) Save(values ...*model.Barrage) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b barrageDo) First() (*model.Barrage, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Barrage), nil
	}
}

func (b barrageDo) Take() (*model.Barrage, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Barrage), nil
	}
}

func (b barrageDo) Last() (*model.Barrage, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Barrage), nil
	}
}

func (b barrageDo) Find() ([]*model.Barrage, error) {
	result, err := b.DO.Find()
	return result.([]*model.Barrage), err
}

func (b barrageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Barrage, err error) {
	buf := make([]*model.Barrage, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b barrageDo) FindInBatches(result *[]*model.Barrage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b barrageDo) Attrs(attrs ...field.AssignExpr) IBarrageDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b barrageDo) Assign(attrs ...field.AssignExpr) IBarrageDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b barrageDo) Joins(fields ...field.RelationField) IBarrageDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b barrageDo) Preload(fields ...field.RelationField) IBarrageDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b barrageDo) FirstOrInit() (*model.Barrage, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Barrage), nil
	}
}

func (b barrageDo) FirstOrCreate() (*model.Barrage, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Barrage), nil
	}
}

func (b barrageDo) FindByPage(offset int, limit int) (result []*model.Barrage, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b barrageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b barrageDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b barrageDo) Delete(models ...*model.Barrage) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *barrageDo) withDO(do gen.Dao) *barrageDo {
	b.DO = *do.(*gen.DO)
	return b
}
