package contract

import (
	"crypto/sha256"

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
	name     sdk.String
	owner    sdk.Address
	balances map[sdk.Address]sdk.U256
}

// abi: sayName() returns (string) view
func (c *FooContract) SayName() (sdk.String, error) {
	return c.name, nil
}

// abi:  transfer(address,uint256)
func (c *FooContract) Transfer(to sdk.Address, v sdk.U256) error {
	ownerS := sdk.SLoad(sdk.FromUInt64(3))
	owner := sdk.AddressFromBytes(ownerS.Bytes())
	if owner != sdk.MsgSender() {
		return sdk.EvmError("not owner")
	}
	res, err := to.Call(sdk.WithMaxGas(), sdk.WithValue(sdk.FromUInt64(123)))
	if err != nil {
		return err
	}
	sdk.LogRawN(res, 0)
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
func (c *FooContract) TransferFrom(from sdk.Address, to sdk.Address, v sdk.U256) (sdk.U256, error) {
	hasher := sha256.New()
	hasher.Write(from.Bytes())
	res := hasher.Sum(nil)
	return sdk.FromBytes(res), nil
}
