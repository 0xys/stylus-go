package main

import (
	asgo "asgo"
	"fmt"
)

func user_entrypoint() {
	cont := NewFooContract()
	cd := asgo.GetCalldata()
	if len(cd) < 4 {
		return err
	}
	sel := asgo.ToSelector(cd[:4])
	switch sel {
	case uint32(0xc49c36c):
		err := cont.SayHi()
		if err != nil {
			asgo.SetReturnString(err.Error())
		}
	case uint32(0xa9059cbb):
		err := cont.Transfer(asgo.DecodeAddress(cd[4:36]), asgo.DecodeU256(cd[36:68]))
		if err != nil {
			asgo.SetReturnString(err.Error())
		}
	case uint32(0x5cdf6ad2):
		err := cont.Receive(asgo.Decodeuint64(cd[4:36]))
		if err != nil {
			asgo.SetReturnString(err.Error())
		}
	case uint32(0x23b872dd):
		ret, err := cont.TransferFrom(asgo.DecodeAddress(cd[4:36]), asgo.DecodeU256(cd[36:68]))
		if err != nil {
			asgo.SetReturnString(err.Error())
			return
		}
		asgo.SetReturnBytes(ret.EncodeToBytes())
	}
	fmt.Println("Hello, world")
}
