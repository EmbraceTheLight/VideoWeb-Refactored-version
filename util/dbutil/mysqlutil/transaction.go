package mysqlutil

// This Interface-Definition file is used by grpc service, like vw_user,vw_video and so on.
import "context"

type Transaction interface {
	WithTx(context.Context, func(context.Context) error) error
}
