package ibcwasm

import (
	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
)

var (
	vm WasmEngine

	// state management
	Schema     collections.Schema
	CodeHashes collections.KeySet[[]byte]

	// ChecksumsKey is the key under which all checksums are stored
	ChecksumsKey = collections.NewPrefix(0)
)

// SetVM sets the wasm VM for the 08-wasm module.
func SetVM(wasmVM WasmEngine) {
	vm = wasmVM
}

// GetVM returns the wasm VM for the 08-wasm module.
func GetVM() WasmEngine {
	return vm
}

// SetupWasmStoreService sets up the 08-wasm module's collections.
func SetupWasmStoreService(storeService storetypes.KVStoreService) {
	sb := collections.NewSchemaBuilder(storeService)

	CodeHashes = collections.NewKeySet(sb, ChecksumsKey, "checksums", collections.BytesKey)

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	Schema = schema
}
