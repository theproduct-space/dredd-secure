package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const (
	DefaultAskCount       = uint64(16)
	DefaultMinCount       = uint64(10)
	DefaultMinDsCount     = uint64(3)
	DefaultPrepareGasBase = uint64(3000)
	DefaultPrepareGasEach = uint64(600)
	DefaultExecuteGasBase = uint64(70000)
	DefaultExecuteGasEach = uint64(7500)
)

var DefaultFeeLimit = sdk.NewCoins(sdk.NewInt64Coin("uband", 1000000))

var (
	KeyAskCount       = []byte("AskCount")
	KeyMinCount       = []byte("MinCount")
	KeyMinDsCount     = []byte("MinDsCount")
	KeyPrepareGasBase = []byte("PrepareGasBase")
	KeyPrepareGasEach = []byte("PrepareGasEach")
	KeyExecuteGasBase = []byte("ExecuteGasBase")
	KeyExecuteGasEach = []byte("ExecuteGasEach")
	KeyFeeLimit       = []byte("FeeLimit")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	askCount, minCount, minDsCount, prepareGasBase, prepareGasEach, executeGasBase, executeGasEach uint64,
	feeLimit sdk.Coins,
) Params {
	return Params{
		AskCount:       askCount,
		MinCount:       minCount,
		MinDsCount:     minDsCount,
		PrepareGasBase: prepareGasBase,
		PrepareGasEach: prepareGasEach,
		ExecuteGasBase: executeGasBase,
		ExecuteGasEach: executeGasEach,
		FeeLimit:       feeLimit,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultAskCount,
		DefaultMinCount,
		DefaultMinDsCount,
		DefaultPrepareGasBase,
		DefaultPrepareGasEach,
		DefaultExecuteGasBase,
		DefaultExecuteGasEach,
		DefaultFeeLimit,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAskCount, &p.AskCount, validateUint64("ask count")),
		paramtypes.NewParamSetPair(KeyMinCount, &p.MinCount, validateUint64("min count")),
		paramtypes.NewParamSetPair(KeyMinDsCount, &p.MinDsCount, validateUint64("min ds count")),
		paramtypes.NewParamSetPair(KeyPrepareGasBase, &p.PrepareGasBase, validateUint64("prepare gas base")),
		paramtypes.NewParamSetPair(KeyPrepareGasEach, &p.PrepareGasEach, validateUint64("prepare gas each")),
		paramtypes.NewParamSetPair(KeyExecuteGasBase, &p.ExecuteGasBase, validateUint64("execute gas base")),
		paramtypes.NewParamSetPair(KeyExecuteGasEach, &p.ExecuteGasEach, validateUint64("execute gas each")),
		paramtypes.NewParamSetPair(KeyFeeLimit, &p.FeeLimit, validateFeeLimit),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateUint64(name string) func(interface{}) error {
	return func(i interface{}) error {
		v, ok := i.(uint64)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}
		if v <= 0 {
			return fmt.Errorf("%s must be positive: %d", name, v)
		}
		return nil
	}
}

func validateFeeLimit(i interface{}) error {
	_, ok := i.(sdk.Coins)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "type: %T, expected sdk.Coins", i)
	}
	return nil
}
