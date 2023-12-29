package lib

import (
	"fmt"
	"os"

	. "github.com/dave/jennifer/jen"
)

func GenContract(cont *Contract, out *os.File) error {
	file := NewFile("main")
	cases := make([]Code, len(cont.Functions))
	for _, f := range cont.Functions {
		sel := f.FuncMetadata.Selector.UInt32()
		params := make([]Code, len(f.Params))
		for j, pin := range f.Params {
			pout := Qual("asgo", fmt.Sprintf("Decode%s", pin.Type)).Call(Id("cd").Index(Lit(4+j*32), Lit(4+(j+1)*32)))
			params = append(params, pout)
		}

		var callFuncCode Code
		if len(f.Returns) > 1 {
			callFuncCode = List(Id("ret"), Err()).Op(":=").Qual("cont", f.Name).Call(params...)
		} else {
			callFuncCode = Err().Op(":=").Qual("cont", f.Name).Call(params...)
		}

		c := Case(Lit(sel)).Block(
			callFuncCode,
			If(Err().Op("!=").Nil()).Block(Return(Err())),
		)
		cases = append(cases, c)
	}
	file.Func().Id("user_entrypoint").Params().Block(
		Id("cont").Op(":=").Id(fmt.Sprintf("New%s", cont.Name)).Call(),
		Id("cd").Op(":=").Qual("asgo", "GetCalldata").Call(),
		If(Len(Id("cd")).Op("<").Lit(4)).Block(
			Return(Err()),
		),
		Id("sel").Op(":=").Qual("asgo", "ToSelector").Call(Id("cd").Index(Empty(), Lit(4))),

		Switch(Id("sel")).Block(
			cases...,
		),
		Qual("fmt", "Println").Call(Lit("Hello, world")),
	)

	fmt.Fprintf(out, "%#v", file)

	return nil

}
