package types_test

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/require"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	wasmtesting "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/testing"
	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
	ibcerrors "github.com/cosmos/ibc-go/v8/modules/core/errors"
	ibctesting "github.com/cosmos/ibc-go/v8/testing"
)

func TestMsgStoreCodeValidateBasic(t *testing.T) {
	signer := sdk.AccAddress(ibctesting.TestAccAddress).String()
	testCases := []struct {
		name   string
		msg    *types.MsgStoreCode
		expErr error
	}{
		{
			"success: valid signer address, valid length code",
			types.NewMsgStoreCode(signer, wasmtesting.Code),
			nil,
		},
		{
			"failure: code is empty",
			types.NewMsgStoreCode(signer, []byte("")),
			types.ErrWasmEmptyCode,
		},
		{
			"failure: code is too large",
			types.NewMsgStoreCode(signer, make([]byte, types.MaxWasmSize+1)),
			types.ErrWasmCodeTooLarge,
		},
		{
			"failure: signer is invalid",
			types.NewMsgStoreCode("invalid", wasmtesting.Code),
			ibcerrors.ErrInvalidAddress,
		},
	}

	for _, tc := range testCases {
		tc := tc

		err := tc.msg.ValidateBasic()
		expPass := tc.expErr == nil
		if expPass {
			require.NoError(t, err)
		} else {
			require.ErrorIs(t, err, tc.expErr)
		}
	}
}

func (suite *TypesTestSuite) TestMsgStoreCodeGetSigners() {
	testCases := []struct {
		name    string
		address sdk.AccAddress
		expPass bool
	}{
		{"success: valid address", sdk.AccAddress(ibctesting.TestAccAddress), true},
		{"failure: nil address", nil, false},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.SetupWasmWithMockVM()

			address := tc.address
			msg := types.NewMsgStoreCode(address.String(), wasmtesting.Code)

			signers, _, err := GetSimApp(suite.chainA).AppCodec().GetMsgV1Signers(msg)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(address.Bytes(), signers[0])
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func TestMsgMigrateContractValidateBasic(t *testing.T) {
	signer := sdk.AccAddress(ibctesting.TestAccAddress).String()
	validCodeHash := sha256.Sum256(wasmtesting.Code)
	validMigrateMsg := []byte("{}")

	testCases := []struct {
		name   string
		msg    *types.MsgMigrateContract
		expErr error
	}{
		{
			"success: valid signer address, valid code hash, valid migrate msg",
			types.NewMsgMigrateContract(signer, defaultWasmClientID, validCodeHash[:], validMigrateMsg),
			nil,
		},
		{
			"failure: invalid signer address",
			types.NewMsgMigrateContract(ibctesting.InvalidID, defaultWasmClientID, validCodeHash[:], validMigrateMsg),
			ibcerrors.ErrInvalidAddress,
		},
		{
			"failure: clientID is not a valid client identifier",
			types.NewMsgMigrateContract(signer, ibctesting.InvalidID, validCodeHash[:], validMigrateMsg),
			host.ErrInvalidID,
		},
		{
			"failure: clientID is not a wasm client identifier",
			types.NewMsgMigrateContract(signer, ibctesting.FirstClientID, validCodeHash[:], validMigrateMsg),
			host.ErrInvalidID,
		},
		{
			"failure: code hash is nil",
			types.NewMsgMigrateContract(signer, defaultWasmClientID, nil, validMigrateMsg),
			errorsmod.Wrap(types.ErrInvalidCodeHash, "code hash cannot be empty"),
		},
		{
			"failure: code hash is empty",
			types.NewMsgMigrateContract(signer, defaultWasmClientID, []byte{}, validMigrateMsg),
			errorsmod.Wrap(types.ErrInvalidCodeHash, "code hash cannot be empty"),
		},
		{
			"failure: code hash is not 32 bytes",
			types.NewMsgMigrateContract(signer, defaultWasmClientID, []byte{1}, validMigrateMsg),
			errorsmod.Wrapf(types.ErrInvalidCodeHash, "expected length of 32 bytes, got %d", 1),
		},
		{
			"failure: migrateMsg is nil",
			types.NewMsgMigrateContract(signer, defaultWasmClientID, validCodeHash[:], nil),
			errorsmod.Wrap(ibcerrors.ErrInvalidRequest, "migrate message cannot be empty"),
		},
		{
			"failure: migrateMsg is empty",
			types.NewMsgMigrateContract(signer, defaultWasmClientID, validCodeHash[:], []byte("")),
			errorsmod.Wrap(ibcerrors.ErrInvalidRequest, "migrate message cannot be empty"),
		},
	}

	for _, tc := range testCases {
		tc := tc

		err := tc.msg.ValidateBasic()
		expPass := tc.expErr == nil
		if expPass {
			require.NoError(t, err)
		} else {
			require.ErrorIs(t, err, tc.expErr, tc.name)
		}
	}
}

func TestMsgRemoveCodeHashValidateBasic(t *testing.T) {
	signer := sdk.AccAddress(ibctesting.TestAccAddress).String()

	codeHash := sha256.Sum256(wasmtesting.Code)

	testCases := []struct {
		name   string
		msg    *types.MsgRemoveChecksum
		expErr error
	}{
		{
			"success: valid signer address, valid length code hash",
			types.NewMsgRemoveChecksum(signer, codeHash[:]),
			nil,
		},
		{
			"failure: code hash is empty",
			types.NewMsgRemoveChecksum(signer, []byte("")),
			types.ErrInvalidCodeHash,
		},
		{
			"failure: code hash is nil",
			types.NewMsgRemoveChecksum(signer, nil),
			types.ErrInvalidCodeHash,
		},
		{
			"failure: signer is invalid",
			types.NewMsgRemoveChecksum(ibctesting.InvalidID, codeHash[:]),
			ibcerrors.ErrInvalidAddress,
		},
	}

	for _, tc := range testCases {
		tc := tc

		err := tc.msg.ValidateBasic()

		if tc.expErr == nil {
			require.NoError(t, err, tc.name)
		} else {
			require.ErrorIs(t, err, tc.expErr, tc.name)
		}
	}
}
