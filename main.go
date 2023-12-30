package main

import "math/bits"

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

func (b Bytes) String() string {
	dst := make([]byte, EncodedLen(len(b)))
	Encode(dst, b)
	return string(dst)
}

func Concat(items ...Bytes) Bytes {
	ret := make([]uint8, 0, 1024)
	for _, item := range items {
		ret = append(ret, item...)
	}
	return ret
}

const AddressLen = 20

type Address [AddressLen]uint8

func (a Address) String() string {
	dst := make([]byte, EncodedLen(len(a[:])))
	Encode(dst, a[:])
	return string(dst)
}

func (a Address) Balance() Bytes {
	bal := make([]byte, 32, 32)
	account_balance(&a[0], &bal[0])
	return bal
}
func (a Address) CodeHash() Bytes {
	hash := make([]byte, 32, 32)
	account_codehash(&a[0], &hash[0])
	return hash
}

func ContractAddress() Address {
	ret := [AddressLen]uint8{0}
	contract_address(&ret[0])
	return ret
}

type EvmError string

const callerror EvmError = "e"

func NewEvmError(e string) EvmError {
	return EvmError(e)
}

func (e EvmError) Error() string {
	return string(e)
}

type callOpt struct {
	calldata Bytes
	value    Bytes
	gas      uint64
}

func WithCalldata(calldata Bytes) func(*callOpt) {
	return func(c *callOpt) {
		c.calldata = calldata
	}
}
func WithValue(v U256) func(*callOpt) {
	return func(c *callOpt) {
		w := v.Word()
		c.value = w[:]
	}
}
func WithGas(gas uint64) func(*callOpt) {
	return func(c *callOpt) {
		c.gas = gas
	}
}
func WithMaxGas() func(*callOpt) {
	return func(c *callOpt) {
		c.gas = 18446744073709551615 // uint.max()
	}
}

func (a Address) Call(opts ...func(*callOpt)) (Bytes, error) {
	opt := &callOpt{}
	for _, o := range opts {
		o(opt)
	}
	retDataLen := uint32(0)
	status := call_contract(&a[0], &opt.calldata[0], uint32(len(opt.calldata)), &opt.value[0], opt.gas, &retDataLen)

	ret := make([]uint8, retDataLen, retDataLen)
	read_return_data(&ret[0], 0, retDataLen)
	if status != 1 {
		return ret, callerror
	}
	return ret, nil
}

func (a Address) StaticCall(opts ...func(*callOpt)) (Bytes, error) {
	opt := &callOpt{}
	for _, o := range opts {
		o(opt)
	}
	retDataLen := uint32(0)
	status := static_call_contract(&a[0], &opt.calldata[0], uint32(len(opt.calldata)), opt.gas, &retDataLen)
	ret := make([]uint8, retDataLen, retDataLen)
	read_return_data(&ret[0], 0, retDataLen)
	if status != 1 {
		return ret, callerror
	}
	return ret, nil
}

func (a Address) DelegateCall(opts ...func(*callOpt)) (Bytes, error) {
	opt := &callOpt{}
	for _, o := range opts {
		o(opt)
	}
	retDataLen := uint32(0)
	status := delegate_call_contract(&a[0], &opt.calldata[0], uint32(len(opt.calldata)), opt.gas, &retDataLen)
	ret := make([]uint8, retDataLen, retDataLen)
	read_return_data(&ret[0], 0, retDataLen)
	if status != 1 {
		return ret, callerror
	}
	return ret, nil
}

// TODO
func Create1() {
}
func Create2() {
}

// auto generate from ABI vvvv
/**
name: SomeContract
methods:
	uint256 transfer(address, uint)


type SomeContract struct {
	addr Address
}

func NewSomeContract(addr Address) *SomeContract {
	return &SomeContract{addr}
}

func (c *SomeContract) Transfer(addr Address, val U256) (U256, error) {
	tag := []uint8{0x01, 0x02, 0x01, 0x02, 0x01, 0x02, 0x01, 0x02}
	arg1 := addr.Encode()
	arg2 := val.Encode()
	res, err := c.addr.Call(WithCalldata(Concat(tag, arg1, arg2)), WithMaxGas())
	ret := FromBytes(res)
	return ret, err
}

// auto generate ^^^^

func Sample() {
	var usdtAddr Address
	addr := ContractAddress()
	usdt := FromAddress(usdtAddr)
	usdt.Transfer(addr, FromUInt64(100))
}
*/

func Keccak256(data Bytes) Word {
	ret := [32]uint8{0}
	native_keccak256(&data[0], uint32(len(data)), &ret[0])
	return ret
}

func BlockBaseFee() Bytes {
	bal := make([]byte, 32, 32)
	block_basefee(&bal[0])
	return bal
}
func ChainID() uint64 {
	return chainid()
}
func BlockCoinbase() Address {
	addr := make([]uint8, AddressLen, AddressLen)
	block_coinbase(&addr[0])
	return Address(addr)
}
func BlockGasLimit() uint64 {
	return block_gas_limit()
}
func BlockNumber() uint64 {
	return block_number()
}
func BlockTimestamp() uint64 {
	return block_timestamp()
}
func GasLeft() uint64 {
	return evm_gas_left()
}
func InkLeft() uint64 {
	return evm_ink_left()
}
func GasPrice() U256 {
	p := make([]uint8, 32, 32)
	tx_gas_price(&p[0])
	return FromBytes(p)
}
func InkPrice() uint32 {
	return tx_ink_price()
}
func ReturnDataSize() uint32 {
	return return_data_size()
}
func ReturnData() Bytes {
	sz := ReturnDataSize()
	ret := make([]uint8, sz, sz)
	read_return_data(&ret[0], 0, sz)
	return ret
}

// https://go.googlesource.com/go/+/go1.9.2/src/encoding/hex/hex.go
const hextable = "0123456789abcdef"

// EncodedLen returns the length of an encoding of n source bytes.
// Specifically, it returns n * 2.
func EncodedLen(n int) int { return n * 2 }

// Encode encodes src into EncodedLen(len(src))
// bytes of dst. As a convenience, it returns the number
// of bytes written to dst, but this value is always EncodedLen(len(src)).
// Encode implements hexadecimal encoding.
func Encode(dst, src []byte) int {
	for i, v := range src {
		dst[i*2] = hextable[v>>4]
		dst[i*2+1] = hextable[v&0x0f]
	}
	return len(src) * 2
}

func ToSelector(in []byte) uint32 {
	ret := uint32(0)
	ret += uint32(in[3])
	ret += uint32(in[2]) << 8
	ret += uint32(in[1]) << 16
	ret += uint32(in[0]) << 24
	return ret
}

// must be initialized in entrypoint
var calldataLen uint32

func GetCalldata() Bytes {
	dest := make([]uint8, calldataLen, calldataLen)
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
func MsgValue() U256 {
	w := [32]uint8{0}
	msg_value(&w[0])
	return FromWord(w)
}

func Log(msg string) {
	LogN(msg, 0)
}

func LogN(msg string, topics uint32) {
	bytes := []byte(msg)
	LogRawN(bytes, topics)
}
func LogUInt8(n uint8, topics uint32) {
	bytes := []byte{n}
	LogRawN(bytes, topics)
}
func LogUInt32(n uint32, topics uint32) {
	b := []byte{0, 0, 0, 0}
	b[0] = byte(n >> 24)
	b[1] = byte(n >> 16)
	b[2] = byte(n >> 8)
	b[3] = byte(n)
	LogRawN(b, topics)
}
func LogRawN(bytes Bytes, topics uint32) {
	length := len(bytes)
	emit_log(&bytes[0], uint32(length), topics)
}
func LogU256(v U256) {
	w := v.Word()
	emit_log(&w[0], uint32(32), 0)
}

var returnStatus uint32

func SetReturnBytes(bytes Bytes) {
	write_result(&bytes[0], uint32(len(bytes)))
}
func SetReturnU256(z U256) {
	bytes := z.Word()
	write_result(&bytes[0], uint32(len(bytes)))
}
func SetReturnUInt8(val uint8) {
	bytes := []uint8{val}
	SetReturnBytes(bytes)
}
func SetReturnUInt64(val uint64) {
	SetReturnU256(FromUInt64(val))
}
func SetReturnString(msg string) {
	SetReturnBytes([]byte(msg))
}
func SetReturnAddress(addr Address) {
	SetReturnBytes(addr[:])
}

func Revert(bytes Bytes) {
	SetReturnBytes(bytes)
	panic(0)
}
func RevertWithString(msg string) {
	Revert([]byte(msg))
}

func uint64ToBytes(v uint64) [8]uint8 {
	b := [8]uint8{0, 0, 0, 0, 0, 0, 0, 0}
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// big-endian
func bytesToUint64(b []uint8) uint64 {
	ret := uint64(0)
	lsbIdx := min(len(b), 8) - 1
	for i := lsbIdx; i >= 0; i-- {
		ret = ret + (uint64(b[i]) << (8 * (lsbIdx - i)))
	}
	return ret
}

func SStore(key, value U256) {
	k := key.Word()
	v := value.Word()
	storage_store_bytes32(&k[0], &v[0])
}
func SLoad(key U256) U256 {
	ret := make([]uint8, 32, 32)
	k := key.Word()
	storage_load_bytes32(&k[0], &ret[0])
	return FromBytes(ret)
}

type Word [32]uint8

type U256 [4]uint64

func NewU256() U256 {
	return [4]uint64{0, 0, 0, 0}
}

func FromUInt64(v uint64) U256 {
	return [4]uint64{v, 0, 0, 0}
}
func FromWord(w Word) U256 {
	ret := NewU256()
	ret[0] = bytesToUint64(w[24:32])
	ret[1] = bytesToUint64(w[16:24])
	ret[2] = bytesToUint64(w[8:16])
	ret[3] = bytesToUint64(w[:8])
	return ret
}
func FromBytes(b Bytes) U256 {
	ret := NewU256()
	l := len(b)
	if l == 0 {
		return ret
	}
	if l <= 8 {
		ret[0] = bytesToUint64(b)
		return ret
	}
	if l <= 16 {
		ret[0] = bytesToUint64(b[l-8 : l])
		ret[1] = bytesToUint64(b[:l-8])
		return ret
	}
	if l <= 24 {
		ret[0] = bytesToUint64(b[l-8 : l])
		ret[1] = bytesToUint64(b[l-16 : l-8])
		ret[2] = bytesToUint64(b[:l-16])
		return ret
	}
	if l <= 32 {
		ret[0] = bytesToUint64(b[l-8 : l])
		ret[1] = bytesToUint64(b[l-16 : l-8])
		ret[2] = bytesToUint64(b[l-24 : l-16])
		ret[3] = bytesToUint64(b[:l-24])
		return ret
	}

	ret[0] = bytesToUint64(b[24:32])
	ret[1] = bytesToUint64(b[16:24])
	ret[2] = bytesToUint64(b[8:16])
	ret[3] = bytesToUint64(b[:8])
	return ret
}

func (z *U256) IsZero() bool {
	return z[0] == uint64(0) && z[1] == uint64(0) && z[2] == uint64(0) && z[3] == uint64(0)
}

func (z *U256) Add(x, y *U256) *U256 {
	var carry uint64
	z[0], carry = bits.Add64(x[0], y[0], 0)
	z[1], carry = bits.Add64(x[1], y[1], carry)
	z[2], carry = bits.Add64(x[2], y[2], carry)
	z[3], _ = bits.Add64(x[3], y[3], carry)
	return z
}

func (z *U256) Word() Word {
	dst := [32]uint8{0}
	b1 := uint64ToBytes(z[0])
	b2 := uint64ToBytes(z[1])
	b3 := uint64ToBytes(z[2])
	b4 := uint64ToBytes(z[3])
	copy(dst[:8], b4[:8])
	copy(dst[8:16], b3[:8])
	copy(dst[16:24], b2[:8])
	copy(dst[24:32], b1[:8])
	return dst
}

func testReturn() uint32 {
	SetReturnBytes([]byte{0x32, 0x64, 0x1})
	return 0
}

const (
	StoreMarker uint8 = 0x01
	LoadMarker  uint8 = 0x02
	Load2Marker uint8 = 0x03 // load with log
)

func testStorage() uint32 {
	cd := GetCalldata()
	defaultKey := 0
	if len(cd) < 1 {
		RevertWithString("not enough calldata")
	}
	switch cd[0] {
	case StoreMarker:
		v := FromBytes(cd[1:])
		LogU256(v)
		SStore(FromUInt64(uint64(defaultKey)), v)
	case LoadMarker:
		val := SLoad(FromUInt64(uint64(defaultKey)))
		SetReturnU256(val)
	case Load2Marker:
		val := SLoad(FromUInt64(uint64(defaultKey)))
		LogU256(val)
		SetReturnU256(val)
	}
	return 0

}
func testCall() uint32 {
	return 0
}
func testPanic() uint32 {
	Revert([]byte{0, 1, 2, 0x43, 0x88})
	return 0
}

//export user_entrypoint
func user_entrypoint(args_len uint32) uint32 {
	returnStatus = 0
	calldataLen = args_len
	// LogUInt32(args_len, 0)
	// Log(TxOrigin().String())
	// calldata := GetCalldata()
	// LogRawN(calldata, 0)

	// addr := Address(calldata[:20])
	// LogRawN(addr[:], 0)

	// LogRawN(addr.Balance(), 0)
	// LogRawN(MsgSender().Balance(), 0)

	// val := MsgValue()
	// LogU256(val)

	// bn := FromUInt64(BlockNumber())
	// LogU256(bn)

	// bs := FromUInt64(BlockTimestamp())
	// LogU256(bs)

	// sum := NewU256()
	// LogU256(*sum.Add(&bn, &bs))

	/*
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
	*/
	return testStorage()
}

func main() {
	memory_grow(0)
	user_entrypoint(1)

}
