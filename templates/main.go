package template

import (
	"github.com/0xys/stylus-go/sdk"
)

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

// contract
type FooContract struct {
	value    int
	name     string
	owner    sdk.Address
	balances map[sdk.Address]sdk.U256
}

// abi: sayHi() pure
func (c *FooContract) SayHi() error {
	sdk.Log("hi")
	return nil
}

// abi:  transfer(address,uint256)
func (c *FooContract) Transfer(to sdk.Address, v sdk.U256) error {
	return nil
}

// abi: receive(uint64) payable
func (c *FooContract) Receive(a uint64) error {
	return nil
}

// abi: balance() view
func (c *FooContract) Balance() (sdk.U256, error) {
	return sdk.U256([4]uint64{0}), nil
}

// abi: transferFrom(address,address,uint256) returns (bool)
func (c *FooContract) TransferFrom(from, to sdk.Address, v sdk.U256) (bool, error) {
	return false, nil
}
