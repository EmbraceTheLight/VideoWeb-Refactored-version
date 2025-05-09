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

func newVideo(db *gorm.DB, opts ...gen.DOOption) video {
	_video := video{}

	_video.videoDo.UseDB(db, opts...)
	_video.videoDo.UseModel(&model.Video{})

	tableName := _video.videoDo.TableName()
	_video.ALL = field.NewAsterisk(tableName)
	_video.CreatedAt = field.NewTime(tableName, "created_at")
	_video.UpdatedAt = field.NewTime(tableName, "updated_at")
	_video.DeletedAt = field.NewField(tableName, "deleted_at")
	_video.Title = field.NewString(tableName, "title")
	_video.Description = field.NewString(tableName, "description")
	_video.Class = field.NewString(tableName, "class")
	_video.Hot = field.NewInt64(tableName, "hot")
	_video.Tags = field.NewString(tableName, "tags")
	_video.VideoPath = field.NewString(tableName, "video_path")
	_video.VideoID = field.NewInt64(tableName, "video_id")
	_video.PublisherID = field.NewInt64(tableName, "publisher_id")
	_video.PublisherName = field.NewString(tableName, "publisher_name")
	_video.Likes = field.NewInt64(tableName, "likes")
	_video.Shells = field.NewInt64(tableName, "shells")
	_video.CntBarrages = field.NewInt64(tableName, "cnt_barrages")
	_video.CntShares = field.NewInt64(tableName, "cnt_shares")
	_video.CntFavorited = field.NewInt64(tableName, "cnt_favorited")
	_video.CntViewed = field.NewInt64(tableName, "cnt_viewed")
	_video.CntComments = field.NewInt64(tableName, "cnt_comments")
	_video.Duration = field.NewString(tableName, "duration")
	_video.Size = field.NewInt64(tableName, "size")
	_video.CoverPath = field.NewString(tableName, "cover_path")
	_video.Version = field.NewField(tableName, "version")

	_video.fillFieldMap()

	return _video
}

type video struct {
	videoDo

	ALL           field.Asterisk
	CreatedAt     field.Time   // 创建时间
	UpdatedAt     field.Time   // 更新时间
	DeletedAt     field.Field  // 删除时间
	Title         field.String //  视频标题
	Description   field.String // 视频描述
	Class         field.String // 视频所属类别（以英文逗号 , 分隔多个类别）
	Hot           field.Int64  // 视频热度
	Tags          field.String // 视频标签，以英文逗号分隔
	VideoPath     field.String // 视频文件路径
	VideoID       field.Int64  // 视频ID
	PublisherID   field.Int64  // 上传者用户ID
	PublisherName field.String // 上传者用户名
	Likes         field.Int64  // 视频点赞数
	Shells        field.Int64  // 视频获得的贝壳数
	CntBarrages   field.Int64  // 视频弹幕数
	CntShares     field.Int64  // 视频分享数
	CntFavorited  field.Int64  // 视频收藏数
	CntViewed     field.Int64  // 视频观看数（点开就算看）
	CntComments   field.Int64
	Duration      field.String // 视频时长
	Size          field.Int64  // 视频文件大小
	CoverPath     field.String // 视频封面路径
	Version       field.Field  // 乐观锁版本号

	fieldMap map[string]field.Expr
}

func (v video) Table(newTableName string) *video {
	v.videoDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v video) As(alias string) *video {
	v.videoDo.DO = *(v.videoDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *video) updateTableName(table string) *video {
	v.ALL = field.NewAsterisk(table)
	v.CreatedAt = field.NewTime(table, "created_at")
	v.UpdatedAt = field.NewTime(table, "updated_at")
	v.DeletedAt = field.NewField(table, "deleted_at")
	v.Title = field.NewString(table, "title")
	v.Description = field.NewString(table, "description")
	v.Class = field.NewString(table, "class")
	v.Hot = field.NewInt64(table, "hot")
	v.Tags = field.NewString(table, "tags")
	v.VideoPath = field.NewString(table, "video_path")
	v.VideoID = field.NewInt64(table, "video_id")
	v.PublisherID = field.NewInt64(table, "publisher_id")
	v.PublisherName = field.NewString(table, "publisher_name")
	v.Likes = field.NewInt64(table, "likes")
	v.Shells = field.NewInt64(table, "shells")
	v.CntBarrages = field.NewInt64(table, "cnt_barrages")
	v.CntShares = field.NewInt64(table, "cnt_shares")
	v.CntFavorited = field.NewInt64(table, "cnt_favorited")
	v.CntViewed = field.NewInt64(table, "cnt_viewed")
	v.CntComments = field.NewInt64(table, "cnt_comments")
	v.Duration = field.NewString(table, "duration")
	v.Size = field.NewInt64(table, "size")
	v.CoverPath = field.NewString(table, "cover_path")
	v.Version = field.NewField(table, "version")

	v.fillFieldMap()

	return v
}

func (v *video) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *video) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 23)
	v.fieldMap["created_at"] = v.CreatedAt
	v.fieldMap["updated_at"] = v.UpdatedAt
	v.fieldMap["deleted_at"] = v.DeletedAt
	v.fieldMap["title"] = v.Title
	v.fieldMap["description"] = v.Description
	v.fieldMap["class"] = v.Class
	v.fieldMap["hot"] = v.Hot
	v.fieldMap["tags"] = v.Tags
	v.fieldMap["video_path"] = v.VideoPath
	v.fieldMap["video_id"] = v.VideoID
	v.fieldMap["publisher_id"] = v.PublisherID
	v.fieldMap["publisher_name"] = v.PublisherName
	v.fieldMap["likes"] = v.Likes
	v.fieldMap["shells"] = v.Shells
	v.fieldMap["cnt_barrages"] = v.CntBarrages
	v.fieldMap["cnt_shares"] = v.CntShares
	v.fieldMap["cnt_favorited"] = v.CntFavorited
	v.fieldMap["cnt_viewed"] = v.CntViewed
	v.fieldMap["cnt_comments"] = v.CntComments
	v.fieldMap["duration"] = v.Duration
	v.fieldMap["size"] = v.Size
	v.fieldMap["cover_path"] = v.CoverPath
	v.fieldMap["version"] = v.Version
}

func (v video) clone(db *gorm.DB) video {
	v.videoDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v video) replaceDB(db *gorm.DB) video {
	v.videoDo.ReplaceDB(db)
	return v
}

type videoDo struct{ gen.DO }

type IVideoDo interface {
	gen.SubQuery
	Debug() IVideoDo
	WithContext(ctx context.Context) IVideoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVideoDo
	WriteDB() IVideoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVideoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVideoDo
	Not(conds ...gen.Condition) IVideoDo
	Or(conds ...gen.Condition) IVideoDo
	Select(conds ...field.Expr) IVideoDo
	Where(conds ...gen.Condition) IVideoDo
	Order(conds ...field.Expr) IVideoDo
	Distinct(cols ...field.Expr) IVideoDo
	Omit(cols ...field.Expr) IVideoDo
	Join(table schema.Tabler, on ...field.Expr) IVideoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVideoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVideoDo
	Group(cols ...field.Expr) IVideoDo
	Having(conds ...gen.Condition) IVideoDo
	Limit(limit int) IVideoDo
	Offset(offset int) IVideoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVideoDo
	Unscoped() IVideoDo
	Create(values ...*model.Video) error
	CreateInBatches(values []*model.Video, batchSize int) error
	Save(values ...*model.Video) error
	First() (*model.Video, error)
	Take() (*model.Video, error)
	Last() (*model.Video, error)
	Find() ([]*model.Video, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Video, err error)
	FindInBatches(result *[]*model.Video, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Video) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVideoDo
	Assign(attrs ...field.AssignExpr) IVideoDo
	Joins(fields ...field.RelationField) IVideoDo
	Preload(fields ...field.RelationField) IVideoDo
	FirstOrInit() (*model.Video, error)
	FirstOrCreate() (*model.Video, error)
	FindByPage(offset int, limit int) (result []*model.Video, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVideoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v videoDo) Debug() IVideoDo {
	return v.withDO(v.DO.Debug())
}

func (v videoDo) WithContext(ctx context.Context) IVideoDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v videoDo) ReadDB() IVideoDo {
	return v.Clauses(dbresolver.Read)
}

func (v videoDo) WriteDB() IVideoDo {
	return v.Clauses(dbresolver.Write)
}

func (v videoDo) Session(config *gorm.Session) IVideoDo {
	return v.withDO(v.DO.Session(config))
}

func (v videoDo) Clauses(conds ...clause.Expression) IVideoDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v videoDo) Returning(value interface{}, columns ...string) IVideoDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v videoDo) Not(conds ...gen.Condition) IVideoDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v videoDo) Or(conds ...gen.Condition) IVideoDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v videoDo) Select(conds ...field.Expr) IVideoDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v videoDo) Where(conds ...gen.Condition) IVideoDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v videoDo) Order(conds ...field.Expr) IVideoDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v videoDo) Distinct(cols ...field.Expr) IVideoDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v videoDo) Omit(cols ...field.Expr) IVideoDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v videoDo) Join(table schema.Tabler, on ...field.Expr) IVideoDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v videoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVideoDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v videoDo) RightJoin(table schema.Tabler, on ...field.Expr) IVideoDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v videoDo) Group(cols ...field.Expr) IVideoDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v videoDo) Having(conds ...gen.Condition) IVideoDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v videoDo) Limit(limit int) IVideoDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v videoDo) Offset(offset int) IVideoDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v videoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVideoDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v videoDo) Unscoped() IVideoDo {
	return v.withDO(v.DO.Unscoped())
}

func (v videoDo) Create(values ...*model.Video) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v videoDo) CreateInBatches(values []*model.Video, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v videoDo) Save(values ...*model.Video) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v videoDo) First() (*model.Video, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Video), nil
	}
}

func (v videoDo) Take() (*model.Video, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Video), nil
	}
}

func (v videoDo) Last() (*model.Video, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Video), nil
	}
}

func (v videoDo) Find() ([]*model.Video, error) {
	result, err := v.DO.Find()
	return result.([]*model.Video), err
}

func (v videoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Video, err error) {
	buf := make([]*model.Video, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v videoDo) FindInBatches(result *[]*model.Video, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v videoDo) Attrs(attrs ...field.AssignExpr) IVideoDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v videoDo) Assign(attrs ...field.AssignExpr) IVideoDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v videoDo) Joins(fields ...field.RelationField) IVideoDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v videoDo) Preload(fields ...field.RelationField) IVideoDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v videoDo) FirstOrInit() (*model.Video, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Video), nil
	}
}

func (v videoDo) FirstOrCreate() (*model.Video, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Video), nil
	}
}

func (v videoDo) FindByPage(offset int, limit int) (result []*model.Video, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v videoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v videoDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v videoDo) Delete(models ...*model.Video) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *videoDo) withDO(do gen.Dao) *videoDo {
	v.DO = *do.(*gen.DO)
	return v
}
