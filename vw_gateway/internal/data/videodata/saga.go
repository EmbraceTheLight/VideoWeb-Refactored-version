package videodata

import (
	"context"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"google.golang.org/protobuf/proto"
	"util/getid"
)

type sagaMember struct {
	method           string
	compensateMethod string
	payload          proto.Message
}

func newSaga(ctx context.Context, dtmServerAddr string) (string, *dtmgrpc.SagaGrpc) {
	gid := getid.GetUUID()
	saga := dtmgrpc.NewSagaGrpcWithContext(ctx, dtmServerAddr, gid)
	return gid, saga
}
func wrapSagaAdd(saga *dtmgrpc.SagaGrpc, member ...*sagaMember) *dtmgrpc.SagaGrpc {
	for _, s := range member {
		saga.Add(s.method, s.compensateMethod, s.payload)
	}
	saga.WaitResult = true
	return saga
}

func wrapSaga(ctx context.Context, dtmServerAddr string, member ...*sagaMember) (string, *dtmgrpc.SagaGrpc) {
	gid, saga := newSaga(ctx, dtmServerAddr)
	wrapSagaAdd(saga, member...)
	return gid, saga
}
