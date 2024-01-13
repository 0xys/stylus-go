package sdk

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
//export contract_address
func contract_address(address *uint8)

//go:wasm-module vm_hooks
//export create1
func create1(code *uint8, code_len uint32, endowment *uint8, contract *uint8, return_data_len *uint32)

//go:wasm-module vm_hooks
//export create2
func create2(code *uint8, code_len uint32, endowment *uint8, salt *uint8, contract *uint8, return_data_len *uint32)

//go:wasm-module vm_hooks
//export delegate_call_contract
func delegate_call_contract(contract *uint8, calldata *uint8, calldata_len uint32, gas uint64, return_data_len *uint32) uint8

//go:wasm-module vm_hooks
//export emit_log
func emit_log(data *uint8, length uint32, topics uint32)

//go:wasm-module vm_hooks
//export evm_gas_left
func evm_gas_left() uint64

//go:wasm-module vm_hooks
//export evm_ink_left
func evm_ink_left() uint64

//go:wasm-module vm_hooks
//export memory_grow
func memory_grow(pages uint16)

//go:wasm-module vm_hooks
//export msg_sender
func msg_sender(sender *uint8)

//go:wasm-module vm_hooks
//export msg_value
func msg_value(value *uint8)

//go:wasm-module vm_hooks
//export native_keccak256
func native_keccak256(bytes *uint8, length uint32, output *uint8)

//go:wasm-module vm_hooks
//export read_args
func read_args(data *uint8)

//go:wasm-module vm_hooks
//export read_return_data
func read_return_data(dest *uint8, offset uint32, size uint32) uint32

//go:wasm-module vm_hooks
//export write_result
func write_result(data *uint8, length uint32)

//go:wasm-module vm_hooks
//export return_data_size
func return_data_size() uint32

//go:wasm-module vm_hooks
//export static_call_contract
func static_call_contract(contract *uint8, calldata *uint8, calldata_len uint32, gas uint64, return_data_len *uint32) uint8

//go:wasm-module vm_hooks
//export tx_gas_price
func tx_gas_price(gas_price *uint8)

//go:wasm-module vm_hooks
//export tx_ink_price
func tx_ink_price() uint32

//go:wasm-module vm_hooks
//export tx_origin
func tx_origin(origin *uint8)
