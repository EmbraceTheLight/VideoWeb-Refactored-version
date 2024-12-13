package snowflake

import (
	"github.com/yitter/idgenerator-go/idgen"
)

func init() {
	var options = idgen.NewIdGeneratorOptions(1)
	options.BaseTime = baseTime
	idgen.SetIdGenerator(options)
	//TODO:将 雪花算法组件初始化成功 记录日志
}

// GetID generates ID by snowflake algorithm
func GetID() int64 {
	return idgen.NextId()
}
