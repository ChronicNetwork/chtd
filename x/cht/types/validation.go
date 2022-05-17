package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// MaxLabelSize is the longest label that can be used when Instantiating a contract
	MaxLabelSize = 128 // extension point for chains to customize via compile flag.

	// MaxChtSize is the largest a compiled contract code can be when storing code on chain
	MaxChtSize = 800 * 1024 // extension point for chains to customize via compile flag.
)

func validateChtCode(s []byte) error {
	if len(s) == 0 {
		return sdkerrors.Wrap(ErrEmpty, "is required")
	}
	if len(s) > MaxChtSize {
		return sdkerrors.Wrapf(ErrLimit, "cannot be longer than %d bytes", MaxChtSize)
	}
	return nil
}

func validateLabel(label string) error {
	if label == "" {
		return sdkerrors.Wrap(ErrEmpty, "is required")
	}
	if len(label) > MaxLabelSize {
		return sdkerrors.Wrap(ErrLimit, "cannot be longer than 128 characters")
	}
	return nil
}
