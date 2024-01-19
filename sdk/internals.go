package sdk

var (
	isPure bool
	isView bool
)

func SetPure() {
	isPure = true
}

func SetView() {
	isView = true
}

func Init(argLen uint32) {
	calldataLen = argLen
	returnStatus = 0
	memory_grow(0) // memory grow must be included in wasm binary
}

func Flush() {
}
