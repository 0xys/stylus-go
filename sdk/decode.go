package sdk

func DecodeAddress(data Bytes) (Address, error) {
	if len(data) < AddressLen {
		return NilAddress, EvmError("not address")
	}
	ret := [AddressLen]uint8{0}
	for i := 0; i <= AddressLen; i++ {
		ret[AddressLen-1-i] = data[len(data)-i]
	}
	return ret, nil
}

func DecodeU256(data Bytes) (U256, error) {
	if len(data) < 32 {
		return Zero, EvmError("not u256")
	}
	return FromBytes(data), nil
}

func DecodeUint64(data Bytes) (uint64, error) {
	if len(data) < 8 {
		return 0, EvmError("not uint64")
	}
	ret := uint64(0)
	lsbIdx := len(data) - 1
	for i := lsbIdx; i >= 0; i-- {
		ret = ret + (uint64(data[i]) << (8 * (lsbIdx - i)))
	}
	return ret, nil
}

func DecodeBool(data Bytes) (bool, error) {
	if len(data) < 1 {
		return false, EvmError("not bool")
	}
	if data[len(data)-1] == 1 {
		return true, nil
	} else {
		return false, nil
	}
}
