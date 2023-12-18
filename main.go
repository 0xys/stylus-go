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

type Bytes []uint8

const AddressLen = 20

type Address [AddressLen]uint8

func (a Address) String() string {
	//return hex.EncodeToString([]byte(a[:]))
	return "todo"
}

func GetCalldata(length uint32) Bytes {
	dest := make([]uint8, length, length)
	read_args(&dest[0])
	return Bytes(dest)
}

func TxOrigin() Address {
	addr := make([]uint8, AddressLen, AddressLen)
	tx_origin(&addr[0])
	return Address(addr)
}

func MsgSender() Address {
	addr := make([]uint8, AddressLen, AddressLen)
	msg_sender(&addr[0])
	return Address(addr)
}

func Log(msg string) {
	LogN(msg, 0)
}

func LogN(msg string, topics uint32) {
	bytes := []byte(msg)
	length := len(bytes)
	emit_log(&bytes[0], uint32(length), topics)
}

//export user_entrypoint
func user_entrypoint(args_len uint32) uint32 {
	addr := []uint8{0}
	dest := []uint8{0}
	val := []uint8{0}
	ll := []uint32{0}
	account_balance(&addr[0], &dest[0])
	account_codehash(&addr[0], &dest[0])

	GetCalldata(args_len)
	TxOrigin()
	Log("hello world")

	storage_load_bytes32(&addr[0], &dest[0])
	storage_store_bytes32(&addr[0], &dest[0])

	block_basefee(&val[0])

	chainid()

	block_coinbase(&addr[0])
	block_gas_limit()
	block_number()
	block_timestamp()

	call_contract(&addr[0], &dest[0], uint32(10), &val[0], uint64(11), &ll[0])
	contract_address(&addr[0])
	create1(&val[0], uint32(10), &addr[0], &addr[0], &ll[0])
	create2(&val[0], uint32(10), &addr[0], &val[0], &addr[0], &ll[0])
	delegate_call_contract(&addr[0], &dest[0], uint32(10), uint64(11), &ll[0])

	emit_log(&val[0], uint32(10), uint32(11))

	evm_gas_left()
	evm_ink_left()

	memory_grow(0)

	msg_sender(&addr[0])
	msg_value(&val[0])

	native_keccak256(&addr[0], uint32(10), &dest[0])

	read_args(&val[0])
	read_return_data(&dest[0], uint32(10), uint32(32))
	write_result(&dest[0], uint32(10))
	return_data_size()
	static_call_contract(&addr[0], &dest[0], uint32(10), uint64(11), &ll[0])
	tx_gas_price(&val[0])
	tx_ink_price()
	tx_origin(&addr[0])

	return 0
}

func main() {
	user_entrypoint(1)
}
