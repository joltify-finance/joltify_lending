package keeper

import (
	"encoding/base64"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var param types.Params
	k.paramstore.GetParamSet(ctx, &param)
	return param
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetSubmitter(ctx sdk.Context) (submitters []sdk.AccAddress) {
	k.paramstore.Get(ctx, types.KeyKycSubmitter, &submitters)
	return submitters
}

func (k Keeper) GetProjects(ctx sdk.Context) []*types.ProjectInfo {
	var projectsEncoded string
	k.paramstore.Get(ctx, types.KeyProjects, &projectsEncoded)

	val, err := base64.StdEncoding.DecodeString(projectsEncoded)
	if err != nil {
		panic("invalid encoded string")
	}

	var projects types.Projects
	err = proto.Unmarshal(val, &projects)
	if err != nil {
		panic("invalid encoded string")
	}

	return projects.Items
}
