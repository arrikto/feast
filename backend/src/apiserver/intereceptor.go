package main

import (
	"context"

	util "github.com/feast-dev/feast/backend/src/utils"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

// apiServerInterceptor implements UnaryServerInterceptor that provides the common wrapping logic
// to be executed before and after all API handler calls, e.g. Logging, error handling.
// For more details, see https://github.com/grpc/grpc-go/blob/master/interceptor.go
func apiServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	glog.Infof("%v handler starting", info.FullMethod)
	resp, err = handler(ctx, req)
	if err != nil {
		util.LogError(util.Wrapf(err, "%s call failed", info.FullMethod))
		// Convert error to gRPC errors
		err = util.ToGRPCError(err)
		return
	}
	glog.Infof("%v handler finished", info.FullMethod)

	return
}
