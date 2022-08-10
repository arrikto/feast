package main

import (
	"context"
	"flag"
	"math"
	"net"
	"net/http"
	"strings"

	api "github.com/feast-dev/feast/backend/api/go_client"

	"github.com/feast-dev/feast/backend/src/apiserver/common"
	"github.com/feast-dev/feast/backend/src/apiserver/resource"
	"github.com/feast-dev/feast/backend/src/apiserver/server"

	"github.com/fsnotify/fsnotify"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var (
	rpcPortFlag  = flag.String("rpcPortFlag", ":8887", "RPC Port")
	httpPortFlag = flag.String("httpPortFlag", ":8888", "Http Proxy Port")
	configPath   = flag.String("config", "", "Path to JSON file containing config")
)

type RegisterHttpHandlerFromEndpoint func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

func main() {
	flag.Parse()

	initConfig()
	clientManager := newClientManager()
	resourceManager := resource.NewResourceManager(&clientManager)

	go startRpcServer(resourceManager)
	startHttpProxy(resourceManager)

	clientManager.Close()
}

// A custom http request header matcher to pass on the user identity
// Reference: https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_your_gateway/#mapping-from-http-request-headers-to-grpc-client-metadata
func grpcCustomMatcher(key string) (string, bool) {
	if strings.EqualFold(key, common.GetKubeflowUserIDHeader()) {
		return strings.ToLower(key), true
	}

	return strings.ToLower(key), false
}

func startRpcServer(resourceManager *resource.ResourceManager) {
	glog.Info("Starting RPC server")
	listener, err := net.Listen("tcp", *rpcPortFlag)
	if err != nil {
		glog.Fatalf("Failed to start RPC server: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(apiServerInterceptor), grpc.MaxRecvMsgSize(math.MaxInt32))
	api.RegisterDataSourceServiceServer(s, server.NewDataSourceServer(resourceManager, &server.DataSourceServerOptions{}))
	api.RegisterEntityServiceServer(s, server.NewEntityServer(resourceManager, &server.EntityServerOptions{}))
	api.RegisterFeatureServiceServiceServer(s, server.NewFeatureServiceServer(resourceManager, &server.FeatureServiceServerOptions{}))
	api.RegisterFeatureViewServiceServer(s, server.NewFeatureViewServer(resourceManager, &server.FeatureViewServerOptions{}))
	api.RegisterInfraObjectServiceServer(s, server.NewInfraObjectServer(resourceManager, &server.InfraObjectServerOptions{}))
	api.RegisterOnDemandFeatureViewServiceServer(s, server.NewOnDemandFeatureViewServer(resourceManager, &server.OnDemandFeatureViewServerOptions{}))
	api.RegisterProjectServiceServer(s, server.NewProjectServer(resourceManager, &server.ProjectServerOptions{}))
	api.RegisterRequestFeatureViewServiceServer(s, server.NewRequestFeatureViewServer(resourceManager, &server.RequestFeatureViewServerOptions{}))
	api.RegisterSavedDatasetServiceServer(s, server.NewSavedDatasetServer(resourceManager, &server.SavedDatasetServerOptions{}))

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		glog.Fatalf("Failed to serve rpc listener: %v", err)
	}
	glog.Info("RPC server started")
}

func startHttpProxy(resourceManager *resource.ResourceManager) {
	glog.Info("Starting Http Proxy")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create gRPC HTTP MUX and register services.
	runtimeMux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(grpcCustomMatcher))
	registerHttpHandlerFromEndpoint(api.RegisterDataSourceServiceHandlerFromEndpoint, "DataSourceService", ctx, runtimeMux)
	registerHttpHandlerFromEndpoint(api.RegisterEntityServiceHandlerFromEndpoint, "EntityService", ctx, runtimeMux)
	registerHttpHandlerFromEndpoint(api.RegisterFeatureServiceServiceHandlerFromEndpoint, "FeatureServiceService", ctx, runtimeMux)
	registerHttpHandlerFromEndpoint(api.RegisterFeatureViewServiceHandlerFromEndpoint, "FeatureViewService", ctx, runtimeMux)
	registerHttpHandlerFromEndpoint(api.RegisterInfraObjectServiceHandlerFromEndpoint, "InfraObjectService", ctx, runtimeMux)
	registerHttpHandlerFromEndpoint(api.RegisterOnDemandFeatureViewServiceHandlerFromEndpoint, "OnDemandFeatureViewService", ctx, runtimeMux)
	registerHttpHandlerFromEndpoint(api.RegisterProjectServiceHandlerFromEndpoint, "ProjectService", ctx, runtimeMux)
	registerHttpHandlerFromEndpoint(api.RegisterRequestFeatureViewServiceHandlerFromEndpoint, "RequestFeatureViewService", ctx, runtimeMux)
	registerHttpHandlerFromEndpoint(api.RegisterSavedDatasetServiceHandlerFromEndpoint, "SavedDatasetService", ctx, runtimeMux)

	http.ListenAndServe(*httpPortFlag, runtimeMux)
	glog.Info("Http Proxy started")
}

func registerHttpHandlerFromEndpoint(handler RegisterHttpHandlerFromEndpoint, serviceName string, ctx context.Context, mux *runtime.ServeMux) {
	endpoint := "localhost" + *rpcPortFlag
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32))}

	if err := handler(ctx, mux, endpoint, opts); err != nil {
		glog.Fatalf("Failed to register %v handler: %v", serviceName, err)
	}
}

func initConfig() {
	// Import environment variable, support nested vars e.g. DBCONFIG_DRIVERNAME
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
	// We need empty string env var for e.g. KUBEFLOW_USERID_PREFIX.
	viper.AllowEmptyEnv(true)

	// Set configuration file name. The format is auto detected in this case.
	viper.SetConfigName("config")
	viper.AddConfigPath(*configPath)
	err := viper.ReadInConfig()
	if err != nil {
		glog.Fatalf("Fatal error config file: %s", err)
	}

	// Watch for configuration change
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// Read in config again
		viper.ReadInConfig()
	})
}
