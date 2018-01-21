package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/johanbrandhorst/chronic-pain-tracker/api"
	"github.com/johanbrandhorst/chronic-pain-tracker/proto"
)

type psqlURL url.URL

func (p *psqlURL) Decode(in string) error {
	u, err := url.Parse(in)
	if err != nil {
		return err
	}

	if u.Scheme != "postgres" && u.Scheme != "postgresql" {
		return errors.New(`schema should be "postgres" or "postgresql"`)
	}

	*p = psqlURL(*u)
	return nil
}

func (p *psqlURL) URL() url.URL {
	return (url.URL)(*p)
}

type config struct {
	GRPCAddr    string  `default:"localhost:8081" envconfig:"GRPC_ADDR" desc:"Address to serve the gRPC Server on."`
	GatewayAddr string  `default:"0.0.0.0:8080" split_words:"true" desc:"Address to serve the gRPC-Gateway on."`
	PostgresURL psqlURL `required:"true" envconfig:"POSTGRES_URL" desc:"URL to the Postgres database used."`
}

func main() {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Formatter = &logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: time.StampMilli,
		FullTimestamp:   true,
	}

	envOpts := config{}
	err := envconfig.Process("", &envOpts)
	if err != nil {
		envconfig.Usage("", &envOpts)
		logger.WithError(err).Fatal()
	}

	s := grpc.NewServer()

	srv, err := api.NewServer(logger, envOpts.PostgresURL.URL())
	if err != nil {
		logger.WithError(err).Fatal("Failed to connect to Postgres")
	}

	proto.RegisterPainTrackerServer(s, srv)
	proto.RegisterMonitorServer(s, srv)

	go func() {
		lis, err := net.Listen("tcp", envOpts.GRPCAddr)
		if err != nil {
			logger.WithError(err).Fatal("Failed to start grpc listener")
		}

		err = s.Serve(lis)
		if err != nil {
			logger.WithError(err).Fatal("Failed to serve gRPC server")
		}
	}()

	cc, err := grpc.Dial(envOpts.GRPCAddr, grpc.WithInsecure())
	if err != nil {
		logger.WithError(err).Fatal("Failed to dial gRPC server")
	}

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption("*", &runtime.JSONPb{
			EmitDefaults: true,
		}),
	)
	err = proto.RegisterPainTrackerHandler(context.Background(), mux, cc)
	if err != nil {
		logger.WithError(err).Fatal("Failed to register pain tracker in gRPC-gateway")
	}
	err = proto.RegisterMonitorHandler(context.Background(), mux, cc)
	if err != nil {
		logger.WithError(err).Fatal("Failed to register monitor in gRPC-gateway")
	}

	logger.Infoln("Serving on", envOpts.GatewayAddr)
	err = http.ListenAndServe(envOpts.GatewayAddr, mux)
	if err != nil {
		logger.WithError(err).Fatal("Failed to serve gRPC-gateway")
	}
}
