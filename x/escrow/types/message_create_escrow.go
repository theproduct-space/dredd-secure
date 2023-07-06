package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"strconv"
	"time"
)

const TypeMsgCreateEscrow = "create_escrow"

var _ sdk.Msg = &MsgCreateEscrow{}

func NewMsgCreateEscrow(creator string, initiatorCoins sdk.Coins, fulfillerCoins sdk.Coins, startDate string, endDate string) *MsgCreateEscrow {
	return &MsgCreateEscrow{
		Creator:        creator,
		InitiatorCoins: initiatorCoins,
		FulfillerCoins: fulfillerCoins,
		StartDate:      startDate,
		EndDate:        endDate,
	}
}

func (msg *MsgCreateEscrow) Route() string {
	return RouterKey
}

func (msg *MsgCreateEscrow) Type() string {
	return TypeMsgCreateEscrow
}

func (msg *MsgCreateEscrow) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEscrow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEscrow) ValidateBasic() error {
	// Account validation
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Validating InitiatorCoins & FulfillerCoins are defined
	if (msg.InitiatorCoins == nil) {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "InitiatorCoins does not contain any Coin")
	}
	if (msg.FulfillerCoins == nil) {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "FulfillerCoins does not contain any Coin")
	}

	// Validating every coins
	for i := 0; i < len(msg.InitiatorCoins); i++ {
		if (!msg.InitiatorCoins[i].IsValid()) {
			return errors.Wrapf(sdkerrors.ErrInvalidCoins , "InitiatorCoins with denom %v is not a valid Coin object", msg.InitiatorCoins[i].Denom)
		}
		if (!msg.FulfillerCoins[i].IsValid()) {
			return errors.Wrapf(sdkerrors.ErrInvalidCoins , "FulfillerCoins with denom %v is not a valid Coin object", msg.FulfillerCoins[i].Denom)
		}
	}

	// Dates validation
	now := time.Now()
	unixTimeNow := now.Unix()

	endDateInt, errParseIntEndDate := strconv.ParseInt(msg.EndDate, 10, 64)
	if errParseIntEndDate != nil {
		return errors.Wrap(sdkerrors.ErrInvalidRequest , "Invalid endDate")
	}
	startDateInt, errParseIntStartDate := strconv.ParseInt(msg.StartDate, 10, 64)
	if errParseIntStartDate != nil {
		return errors.Wrap(sdkerrors.ErrInvalidRequest , "Invalid startdate")
	}

	// Validating end_date is in the future
	if (endDateInt < unixTimeNow) {
		return errors.Wrap(sdkerrors.ErrInvalidRequest , "End date is not in the future")
	}

	// Validating start_date is before end_date
	if (endDateInt < startDateInt) {
		return errors.Wrap(sdkerrors.ErrInvalidRequest , "End date is before start date")
	}

	return nil
}
