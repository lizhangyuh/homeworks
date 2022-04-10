package service

import (
	"context"

	pb "app/api/app/v1"

	"go.opentelemetry.io/otel"
)

func (s *AppService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	//  add trace
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "regist user")
	span.SpanContext()
	defer span.End()

	return s.uc.CreateUser(ctx, req)
}

func (s *AppService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	//  add trace
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "auth user by phone")
	span.SpanContext()
	defer span.End()

	return s.uc.SignInByPassword(ctx, req)
}
