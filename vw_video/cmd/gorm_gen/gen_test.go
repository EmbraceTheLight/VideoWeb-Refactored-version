package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"gorm.io/gen"
	"gorm.io/gorm"
	"sync"
	"testing"
	"time"
	"vw_video/internal/data/dal/query"

	"vw_video/internal/data"
	"vw_video/internal/data/dal/model"
)

func TestGen(t *testing.T) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../internal/data/dal1/query",

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		//FieldNullable: true,

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false,

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false,

		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: true,

		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true,
	})
	//设置目标db
	g.UseDB(data.GetDB())

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
	}

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	g.WithDataTypeMap(dataMap)

	user := g.GenerateModel("userbiz", gen.FieldType("version", "optimisticlock.Version"))
	g.ApplyBasic(user)

	g.Execute()

}

// 测试生成的查询代码
func TestGetAndInsertUser(t *testing.T) {
	require.NotNil(t, query.Video)
	user, err := query.Video.Last()

	var start int64
	if errors.Is(err, gorm.ErrRecordNotFound) {
		start = 1
	} else {
		start = user.UserID + 1
	}
	for i := start; i < start+3; i++ {
		err = query.User.WithContext(context.Background()).Create(&model.User{
			UserID:   i,
			Username: fmt.Sprintf("test %d", i),
			Password: fmt.Sprintf("test %d: %s", i, "123456"),
		})
		require.NoError(t, err)
	}

}

func TestOptimisticLock(t *testing.T) {
	t.Parallel()
	query.SetDefault(data.GetDB())
	require.NotNil(t, query.User)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		_, err := updateShellsUseGen(1, 100, 0)
		if err != nil {
			t.Error(err)
		}
	}()

	go func() {
		defer wg.Done()
		info, err := updateShellsUseGen(1, 200, 1000)
		if info.RowsAffected == 0 {
			fmt.Println("optimistic lock is working!")
		}
		if err != nil {
			t.Error(err)
		}
	}()
	wg.Wait()

}

// 测试gen结合乐观锁的平均消耗
func BenchmarkTimeDiffWithGen(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		updateShellsUseGen(1, 100, 0)
	}
}

// 测试gen结合乐观锁的平均消耗
func BenchmarkTimeDiffWithGorm(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		updateShellsUseCommonGorm(1, 100, 0)
	}
}

// 测试gorm乐观锁:使用gen生成的查询代码
func updateShellsUseGen(id int64, value int64, sleep int) (info gen.ResultInfo, err error) {
	tx := query.Q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			//回滚事务
			_ = tx.Rollback()
		}
	}()

	do := tx.User.WithContext(context.Background())
	user, err := do.Where(query.User.UserID.Eq(id)).First()
	if err != nil {
		return info, err
	}
	//fmt.Println("====================>", userbiz.Shells)

	//模拟并发更新，事务存在延迟的情况
	if sleep > 0 {
		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}
	do.ReplaceDB(do.UnderlyingDB().Model(user)) //指定底层gorm.DB的model为user，这样才能使用乐观锁插件
	//info, err = do.Debug().Where(query.User.UserID.Eq(id)).Update(tx.User.Shells, tx.User.Shells.Add(value)) //这行代码就不能使用乐观锁插件，尚不清楚原因
	//info, err = do.Debug().Where(query.User.UserID.Eq(id)).Updates(&model.User{
	info, err = do.Where(query.User.UserID.Eq(id)).Updates(&model.User{
		Shells: user.Shells + value,
	})
	if err != nil {
		fmt.Println(err)
		return info, err
	}

	err = tx.Commit()
	return info, err
}

// 测试gorm乐观锁:使用gorm原生方法
func updateShellsUseCommonGorm(id int64, value int64, sleep int) (info gen.ResultInfo, err error) {
	db := data.GetDB()
	user := model.User{}
	res := db.Where(query.User.UserID.Eq(id)).First(&user)
	if res.Error != nil {
		return info, res.Error
	}
	tx := db.Begin()
	defer func() {
		if recover() != nil || err != nil {
			//回滚事务
			_ = tx.Rollback()
		}
	}()
	if sleep > 0 {
		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}
	//result := tx.Model(&userbiz).Debug().Update("shells", userbiz.Shells+value) //使用乐观锁的关键！它指定了具体的某条user记录
	result := tx.Model(&user).Update("shells", user.Shells+value) //使用乐观锁的关键！它指定了具体的某条user记录
	info.RowsAffected = result.RowsAffected
	info.Error = result.Error
	if result.Error != nil {
		fmt.Println(result.Error)
		return info, result.Error
	}

	tx.Commit()
	return info, err
}

func TestA(t *testing.T) {

}
