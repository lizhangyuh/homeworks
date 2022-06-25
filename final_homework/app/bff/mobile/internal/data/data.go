package data

import (
	"context"
	"time"

	postV1 "nvm.com/mrc-api-go/api/post/service/v1"
	userV1 "nvm.com/mrc-api-go/api/user/service/v1"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewUserServiceClient, NewRegistrar, NewDiscovery)

// Data .
type Data struct {
	log *log.Helper
	uc  userV1.UserClient // 用户服务的客户端
	pc  postV1.PostClient // 用户服务的客户端
}

// NewData .
func NewData(c *conf.Data, uc userV1.UserClient, logger log.Logger) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, uc: uc}, nil
}

// NewUserServiceClient 链接用户服务
func NewUserServiceClient(auth *conf.Auth, sr *conf.Service, rr registry.Discovery, tp *tracesdk.TracerProvider) userV1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.User.Endpoint), // consul
		grpc.WithDiscovery(rr),              // consul
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
			jwt.Client(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(auth.ServiceKey), nil
			}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)),
		),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
	c := userV1.NewUserClient(conn)
	return c
}

// NewPostServiceClient 链接文章服务
func NewPostUserServiceClient(auth *conf.Auth, sr *conf.Service, rr registry.Discovery, tp *tracesdk.TracerProvider) postV1.PostClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.User.Endpoint), // consul
		grpc.WithDiscovery(rr),              // consul
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
			jwt.Client(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(auth.ServiceKey), nil
			}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)),
		),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
	c := postV1.NewPostClient(conn)
	return c
}

// NewRegistrar add consul
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
