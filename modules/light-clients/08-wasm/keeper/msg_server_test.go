package keeper_test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"

	wasmvm "github.com/CosmWasm/wasmvm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/internal/ibcwasm"
	wasmtesting "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/testing"
	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
	ibcerrors "github.com/cosmos/ibc-go/v8/modules/core/errors"
	localhost "github.com/cosmos/ibc-go/v8/modules/light-clients/09-localhost"
	ibctesting "github.com/cosmos/ibc-go/v8/testing"
)

func (suite *KeeperTestSuite) TestMsgStoreCode() {
	var (
		msg    *types.MsgStoreCode
		signer string
		data   []byte
	)

	testCases := []struct {
		name     string
		malleate func()
		expError error
	}{
		{
			"success",
			func() {
				msg = types.NewMsgStoreCode(signer, data)
			},
			nil,
		},
		{
			"fails with duplicate wasm code",
			func() {
				msg = types.NewMsgStoreCode(signer, data)

				_, err := GetSimApp(suite.chainA).WasmClientKeeper.StoreCode(suite.chainA.GetContext(), msg)
				suite.Require().NoError(err)
			},
			types.ErrWasmCodeExists,
		},
		{
			"fails with invalid wasm code",
			func() {
				msg = types.NewMsgStoreCode(signer, []byte{})
			},
			types.ErrWasmEmptyCode,
		},
		{
			"fails with unauthorized signer",
			func() {
				signer = suite.chainA.SenderAccount.GetAddress().String()
				msg = types.NewMsgStoreCode(signer, data)
			},
			ibcerrors.ErrUnauthorized,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			signer = authtypes.NewModuleAddress(govtypes.ModuleName).String()
			data, _ = os.ReadFile("../test_data/ics10_grandpa_cw.wasm.gz")

			tc.malleate()

			ctx := suite.chainA.GetContext()
			res, err := GetSimApp(suite.chainA).WasmClientKeeper.StoreCode(ctx, msg)
			events := ctx.EventManager().Events()

			if tc.expError == nil {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().NotEmpty(res.Checksum)

				// Verify events
				expectedEvents := sdk.Events{
					sdk.NewEvent(
						"store_wasm_code",
						sdk.NewAttribute(types.AttributeKeyWasmCodeHash, hex.EncodeToString(res.Checksum)),
					),
					sdk.NewEvent(
						sdk.EventTypeMessage,
						sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					),
				}

				for _, evt := range expectedEvents {
					suite.Require().Contains(events, evt)
				}
			} else {
				suite.Require().ErrorIs(err, tc.expError)
				suite.Require().Nil(res)
				suite.Require().Empty(events)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgMigrateContract() {
	oldCodeHash := sha256.Sum256(wasmtesting.Code)

	newByteCode := []byte("MockByteCode-TestMsgMigrateContract")

	govAcc := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	var (
		newCodeHash    []byte
		msg            *types.MsgMigrateContract
		expClientState *types.ClientState
	)

	testCases := []struct {
		name     string
		malleate func()
		expError error
	}{
		{
			"success: no update to client state",
			func() {
				msg = types.NewMsgMigrateContract(govAcc, defaultWasmClientID, newCodeHash, []byte("{}"))

				suite.mockVM.MigrateFn = func(_ wasmvm.Checksum, _ wasmvmtypes.Env, _ []byte, _ wasmvm.KVStore, _ wasmvm.GoAPI, _ wasmvm.Querier, _ wasmvm.GasMeter, _ uint64, _ wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error) {
					data, err := json.Marshal(types.EmptyResult{})
					suite.Require().NoError(err)

					return &wasmvmtypes.Response{Data: data}, wasmtesting.DefaultGasUsed, nil
				}
			},
			nil,
		},
		{
			"success: update client state",
			func() {
				msg = types.NewMsgMigrateContract(govAcc, defaultWasmClientID, newCodeHash, []byte("{}"))

				suite.mockVM.MigrateFn = func(_ wasmvm.Checksum, _ wasmvmtypes.Env, _ []byte, store wasmvm.KVStore, _ wasmvm.GoAPI, _ wasmvm.Querier, _ wasmvm.GasMeter, _ uint64, _ wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error) {
					// the code hash written in the client state will later be overwritten by the message server.
					expClientState = types.NewClientState([]byte{1}, []byte("invalid code hash"), clienttypes.NewHeight(2000, 2))
					store.Set(host.ClientStateKey(), clienttypes.MustMarshalClientState(suite.chainA.App.AppCodec(), expClientState))

					data, err := json.Marshal(types.EmptyResult{})
					suite.Require().NoError(err)

					return &wasmvmtypes.Response{Data: data}, wasmtesting.DefaultGasUsed, nil
				}
			},
			nil,
		},
		{
			"failure: same code hash",
			func() {
				msg = types.NewMsgMigrateContract(govAcc, defaultWasmClientID, oldCodeHash[:], []byte("{}"))

				suite.mockVM.MigrateFn = func(_ wasmvm.Checksum, _ wasmvmtypes.Env, _ []byte, _ wasmvm.KVStore, _ wasmvm.GoAPI, _ wasmvm.Querier, _ wasmvm.GasMeter, _ uint64, _ wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error) {
					panic("unreachable")
				}
			},
			types.ErrWasmCodeExists,
		},
		{
			"failure: unauthorized signer",
			func() {
				msg = types.NewMsgMigrateContract(suite.chainA.SenderAccount.GetAddress().String(), defaultWasmClientID, newCodeHash, []byte("{}"))
			},
			ibcerrors.ErrUnauthorized,
		},
		{
			"failure: invalid wasm code hash",
			func() {
				msg = types.NewMsgMigrateContract(govAcc, defaultWasmClientID, []byte(ibctesting.InvalidID), []byte("{}"))
			},
			types.ErrWasmCodeHashNotFound,
		},
		{
			"failure: invalid client id",
			func() {
				msg = types.NewMsgMigrateContract(govAcc, ibctesting.InvalidID, newCodeHash, []byte("{}"))
			},
			clienttypes.ErrClientTypeNotFound,
		},
		{
			"failure: contract returns error",
			func() {
				msg = types.NewMsgMigrateContract(govAcc, defaultWasmClientID, newCodeHash, []byte("{}"))

				suite.mockVM.MigrateFn = func(_ wasmvm.Checksum, _ wasmvmtypes.Env, _ []byte, _ wasmvm.KVStore, _ wasmvm.GoAPI, _ wasmvm.Querier, _ wasmvm.GasMeter, _ uint64, _ wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error) {
					return nil, wasmtesting.DefaultGasUsed, wasmtesting.ErrMockContract
				}
			},
			types.ErrWasmContractCallFailed,
		},
		{
			"failure: incorrect state update",
			func() {
				msg = types.NewMsgMigrateContract(govAcc, defaultWasmClientID, newCodeHash, []byte("{}"))

				suite.mockVM.MigrateFn = func(_ wasmvm.Checksum, _ wasmvmtypes.Env, _ []byte, store wasmvm.KVStore, _ wasmvm.GoAPI, _ wasmvm.Querier, _ wasmvm.GasMeter, _ uint64, _ wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error) {
					// the code hash written in here will be overwritten
					newClientState := localhost.NewClientState(clienttypes.NewHeight(1, 1))

					store.Set(host.ClientStateKey(), clienttypes.MustMarshalClientState(suite.chainA.App.AppCodec(), newClientState))

					data, err := json.Marshal(types.EmptyResult{})
					suite.Require().NoError(err)

					return &wasmvmtypes.Response{Data: data}, wasmtesting.DefaultGasUsed, nil
				}
			},
			types.ErrWasmInvalidContractModification,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupWasmWithMockVM()

			newCodeHash = storeWasmCode(suite, newByteCode)

			endpoint := wasmtesting.NewWasmEndpoint(suite.chainA)
			err := endpoint.CreateClient()
			suite.Require().NoError(err)

			// this is the old client state
			expClientState = endpoint.GetClientState().(*types.ClientState)

			tc.malleate()

			ctx := suite.chainA.GetContext()
			res, err := GetSimApp(suite.chainA).WasmClientKeeper.MigrateContract(ctx, msg)
			events := ctx.EventManager().Events().ToABCIEvents()

			if tc.expError == nil {
				expClientState.Checksum = newCodeHash

				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				// updated client state
				clientState, ok := endpoint.GetClientState().(*types.ClientState)
				suite.Require().True(ok)

				suite.Require().Equal(expClientState, clientState)

				// Verify events
				expectedEvents := sdk.Events{
					sdk.NewEvent(
						"migrate_contract",
						sdk.NewAttribute(types.AttributeKeyClientID, defaultWasmClientID),
						sdk.NewAttribute(types.AttributeKeyWasmCodeHash, hex.EncodeToString(oldCodeHash[:])),
						sdk.NewAttribute(types.AttributeKeyNewCodeHash, hex.EncodeToString(newCodeHash)),
					),
					sdk.NewEvent(
						sdk.EventTypeMessage,
						sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					),
				}.ToABCIEvents()

				for _, evt := range expectedEvents {
					suite.Require().Contains(events, evt)
				}
			} else {
				suite.Require().ErrorIs(err, tc.expError)
				suite.Require().Nil(res)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgRemoveCodeHash() {
	codeHash := sha256.Sum256(wasmtesting.Code)

	govAcc := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	var (
		msg           *types.MsgRemoveChecksum
		expCodeHashes []types.CodeHash
	)

	testCases := []struct {
		name     string
		malleate func()
		expError error
	}{
		{
			"success",
			func() {
				msg = types.NewMsgRemoveChecksum(govAcc, codeHash[:])

				expCodeHashes = []types.CodeHash{}
			},
			nil,
		},
		{
			"success: many code hashes",
			func() {
				msg = types.NewMsgRemoveChecksum(govAcc, codeHash[:])

				expCodeHashes = []types.CodeHash{}

				for i := 0; i < 20; i++ {
					codeHash := sha256.Sum256([]byte{byte(i)})
					err := ibcwasm.CodeHashes.Set(suite.chainA.GetContext(), codeHash[:])
					suite.Require().NoError(err)

					expCodeHashes = append(expCodeHashes, codeHash[:])
				}
			},
			nil,
		},
		{
			"failure: code hash is missing",
			func() {
				msg = types.NewMsgRemoveChecksum(govAcc, []byte{1})
			},
			types.ErrWasmCodeHashNotFound,
		},
		{
			"failure: unauthorized signer",
			func() {
				msg = types.NewMsgRemoveChecksum(suite.chainA.SenderAccount.GetAddress().String(), codeHash[:])
			},
			ibcerrors.ErrUnauthorized,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupWasmWithMockVM()

			endpoint := wasmtesting.NewWasmEndpoint(suite.chainA)
			err := endpoint.CreateClient()
			suite.Require().NoError(err)

			tc.malleate()

			ctx := suite.chainA.GetContext()
			res, err := GetSimApp(suite.chainA).WasmClientKeeper.RemoveChecksum(ctx, msg)
			events := ctx.EventManager().Events().ToABCIEvents()

			if tc.expError == nil {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				codeHashes, err := types.GetAllCodeHashes(suite.chainA.GetContext())
				suite.Require().NoError(err)

				// Check equality of code hashes up to order
				suite.Require().ElementsMatch(expCodeHashes, codeHashes)

				// Verify events
				suite.Require().Len(events, 0)
			} else {
				suite.Require().ErrorIs(err, tc.expError)
				suite.Require().Nil(res)
			}
		})
	}
}
