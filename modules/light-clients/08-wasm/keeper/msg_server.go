package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/internal/ibcwasm"
	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"
	ibcerrors "github.com/cosmos/ibc-go/v8/modules/core/errors"
)

var _ types.MsgServer = (*Keeper)(nil)

// StoreCode defines a rpc handler method for MsgStoreCode
func (k Keeper) StoreCode(goCtx context.Context, msg *types.MsgStoreCode) (*types.MsgStoreCodeResponse, error) {
	if k.GetAuthority() != msg.Signer {
		return nil, errorsmod.Wrapf(ibcerrors.ErrUnauthorized, "expected %s, got %s", k.GetAuthority(), msg.Signer)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	codeHash, err := k.storeWasmCode(ctx, msg.WasmByteCode)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to store wasm bytecode")
	}

	emitStoreWasmCodeEvent(ctx, codeHash)

	return &types.MsgStoreCodeResponse{
		Checksum: codeHash,
	}, nil
}

// RemoveChecksum defines a rpc handler method for MsgRemoveChecksum
func (k Keeper) RemoveChecksum(goCtx context.Context, msg *types.MsgRemoveChecksum) (*types.MsgRemoveChecksumResponse, error) {
	if k.GetAuthority() != msg.Signer {
		return nil, errorsmod.Wrapf(ibcerrors.ErrUnauthorized, "expected %s, got %s", k.GetAuthority(), msg.Signer)
	}

	if !types.HasCodeHash(goCtx, msg.Checksum) {
		return nil, types.ErrWasmCodeHashNotFound
	}

	err := ibcwasm.CodeHashes.Remove(goCtx, msg.Checksum)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to remove code hash")
	}

	return &types.MsgRemoveChecksumResponse{}, nil
}

// MigrateContract defines a rpc handler method for MsgMigrateContract
func (k Keeper) MigrateContract(goCtx context.Context, msg *types.MsgMigrateContract) (*types.MsgMigrateContractResponse, error) {
	if k.GetAuthority() != msg.Signer {
		return nil, errorsmod.Wrapf(ibcerrors.ErrUnauthorized, "expected %s, got %s", k.GetAuthority(), msg.Signer)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.migrateContractCode(ctx, msg.ClientId, msg.Checksum, msg.Msg)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to migrate contract")
	}

	// event emission is handled in migrateContractCode

	return &types.MsgMigrateContractResponse{}, nil
}
