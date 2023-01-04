package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMinConfirmSubmitHeight            = []byte("MinConfirmSubmitHeight")
	DefaultMinConfirmSubmitHeight uint64 = 300
)

var (
	KeyMinConfirmJudgeHeight            = []byte("MinConfirmJudgeHeight")
	DefaultMinConfirmJudgeHeight uint64 = 600
)

var (
	KeyRegistrationFee     = []byte("RegistrationFee")
	DefaultRegistrationFee = "100wrmb"
)

var (
	KeyValidatorCommission            = []byte("ValidatorCommission")
	DefaultValidatorCommission uint64 = 10
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	minConfirmSubmitHeight uint64,
	minConfirmJudgeHeight uint64,
	registrationFee string,
	validatorCommission uint64,
) Params {
	return Params{
		MinConfirmSubmitHeight: minConfirmSubmitHeight,
		MinConfirmJudgeHeight:  minConfirmJudgeHeight,
		RegistrationFee:        registrationFee,
		ValidatorCommission:    validatorCommission,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinConfirmSubmitHeight,
		DefaultMinConfirmJudgeHeight,
		DefaultRegistrationFee,
		DefaultValidatorCommission,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinConfirmSubmitHeight, &p.MinConfirmSubmitHeight, validateMinConfirmSubmitHeight),
		paramtypes.NewParamSetPair(KeyMinConfirmJudgeHeight, &p.MinConfirmJudgeHeight, validateMinConfirmJudgeHeight),
		paramtypes.NewParamSetPair(KeyRegistrationFee, &p.RegistrationFee, validateRegistrationFee),
		paramtypes.NewParamSetPair(KeyValidatorCommission, &p.ValidatorCommission, validateValidatorCommission),
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

	if err := validateRegistrationFee(p.RegistrationFee); err != nil {
		return err
	}

	if err := validateValidatorCommission(p.ValidatorCommission); err != nil {
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

// validateRegistrationFee validates the RegistrationFee param
func validateRegistrationFee(v interface{}) error {
	registrationFee, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = registrationFee

	return nil
}

// validateValidatorCommission validates the ValidatorCommission param
func validateValidatorCommission(v interface{}) error {
	validatorCommission, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = validatorCommission

	return nil
}
