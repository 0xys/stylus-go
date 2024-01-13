package sdk

func Init(argLen uint32) {
	calldataLen = argLen
	returnStatus = 0
	memory_grow(0) // memory grow must be included in wasm binary
}

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

// TODO
func Create1() {
}
func Create2() {
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
func LogAddress(addr Address, topics uint32) {
	LogRawN(addr.Bytes(), topics)
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
