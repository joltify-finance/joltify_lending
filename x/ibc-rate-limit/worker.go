package ibc_rate_limit

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

func (i *ICS4Wrapper) whecherOnWhiteBanList(ctx sdk.Context, data []byte) (bool, bool, error) {
	var tdata transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(data, &tdata); err != nil {
		return false, false, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-20 transfer packet data: %s", err.Error())
	}
	ret := i.quotaKeeper.WhetherOnwhitelist(ctx, "ibc", tdata.Sender)
	retBan := i.quotaKeeper.WhetherOnBanlist(ctx, "ibc", tdata.Sender)
	return ret, retBan, nil
}

func (i *ICS4Wrapper) UpdateQuota(ctx sdk.Context, seq uint64, data []byte) error {
	var tdata transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(data, &tdata); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-20 transfer packet data: %s", err.Error())
	}

	denom := tdata.Denom
	if strings.ContainsAny(tdata.Denom, "/") {
		dataHash := sha256.Sum256([]byte(tdata.Denom))
		denom = hex.EncodeToString(dataHash[:])
	}

	token := strings.Join([]string{tdata.Amount, denom}, "")
	tokenAmount, err := sdk.ParseCoinNormalized(token)
	if err != nil {
		return err
	}
	err = i.quotaKeeper.UpdateQuota(ctx, sdk.NewCoins(tokenAmount), tdata.Sender, seq, "ibc")
	return err
}

func (i *ICS4Wrapper) RevokeQuotaHistory(ctx sdk.Context, seq uint64) {
	i.quotaKeeper.RevokeHistory(ctx, "ibc", seq)
}
