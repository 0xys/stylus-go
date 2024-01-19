package sdk

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
