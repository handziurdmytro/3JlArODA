package middleware

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
)

func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	slog.Info("Received gRPC request", slog.String("method", info.FullMethod))

	resp, err := handler(ctx, req)

	if err != nil {
		slog.Error("gRPC request failed",
			slog.String("method", info.FullMethod),
			slog.String("error", err.Error()),
		)
	} else {
		slog.Info("gRPC request completed", slog.String("method", info.FullMethod))
	}

	return resp, err
}
