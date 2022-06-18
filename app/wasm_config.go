package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

const (
	// DefaultRXPInstanceCost is initially set the same as in wasmd
	DefaultRXPInstanceCost uint64 = 60_000
	// DefaultRXPCompileCost set to a large number for testing
	DefaultRXPCompileCost uint64 = 100
)

// RXPGasRegisterConfig is defaults plus a custom compile amount
func RXPGasRegisterConfig() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultRXPInstanceCost
	gasConfig.CompileCost = DefaultRXPCompileCost

	return gasConfig
}

func NewRXPWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(RXPGasRegisterConfig())
}
