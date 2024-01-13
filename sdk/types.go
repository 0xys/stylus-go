package sdk

import "math/bits"

type Word [32]uint8

func Zero() Word {
	return [32]uint8{0}
}

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

func AddressFromBytes(bytes Bytes) Address {
	ret := [AddressLen]uint8{0}
	rightMostIndex := min(AddressLen, len(bytes)) - 1
	for i := 0; i <= rightMostIndex; i++ {
		ret[AddressLen-1-i] = bytes[rightMostIndex-i]
	}
	return ret
}

func (a Address) String() string {
	dst := make([]byte, EncodedLen(len(a[:])))
	Encode(dst, a[:])
	return string(dst)
}

func (a Address) Bytes() Bytes {
	return a[:]
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

const uintMax = 18446744073709551615 // uint.max()
func WithMaxGas() func(*callOpt) {
	return func(c *callOpt) {
		c.gas = uintMax
	}
}

func (a Address) Call(opts ...func(*callOpt)) (Bytes, error) {
	zero := Zero()
	opt := &callOpt{
		gas:   uintMax,
		value: zero[:], // value must be filled with zeros. TODO: optimize
	}
	for _, o := range opts {
		o(opt)
	}
	var retDataLen uint32
	status := call_contract(&a[0], &opt.calldata[0], uint32(len(opt.calldata)), &opt.value[0], opt.gas, &retDataLen)
	ret := make([]uint8, retDataLen, retDataLen)
	if retDataLen > 0 {
		read_return_data(&ret[0], 0, retDataLen)
	}
	if status != 0 {
		return ret, callerror
	}
	return ret, nil
}

func (a Address) StaticCall(opts ...func(*callOpt)) (Bytes, error) {
	zero := Zero()
	opt := &callOpt{
		gas:   uintMax,
		value: zero[:],
	}
	for _, o := range opts {
		o(opt)
	}
	var retDataLen uint32
	status := static_call_contract(&a[0], &opt.calldata[0], uint32(len(opt.calldata)), opt.gas, &retDataLen)
	ret := make([]uint8, retDataLen, retDataLen)
	if retDataLen > 0 {
		read_return_data(&ret[0], 0, retDataLen)
	}
	if status != 0 {
		return ret, callerror
	}
	return ret, nil
}

func (a Address) DelegateCall(opts ...func(*callOpt)) (Bytes, error) {
	zero := Zero()
	opt := &callOpt{
		gas:   uintMax,
		value: zero[:],
	}
	for _, o := range opts {
		o(opt)
	}
	var retDataLen uint32
	status := delegate_call_contract(&a[0], &opt.calldata[0], uint32(len(opt.calldata)), opt.gas, &retDataLen)
	ret := make([]uint8, retDataLen, retDataLen)
	if retDataLen > 0 {
		read_return_data(&ret[0], 0, retDataLen)
	}
	if status != 0 {
		return ret, callerror
	}
	return ret, nil
}
