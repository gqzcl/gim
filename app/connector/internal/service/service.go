package service

import (
	pb "arod-im/api/connector/v1"
	logicV1 "arod-im/api/logic/v1"
	"arod-im/app/connector/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewConnectorService)

var (
	// grpc options
	grpcKeepAliveTime    = time.Duration(10) * time.Second
	grpcKeepAliveTimeout = time.Duration(3) * time.Second
	grpcBackoffMaxDelay  = time.Duration(3) * time.Second
	grpcMaxSendMsgSize   = 1 << 24
	grpcMaxCallMsgSize   = 1 << 24
)

const (
	// grpc options
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
)

type ConnectorService struct {
	pb.UnimplementedConnectorServer

	bc  *biz.BucketUsecase
	log *log.Helper
	//ws          *websocket.Server
	LogicClient []logicV1.LogicClient

	// server ip
	Address string
}

func NewConnectorService(bc *biz.BucketUsecase, logger log.Logger) *ConnectorService {
	c := &ConnectorService{
		bc:  bc,
		log: log.NewHelper(log.With(logger, "module", "connector")),
	}
	ctx := context.Background()
	conn, _ := grpc.DialContext(ctx, "127.0.0.1:9003",
		[]grpc.DialOption{
			grpc.WithInitialWindowSize(grpcInitialWindowSize),
			grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
			grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                grpcKeepAliveTime,
				Timeout:             grpcKeepAliveTimeout,
				PermitWithoutStream: true,
			}),
		}...,
	)
	c.LogicClient = append(c.LogicClient, logicV1.NewLogicClient(conn))
	return c
}
