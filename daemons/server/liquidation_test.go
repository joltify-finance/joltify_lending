package server_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/daemons/liquidation/api"
	liquidationtypes "github.com/joltify-finance/joltify_lending/daemons/server/types/liquidations"
	"github.com/joltify-finance/joltify_lending/dydx_helper/mocks"
	"github.com/joltify-finance/joltify_lending/dydx_helper/testutil/constants"
	"github.com/joltify-finance/joltify_lending/dydx_helper/testutil/grpc"
	satypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/subaccounts/types"
	"github.com/stretchr/testify/require"
)

func TestLiquidateSubaccounts_Empty_Update_Liquidatable_SubaccountIds(t *testing.T) {
	mockGrpcServer := &mocks.GrpcServer{}
	mockFileHandler := &mocks.FileHandler{}
	daemonLiquidationInfo := liquidationtypes.NewDaemonLiquidationInfo()

	s := createServerWithMocks(
		t,
		mockGrpcServer,
		mockFileHandler,
	).WithDaemonLiquidationInfo(
		daemonLiquidationInfo,
	)
	_, err := s.LiquidateSubaccounts(grpc.Ctx, &api.LiquidateSubaccountsRequest{
		LiquidatableSubaccountIds: []satypes.SubaccountId{},
	})
	require.NoError(t, err)
	require.Empty(t, daemonLiquidationInfo.GetLiquidatableSubaccountIds())
}

func TestLiquidateSubaccounts_Multiple_Liquidatable_Subaccount_Ids(t *testing.T) {
	mockGrpcServer := &mocks.GrpcServer{}
	mockFileHandler := &mocks.FileHandler{}
	daemonLiquidationInfo := liquidationtypes.NewDaemonLiquidationInfo()

	s := createServerWithMocks(
		t,
		mockGrpcServer,
		mockFileHandler,
	).WithDaemonLiquidationInfo(
		daemonLiquidationInfo,
	)

	expectedSubaccountIds := []satypes.SubaccountId{
		constants.Alice_Num1,
		constants.Bob_Num0,
		constants.Carl_Num0,
	}
	_, err := s.LiquidateSubaccounts(grpc.Ctx, &api.LiquidateSubaccountsRequest{
		LiquidatableSubaccountIds: expectedSubaccountIds,
	})
	require.NoError(t, err)

	actualSubaccountIds := daemonLiquidationInfo.GetLiquidatableSubaccountIds()
	require.Equal(t, expectedSubaccountIds, actualSubaccountIds)
}

func TestLiquidateSubaccounts_GetSetBlockHeight(t *testing.T) {
	mockGrpcServer := &mocks.GrpcServer{}
	mockFileHandler := &mocks.FileHandler{}
	daemonLiquidationInfo := liquidationtypes.NewDaemonLiquidationInfo()

	s := createServerWithMocks(
		t,
		mockGrpcServer,
		mockFileHandler,
	).WithDaemonLiquidationInfo(
		daemonLiquidationInfo,
	)
	_, err := s.LiquidateSubaccounts(grpc.Ctx, &api.LiquidateSubaccountsRequest{
		BlockHeight:               uint32(123),
		LiquidatableSubaccountIds: []satypes.SubaccountId{},
	})
	require.NoError(t, err)
	require.Equal(t, uint32(123), daemonLiquidationInfo.GetBlockHeight())
}

func TestLiquidateSubaccounts_Empty_Update_Negative_TNC_SubaccountIds(t *testing.T) {
	mockGrpcServer := &mocks.GrpcServer{}
	mockFileHandler := &mocks.FileHandler{}
	daemonLiquidationInfo := liquidationtypes.NewDaemonLiquidationInfo()

	s := createServerWithMocks(
		t,
		mockGrpcServer,
		mockFileHandler,
	).WithDaemonLiquidationInfo(
		daemonLiquidationInfo,
	)
	_, err := s.LiquidateSubaccounts(grpc.Ctx, &api.LiquidateSubaccountsRequest{
		NegativeTncSubaccountIds: []satypes.SubaccountId{},
	})
	require.NoError(t, err)
	require.Empty(t, daemonLiquidationInfo.GetNegativeTncSubaccountIds())
}

func TestLiquidateSubaccounts_Multiple_Negative_TNC_Subaccount_Ids(t *testing.T) {
	mockGrpcServer := &mocks.GrpcServer{}
	mockFileHandler := &mocks.FileHandler{}
	daemonLiquidationInfo := liquidationtypes.NewDaemonLiquidationInfo()

	s := createServerWithMocks(
		t,
		mockGrpcServer,
		mockFileHandler,
	).WithDaemonLiquidationInfo(
		daemonLiquidationInfo,
	)

	expectedSubaccountIds := []satypes.SubaccountId{
		constants.Alice_Num1,
		constants.Bob_Num0,
		constants.Carl_Num0,
	}
	_, err := s.LiquidateSubaccounts(grpc.Ctx, &api.LiquidateSubaccountsRequest{
		NegativeTncSubaccountIds: expectedSubaccountIds,
	})
	require.NoError(t, err)

	actualSubaccountIds := daemonLiquidationInfo.GetNegativeTncSubaccountIds()
	require.Equal(t, expectedSubaccountIds, actualSubaccountIds)
}
