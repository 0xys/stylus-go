package main

//go:wasm-module vm_hooks
//export account_balance
func account_balance(addr *uint8, dest *uint8)

//go:wasm-module vm_hooks
//export chainid
func chainid() uint64

//go:wasm-module vm_hooks
//export memory_grow
func memory_grow(pages uint16)

//go:wasm-module vm_hooks
//export memory_grow

//export user_entrypoint
func user_entrypoint(args_len uint32) uint32 {
	memory_grow(0)
	chainid()
	addr := []uint8{}
	dest := []uint8{}
	account_balance(&addr[0], &dest[0])
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
