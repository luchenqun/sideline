package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMinConfirmSubmitHeight = []byte("MinConfirmSubmitHeight")
	// TODO: Determine the default value
	DefaultMinConfirmSubmitHeight uint64 = 300
)

var (
	KeyMinConfirmJudgeHeight = []byte("MinConfirmJudgeHeight")
	// TODO: Determine the default value
	DefaultMinConfirmJudgeHeight uint64 = 600
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	minConfirmSubmitHeight uint64,
	minConfirmJudgeHeight uint64,
) Params {
	return Params{
		MinConfirmSubmitHeight: minConfirmSubmitHeight,
		MinConfirmJudgeHeight:  minConfirmJudgeHeight,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinConfirmSubmitHeight,
		DefaultMinConfirmJudgeHeight,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinConfirmSubmitHeight, &p.MinConfirmSubmitHeight, validateMinConfirmSubmitHeight),
		paramtypes.NewParamSetPair(KeyMinConfirmJudgeHeight, &p.MinConfirmJudgeHeight, validateMinConfirmJudgeHeight),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMinConfirmSubmitHeight(p.MinConfirmSubmitHeight); err != nil {
		return err
	}

	if err := validateMinConfirmJudgeHeight(p.MinConfirmJudgeHeight); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateMinConfirmSubmitHeight validates the MinConfirmSubmitHeight param
func validateMinConfirmSubmitHeight(v interface{}) error {
	minConfirmSubmitHeight, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = minConfirmSubmitHeight

	return nil
}

// validateMinConfirmJudgeHeight validates the MinConfirmJudgeHeight param
func validateMinConfirmJudgeHeight(v interface{}) error {
	minConfirmJudgeHeight, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = minConfirmJudgeHeight

	return nil
}
