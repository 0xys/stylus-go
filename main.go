package main

import "github.com/0xys/stylus-go/sdk"

func testReturn() uint32 {
	sdk.SetReturnBytes([]byte{0x32, 0x64, 0x1})
	return 0
}

const (
	StoreMarker uint8 = 0x01
	LoadMarker  uint8 = 0x02
	Load2Marker uint8 = 0x03 // load with log

	CallMarker         uint8 = 0x0a
	DelegateCallMarker uint8 = 0x0b
	StaticCallMarker   uint8 = 0x0c
)

func testStorage() uint32 {
	cd := sdk.GetCalldata()
	defaultKey := 0
	if len(cd) < 1 {
		sdk.RevertWithString("not enough calldata")
	}
	switch cd[0] {
	case StoreMarker:
		v := sdk.FromBytes(cd[1:])
		sdk.LogU256(v)
		sdk.SStore(sdk.FromUInt64(uint64(defaultKey)), v)
	case LoadMarker:
		val := sdk.SLoad(sdk.FromUInt64(uint64(defaultKey)))
		sdk.SetReturnU256(val)
	case Load2Marker:
		val := sdk.SLoad(sdk.FromUInt64(uint64(defaultKey)))
		sdk.LogU256(val)
		sdk.SetReturnU256(val)
	}
	return 0

}
func testCall() uint32 {
	cd := sdk.GetCalldata()
	if len(cd) < 1 {
		sdk.RevertWithString("not enough calldata")
	}
	addr := sdk.AddressFromBytes(cd[1:21])
	sdk.LogUInt8(cd[0], 0)
	sdk.LogAddress(addr, 0)
	defaultKey := 0
	switch cd[0] {
	case StoreMarker:
		v := sdk.FromBytes(cd[1:])
		sdk.LogU256(v)
		sdk.SStore(sdk.FromUInt64(uint64(defaultKey)), v)
	case LoadMarker:
		val := sdk.SLoad(sdk.FromUInt64(uint64(defaultKey)))
		sdk.SetReturnU256(val)
	case Load2Marker:
		val := sdk.SLoad(sdk.FromUInt64(uint64(defaultKey)))
		sdk.LogU256(val)
		sdk.SetReturnU256(val)
	case CallMarker:
		res, err := addr.Call(sdk.WithCalldata(cd[21:]))
		if err != nil {
			sdk.Log("fail call")
			return 0
		}
		if len(res) > 0 {
			sdk.LogRawN(res, 0)
		}
	case DelegateCallMarker:
		res, err := addr.DelegateCall(sdk.WithCalldata(cd[21:]))
		if err != nil {
			sdk.Log("fail delegate call")
			return 0
		}
		if len(res) > 0 {
			sdk.LogRawN(res, 0)
		}
	case StaticCallMarker:
		res, err := addr.StaticCall(sdk.WithCalldata(cd[21:]))
		if err != nil {
			sdk.Log("fail static call")
			return 0
		}
		if len(res) > 0 {
			sdk.LogRawN(res, 0)
		}
	}
	return 0
}
func testPanic() uint32 {
	sdk.Revert([]byte{0, 1, 2, 0x43, 0x88})
	return 0
}

//export user_entrypoint
func user_entrypoint(args_len uint32) uint32 {
	sdk.Init(args_len)
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
	return testCall()
}

func main() {
	user_entrypoint(1)
}
