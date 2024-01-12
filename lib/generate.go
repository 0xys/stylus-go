package lib

import (
	"fmt"
	"os"

	. "github.com/dave/jennifer/jen"
)

const sdkIdentifier = "asgo"

func GenContract(cont *Contract, out *os.File) error {
	file := NewFile("main")
	file.HeaderComment(fmt.Sprintf("Code generated by %s. DO NOT EDIT.", sdkIdentifier))
	cases := make([]Code, len(cont.Functions))
	for _, f := range cont.Functions {
		sel := f.FuncMetadata.Selector.UInt32()
		params := make([]Code, len(f.Params))
		for j, pin := range f.Params {
			pout := Qual(sdkIdentifier, fmt.Sprintf("Decode%s", pin.Type)).Call(Id("cd").Index(Lit(4+j*32), Lit(4+(j+1)*32)))
			params = append(params, pout)
		}

		var callFuncCodes []Code
		if !f.IsPayable() {
			callFuncCodes = append(callFuncCodes, If(Op("!").Qual(sdkIdentifier, "GetMsgValue()").Dot("IsZero").Call()).Block(Return()))
		}
		if f.IsPure() {
			callFuncCodes = append(callFuncCodes, Qual(sdkIdentifier, "SetPure").Call())
		}
		if f.IsView() {
			callFuncCodes = append(callFuncCodes, Qual(sdkIdentifier, "SetView").Call())
		}

		if len(f.Returns) > 1 {
			callFuncCodes = append(callFuncCodes, List(Id("ret"), Err()).Op(":=").Id("cont").Dot(f.Name).Call(params...))
			callFuncCodes = append(callFuncCodes, If(Err().Op("!=").Nil()).Block(Qual(sdkIdentifier, "SetReturnString").Call(Id("err").Dot("Error").Call()), Return()))
			callFuncCodes = append(callFuncCodes, Qual(sdkIdentifier, "SetReturnBytes").Call(Id("ret").Dot("EncodeToBytes").Call()))
		} else {
			callFuncCodes = append(callFuncCodes, Err().Op(":=").Id("cont").Dot(f.Name).Call(params...))
			callFuncCodes = append(callFuncCodes, If(Err().Op("!=").Nil()).Block(Qual(sdkIdentifier, "SetReturnString").Call(Id("err").Dot("Error").Call())))
		}

		c := Case(Lit(sel)).Block(
			callFuncCodes...,
		)
		cases = append(cases, c)
	}
	file.Comment("export user_entrypoint")
	file.Func().Id("user_entrypoint").Params().Block(
		Qual(sdkIdentifier, "Init").Call(),
		Defer().Qual(sdkIdentifier, "Flush").Call(),
		Line(),
		Id("cont").Op(":=").Op("&").Id(cont.Name).Block(),
		Id("cd").Op(":=").Qual(sdkIdentifier, "GetCalldata").Call(),
		If(Len(Id("cd")).Op("<").Lit(4)).Block(
			Return(),
		),
		Id("sel").Op(":=").Qual(sdkIdentifier, "ToSelector").Call(Id("cd").Index(Empty(), Lit(4))),

		Switch(Id("sel")).Block(
			cases...,
		),
	)

	file.Comment("dummy main is needed")
	file.Func().Id("main").Params().Block(
		Id("user_entrypoint").Call(),
	)

	fmt.Fprintf(out, "%#v", file)

	return nil

}
