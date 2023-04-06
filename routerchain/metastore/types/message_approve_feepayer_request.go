package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveFeepayerRequest = "approve_feepayer_request"

var _ sdk.Msg = &MsgApproveFeepayerRequest{}

func NewMsgApproveFeepayerRequest(feepayer string, chainType uint64, chainId string, daapAddress string) *MsgApproveFeepayerRequest {
	return &MsgApproveFeepayerRequest{
		FeePayer:    feepayer,
		ChainType:   chainType,
		ChainId:     chainId,
		DappAddress: daapAddress,
	}
}

func (msg *MsgApproveFeepayerRequest) Route() string {
	return RouterKey
}

func (msg *MsgApproveFeepayerRequest) Type() string {
	return TypeMsgApproveFeepayerRequest
}

func (msg *MsgApproveFeepayerRequest) GetSigners() []sdk.AccAddress {
	feepayer, err := sdk.AccAddressFromBech32(msg.FeePayer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{feepayer}
}

func (msg *MsgApproveFeepayerRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveFeepayerRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FeePayer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid feepayer address (%s)", err)
	}
	return nil
}
