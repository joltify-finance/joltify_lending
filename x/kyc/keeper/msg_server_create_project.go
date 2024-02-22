package keeper

import (
	"context"
	"encoding/base64"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
)

// CreateProject create a new project
func (k Keeper) CreateProject(goCtx context.Context, msg *types.MsgCreateProject) (*types.MSgCreateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != k.authority.String() {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority, expected %s, got %s", k.authority.String(), msg.Creator)
	}

	out, err := base64.StdEncoding.DecodeString(msg.EncodedProject)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidProject, "fail to decode the project base64 string: %v", err)
	}

	var project types.ProjectInfo
	err = proto.Unmarshal(out, &project)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidProject, "fail to unmarshal the project: %v", err)
	}

	err = types.ValidateProject(project)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidProject, "invalid project: %v", err)
	}
	projectID, err := k.SetProject(ctx, &project)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidProject, "invalid project: %v", err)
	}
	return &types.MSgCreateProjectResponse{ProjectId: strconv.FormatInt(int64(projectID), 10)}, nil
}
