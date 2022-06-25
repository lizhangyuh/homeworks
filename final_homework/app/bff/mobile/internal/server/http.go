package server

import (
	"context"

	v1 "nvm.com/mrc-api-go/api/bff/mobile/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/conf"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/service"

	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/bff.mobile.v1.Mobile/SignInByPwd"] = struct{}{}
	whiteList["/bff.mobile.v1.Mobile/SignUp"] = struct{}{}

	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}

		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, auth *conf.Auth, tp *tracesdk.TracerProvider, cs *service.MoibleService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(tracing.WithTracerProvider(tp)),
			logging.Server(logger),
			selector.Server(jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(auth.ApiKey), nil
			}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256), jwt.WithClaims(func() jwtv4.Claims {
				return jwtv4.MapClaims{}
			})),
			).Match(NewWhiteListMatcher()).Build(),
		),
		http.Filter(handlers.CORS( // 浏览器跨域
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	// swagger api

	openAPIHandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIHandler)

	v1.RegisterMobileHTTPServer(srv, cs)
	return srv
}
