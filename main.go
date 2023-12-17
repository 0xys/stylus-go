package main

// #include "./hostio.h"
import "C"

/*
type Address [20]uint8

const evmWordSize = 256

func slice(p *C.uint8_t, length int) []uint8 {
	return ((*[evmWordSize]uint8)(unsafe.Pointer(p)))[0:length:length]
}

func MemoryGrow(pages uint16) {
	p := C.uint16_t(pages)
	C.memory_grow(p)
}

func InkLeft() uint64 {
	return C.evm_ink_left()
}

func GasLeft() uint64 {
	return C.evm_gas_left()
}

func MsgSender() Address {
	var src *C.uint8_t
	C.msg_sender(src)
	var ret Address
	array := slice(src, 20)
	for i := 0; i < 20; i++ {
		ret[i] = array[i]
	}
	return ret
}

func MsgValue() uint256.Int {
	var src *C.uint8_t
	C.msg_value(src)
	array := slice(src, 32)
	ret := &uint256.Int{}
	ret.SetBytes(array)
	return *ret
}
*/

//go:wasm-module vm_hooks
//export chainid
func chainid() uint64

//go:wasm-module vm_hooks
//export memory_grow
func memory_grow(pages uint16)

//export user_entrypoint
func user_entrypoint(args_len uint32) uint32 {
	memory_grow(0)
	chainid()
	return 0
}

func main() {
	user_entrypoint(1)
}

/*
func main() {
	g := GasLeft()
	i := InkLeft()
	MemoryGrow(123)
	_ = g
	_ = i
	MsgSender()
}
*/
