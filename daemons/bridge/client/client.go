package client

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/log"
	"github.com/ethereum/go-ethereum/ethclient"
	appflags "github.com/joltify-finance/joltify_lending/app/flags"
	"github.com/joltify-finance/joltify_lending/daemons/bridge/api"
	"github.com/joltify-finance/joltify_lending/daemons/bridge/client/types"
	"github.com/joltify-finance/joltify_lending/daemons/bridge/client/types/constants"
	daemonflags "github.com/joltify-finance/joltify_lending/daemons/flags"
	daemontypes "github.com/joltify-finance/joltify_lending/daemons/types"
	libtime "github.com/joltify-finance/joltify_lending/lib/time"
	bridgetypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/types"
)

// Client encapsulates the logic and interface for the bridge daemon. The bridge daemon periodically queries the
// Ethereum blockchain for new bridge events and relays them to the Cosmos gRPC server.
type Client struct {
	daemontypes.HealthCheckable

	// logger is the logger used by the bridge daemon.
	logger log.Logger
}

func NewClient(logger log.Logger) *Client {
	logger = logger.With(log.ModuleKey, constants.BridgeDaemonModuleName)
	return &Client{
		HealthCheckable: daemontypes.NewTimeBoundedHealthCheckable(
			constants.BridgeDaemonModuleName,
			&libtime.TimeProviderImpl{},
			logger,
		),
		logger: logger,
	}
}

// Start begins a job that periodically runs the RunBridgeDaemonTaskLoop function.
func (c *Client) Start(
	ctx context.Context,
	flags daemonflags.DaemonFlags,
	appFlags appflags.Flags,
	grpcClient daemontypes.GrpcClient,
) error {
	// Log the daemon flags.
	c.logger.Info(
		"Starting bridge daemon with flags",
		"BridgeFlags", flags.Bridge,
	)

	// Panic if EthRpcEndpoint is empty.
	if flags.Bridge.EthRpcEndpoint == "" {
		return fmt.Errorf("flag %s is not set", daemonflags.FlagBridgeDaemonEthRpcEndpoint)
	}

	// Make a connection to the Cosmos gRPC query services.
	queryConn, err := grpcClient.NewTcpConnection(ctx, appFlags.GrpcAddress)
	if err != nil {
		c.logger.Error("Failed to establish gRPC connection to Cosmos gRPC query services", "error", err)
		return err
	}
	defer func() {
		if connErr := grpcClient.CloseConnection(queryConn); connErr != nil {
			c.logger.Error("Failed to close gRPC connection to Cosmos gRPC query services", "error", connErr)
		}
	}()

	// Make a connection to the private daemon gRPC server.
	daemonConn, err := grpcClient.NewGrpcConnection(ctx, flags.Shared.SocketAddress)
	if err != nil {
		c.logger.Error("Failed to establish gRPC connection to socket address", "error", err)
		return err
	}
	defer func() {
		if connErr := grpcClient.CloseConnection(daemonConn); connErr != nil {
			c.logger.Error("Failed to close gRPC connection to Cosmos gRPC query services", "error", connErr)
		}
	}()

	// Initialize gRPC clients from query connection and daemon server connection.
	queryClient := bridgetypes.NewQueryClient(queryConn)
	serviceClient := api.NewBridgeServiceClient(daemonConn)

	// Initialize an Ethereum client from an RPC endpoint.
	ethClient, err := ethclient.Dial(flags.Bridge.EthRpcEndpoint)
	if err != nil {
		c.logger.Error("Failed to establish connection to Ethereum node", "error", err)
		return err
	}
	defer func() { ethClient.Close() }()

	ticker := time.NewTicker(time.Duration(flags.Bridge.LoopDelayMs) * time.Millisecond)
	stop := make(chan bool, 1)
	// Run the main task loop at an interval.
	StartBridgeDaemonTaskLoop(
		ctx,
		c,
		ticker,
		stop,
		&SubTaskRunnerImpl{},
		ethClient,
		queryClient,
		serviceClient,
	)

	return nil
}

// StartBridgeDaemonTaskLoop operates the continuous loop that runs the bridge daemon. It receives as arguments
// a ticker and a stop channel that are used to control and halt the loop.
func StartBridgeDaemonTaskLoop(
	ctx context.Context,
	c *Client,
	ticker *time.Ticker,
	stop <-chan bool,
	s SubTaskRunner,
	ethClient types.EthClient,
	queryClient bridgetypes.QueryClient,
	serviceClient api.BridgeServiceClient,
) {
	// Run the main task loop at an interval.
	for {
		select {
		case <-ticker.C:
			if err := s.RunBridgeDaemonTaskLoop(
				ctx,
				c.logger,
				ethClient,
				queryClient,
				serviceClient,
			); err == nil {
				c.ReportSuccess()
			} else {
				// TODO(DEC-947): Move daemon shutdown to application.
				c.logger.Error("Bridge daemon returned error", "error", err)
				c.ReportFailure(err)
			}
		case <-stop:
			return
		}
	}
}