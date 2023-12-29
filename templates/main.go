package template

// abi: hello() uint256
func Hello() int {
	return 0
}

// World func is public
// hello
func World(str string) string {
	return "world"
}

func privateSomeFunc(aaa int, b uint64) uint64 {
	return uint64(aaa) + b + 10
}

type Address [20]byte
type U256 [4]uint64

// contract
type FooContract struct {
	value    int
	name     string
	owner    Address
	balances map[Address]U256
}

// abi: sayHi()
func (c *FooContract) SayHi() error {
}

// abi:  transfer(address,uint256)
func (c *FooContract) Transfer(to Address, v U256) error {
}

// abi: receive(uint64) payable
func (c *FooContract) Receive(a uint64) error {
}

// abi: transferFrom(address,address,uint256) returns (bool)
func (c *FooContract) TransferFrom(from, to Address, v U256) (bool, error) {
	return false
}
