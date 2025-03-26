package getid

import (
	"github.com/yitter/idgenerator-go/idgen"
)

func init() {
	var options = idgen.NewIdGeneratorOptions(1)
	options.BaseTime = baseTime
	idgen.SetIdGenerator(options)
}

// GetID generates ID by snowflake algorithm
func GetID() int64 {
	return idgen.NextId()
}
