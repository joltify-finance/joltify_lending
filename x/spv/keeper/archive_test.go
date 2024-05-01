package keeper_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
)

// Test suite used for all keeper tests
type ArchiveTestSuite struct {
	suite.Suite
	keeper *spvkeeper.Keeper
	app    types.MsgServer
	ctx    sdk.Context
}

func TestArchiveTestSuite(t *testing.T) {
	suite.Run(t, new(ArchiveTestSuite))
}

// The default state used by each test
func (suite *ArchiveTestSuite) SetupTest() {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	lapp, k, _, _, _, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%

	suite.ctx = ctx
	suite.keeper = k
	suite.app = lapp
}

func (suite *ArchiveTestSuite) TestArchiveClass() {
	mockClass := nft.Class{
		Id:     "nft1",
		Symbol: "NFT1",
	}

	t1 := suite.ctx.BlockTime().Add(time.Second * 5)
	suite.ctx = suite.ctx.WithBlockTime(t1)
	err := suite.keeper.NftKeeper.SaveClass(suite.ctx, mockClass)
	suite.Require().NoError(err)

	suite.keeper.ArchiveClass(suite.ctx, "nft1")
	_, ok := suite.keeper.NftKeeper.GetClass(suite.ctx, "archive-nft1")
	suite.Assert().True(ok)

	suite.Assert().Panics(func() {
		suite.keeper.ArchiveClass(suite.ctx, "nft2")
	})
}

func (suite *ArchiveTestSuite) TestArchiveNFT() {
	mockNFT := nft.NFT{
		Id:      "nft1",
		ClassId: "class-1",
	}

	t1 := suite.ctx.BlockTime().Add(time.Second * 5)
	suite.ctx = suite.ctx.WithBlockTime(t1)

	mockAddr := sdk.AccAddress("addr1")

	err := suite.keeper.NftKeeper.Mint(suite.ctx, mockNFT, mockAddr)
	suite.Assert().NoError(err)

	mockNFT2 := nft.NFT{
		Id:      "nft2",
		ClassId: "class-2",
	}

	err = suite.keeper.NftKeeper.Mint(suite.ctx, mockNFT2, mockAddr)
	suite.Assert().NoError(err)

	mockclass1 := nft.Class{
		Id: "class-1",
	}
	mockclass2 := nft.Class{
		Id: "class-2",
	}

	err = suite.keeper.NftKeeper.SaveClass(suite.ctx, mockclass1)
	suite.Require().NoError(err)
	err = suite.keeper.NftKeeper.SaveClass(suite.ctx, mockclass2)
	suite.Require().NoError(err)

	err = suite.keeper.ArchiveNFT(suite.ctx, "class-1", "nft1")
	suite.Require().NoError(err)

	_, found := suite.keeper.NftKeeper.GetNFT(suite.ctx, "class-1", "nft1")
	suite.Assert().False(found)
	aa := fmt.Sprintf("archive-nft1-%v", t1.Unix())
	a, found := suite.keeper.NftKeeper.GetNFT(suite.ctx, "archive-class-1", aa)
	suite.Assert().True(found)
	suite.Assert().Equal(aa, a.Id)

	t2 := suite.ctx.BlockTime().Add(time.Second * 10)
	suite.ctx = suite.ctx.WithBlockTime(t2)

	err = suite.keeper.ArchiveNFT(suite.ctx, "class-2", "nft2")
	suite.Require().NoError(err)
	_, found = suite.keeper.NftKeeper.GetNFT(suite.ctx, "class-2", "nft2")
	suite.Assert().False(found)

	aa = fmt.Sprintf("archive-nft2-%v", t2.Unix())
	a, found = suite.keeper.NftKeeper.GetNFT(suite.ctx, "archive-class-2", aa)
	suite.Assert().True(found)
	suite.Assert().Equal(aa, a.Id)

	err = suite.keeper.ArchiveNFT(suite.ctx, "class-2", "nft2")
	suite.Require().ErrorContains(err, "nft not found")
}
