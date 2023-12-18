package main

//go:wasm-module vm_hooks
//export account_balance
func account_balance(addr *uint8, dest *uint8)

//go:wasm-module vm_hooks
//export account_codehash
func account_codehash(addr *uint8, dest *uint8)

//go:wasm-module vm_hooks
//export storage_load_bytes32
func storage_load_bytes32(key *uint8, dest *uint8)

//go:wasm-module vm_hooks
//export storage_store_bytes32
func storage_store_bytes32(key *uint8, value *uint8)

//go:wasm-module vm_hooks
//export block_basefee
func block_basefee(basefee *uint8)

//go:wasm-module vm_hooks
//export chainid
func chainid() uint64

//go:wasm-module vm_hooks
//export block_coinbase
func block_coinbase(coinbase *uint8)

//go:wasm-module vm_hooks
//export block_gas_limit
func block_gas_limit() uint64

//go:wasm-module vm_hooks
//export block_number
func block_number() uint64

//go:wasm-module vm_hooks
//export block_timestamp
func block_timestamp() uint64

//go:wasm-module vm_hooks
//export call_contract
func call_contract(contract *uint8, calldata *uint8, calldata_len uint32, value *uint8, gas uint64, return_data_len *uint32) uint8

//go:wasm-module vm_hooks
//export memory_grow
func memory_grow(pages uint16)

//export user_entrypoint
func user_entrypoint(args_len uint32) uint32 {
	addr := []uint8{0}
	dest := []uint8{0}
	val := []uint8{0}
	ll := []uint32{0}
	account_balance(&addr[0], &dest[0])
	account_codehash(&addr[0], &dest[0])
	storage_load_bytes32(&addr[0], &dest[0])
	storage_store_bytes32(&addr[0], &dest[0])
	block_basefee(&val[0])
	chainid()
	block_coinbase(&addr[0])
	block_gas_limit()
	block_number()
	block_timestamp()
	call_contract(&addr[0], &dest[0], uint32(10), &val[0], uint64(11), &ll[0])
	memory_grow(0)

	return 0
}

func main() {
	user_entrypoint(1)
}
