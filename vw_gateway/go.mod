module vw_gateway

go 1.23.3

require (
	github.com/gin-contrib/cors v1.7.4
	github.com/gin-gonic/gin v1.10.0
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/go-kratos/gin v0.1.0
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20250210091908-15168b5a1b7d
	github.com/go-kratos/kratos/v2 v2.8.3
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/wire v0.6.0
	github.com/gorilla/handlers v1.5.2
	github.com/hashicorp/consul/api v1.26.1
	github.com/mojocn/base64Captcha v1.3.8
	github.com/redis/go-redis/v9 v9.7.3
	github.com/spewerspew/spew v0.0.0-20230513223542-89b69fbbe2bd
	github.com/stretchr/testify v1.10.0
	go.uber.org/automaxprocs v1.6.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250224174004-546df14abb99
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.6
	util v0.0.0-00010101000000-000000000000
	vw_user v0.0.0-00010101000000-000000000000
	vw_video v0.0.0-00010101000000-000000000000
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/bytedance/sonic v1.13.2 // indirect
	github.com/bytedance/sonic/loader v0.2.4 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.5 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/gin-contrib/sse v1.0.0 // indirect
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.25.0 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.10 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	github.com/yitter/idgenerator-go v1.3.3 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	golang.org/x/arch v0.15.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/exp v0.0.0-20230817173708-d852ddb80c63 // indirect
	golang.org/x/image v0.23.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250219182151-9fdb1cabc7b2 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	common => ../util/common_proto
	util => ../util
	vw_user => ../vw_user
	vw_video => ../vw_video
)
