package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc/resolver/discovery"
	consulAPI "github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
	"os"
	"strings"
	"time"
	utilCtx "util/context"
	"util/monitor"
	"vw_gateway/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	// 导入 kratos 的 dtm 驱动
	_ "github.com/dtm-labs/driver-kratos"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "vw_gateway"
	// Version is the version of the compiled software.
	Version string = "vw_gateway.v1"
	// flagconf is the config flag.
	flagconf string

	id = "vw_gateway"
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(rr),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// 配置jaeger链路追踪服务
	if err := monitor.SetTracerProvider(bc.Trace.Endpoint, Name); err != nil {
		panic(err)
	}

	// 注册 DTM 服务
	if err := registerDTM(bc.Registry, bc.Dtm); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(
		bc.Server,
		bc.Data,
		bc.Registry,
		bc.Jwt,
		bc.Email,
		bc.Trace,
		bc.Service,
		bc.Dtm,
		logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func registerDTM(consulConfig *conf.Registry, dtmConfig *conf.DTM) error {
	c := consulAPI.DefaultConfig()
	c.Address = consulConfig.Consul.Address
	c.Scheme = consulConfig.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		return err
	}
	r := consul.New(cli, consul.WithHealthCheck(true))
	resolver.Register(discovery.NewBuilder(r, discovery.WithInsecure(true)))
	ctx, cancel := utilCtx.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dtmService := &registry.ServiceInstance{
		ID:        "DTM_SERVER",
		Name:      "dtm",
		Metadata:  make(map[string]string),
		Endpoints: strings.Split(dtmConfig.DtmConfig.Endpoint, ","),
	}
	return r.Register(ctx, dtmService)
	return nil
}
