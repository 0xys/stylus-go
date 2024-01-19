package lib

import (
	"encoding/hex"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"golang.org/x/crypto/sha3"
)

const ABIDirective string = "abi: "

type FuncSelector []byte

type FuncMetadata struct {
	Signature string
	Selector  FuncSelector
	Modifiers []Modifier
}

func (f *FuncMetadata) IsPayable() bool {
	for _, m := range f.Modifiers {
		if m == Payable {
			return true
		}
	}
	return false
}

func (f *FuncMetadata) IsPure() bool {
	for _, m := range f.Modifiers {
		if m == Pure {
			return true
		}
	}
	return false
}

func (f *FuncMetadata) IsView() bool {
	for _, m := range f.Modifiers {
		if m == View {
			return true
		}
	}
	return false
}

func NewFuncMetadata(comment string) (*FuncMetadata, error) {
	ss := strings.Split(comment, ABIDirective)
	if len(ss) < 2 {
		return nil, fmt.Errorf("not ABI %s", comment)
	}
	noNewLine := strings.Trim(ss[1], "\n")
	noPrefixSpace := strings.TrimPrefix(noNewLine, " ")
	words := strings.Split(noPrefixSpace, " ")
	sig := words[0]

	ret := &FuncMetadata{
		Signature: sig,
		Selector:  toSelectorBytes(sig),
		Modifiers: []Modifier{},
	}

	for _, w := range words[1:] {
		switch w {
		case "payable":
			ret.Modifiers = append(ret.Modifiers, Payable)
		case "view":
			ret.Modifiers = append(ret.Modifiers, View)
		case "pure":
			ret.Modifiers = append(ret.Modifiers, Pure)
		}
	}
	return ret, nil
}

func (s FuncSelector) Hex() string {
	return hex.EncodeToString(s)
}
func (s FuncSelector) Hex0x() string {
	return "0x" + hex.EncodeToString(s)
}
func (s FuncSelector) UInt32() uint32 {
	return ToSelector(s)
}
func ToSelector(in []byte) uint32 {
	ret := uint32(0)
	ret += uint32(in[3])
	ret += uint32(in[2]) << 8
	ret += uint32(in[1]) << 16
	ret += uint32(in[0]) << 24
	return ret
}

func toSelectorBytes(funcSignature string) FuncSelector {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(funcSignature))
	hash := hasher.Sum(nil)
	return hash[:4]
}

type Variable struct {
	Name string
	Type string
}

type Modifier int

const (
	Payable Modifier = 1
	Pure    Modifier = 2
	View    Modifier = 3
)

type Contract struct {
	Name      string
	Functions []*Function
	Fields    []*Variable
}

type Function struct {
	Name    string
	Params  []*Variable
	Returns []string
	*FuncMetadata
}

func (f *Function) Print() {
	fmt.Printf("%s %s %s %+v\n", f.Name, f.Signature, f.Selector.Hex0x(), f.Modifiers)
	if len(f.Params) == 0 {
		return
	}
	fmt.Print(" - ")
	for _, p := range f.Params {
		fmt.Printf("%s %s, ", p.Name, p.Type)
	}
	fmt.Println("")
}

func main() {
	if err := realMain(); err != nil {
		panic(err)
	}
}

func parseStructName(f *ast.File) (*Contract, error) {
	ret := &Contract{}
	found := false
	for _, obj := range f.Scope.Objects {
		if t, ok := obj.Decl.(*ast.TypeSpec); ok {
			if tt, ok := t.Type.(*ast.StructType); ok {
				if found {
					return nil, fmt.Errorf("more than one struct is declared")
				}
				found = true
				ret.Name = t.Name.String()
				for _, f2 := range tt.Fields.List {
					ret.Fields = append(ret.Fields, &Variable{
						Name: f2.Names[0].Name,
						Type: reflect.TypeOf(f2.Type).String(),
					})
				}
			}
		}
	}
	if !found {
		return nil, fmt.Errorf("no struct is found")
	}
	return ret, nil
}

func parseFunctions(cont *Contract, f *ast.File) error {
	for _, node := range f.Decls {
		switch d := node.(type) {
		case *ast.FuncDecl:
			if d.Recv == nil {
				continue
			}

			if len(d.Recv.List) != 1 {
				continue
			}
			rname := getReceiverName(d.Recv.List[0])
			if rname != cont.Name {
				continue
			}

			md, err := getFuncMetadata(d)
			if err != nil {
				return err
			}
			f := &Function{
				Name:         d.Name.String(),
				FuncMetadata: md,
			}

			for _, param := range d.Type.Params.List {
				types := fmt.Sprintf("%s", param.Type)
				if tt, ok := param.Type.(*ast.SelectorExpr); ok {
					types = cases.Title(language.Und).String(tt.Sel.String())
					types = tt.Sel.String()
				} else {
					types = cases.Title(language.Und).String(types)
				}

				f.Params = append(f.Params, &Variable{
					Name: param.Names[0].String(),
					Type: types,
				})
			}
			if d.Type.Results != nil {
				for _, r := range d.Type.Results.List {
					f.Returns = append(f.Returns, reflect.TypeOf(r.Type).String())
				}
			}

			cont.Functions = append(cont.Functions, f)
		}
	}
	return nil
}

func getReceiverName(r *ast.Field) string {
	switch t := r.Type.(type) {
	case *ast.StarExpr:
		if v, ok := t.X.(*ast.Ident); ok {
			return v.Name
		}
	case *ast.Ident:
		return t.Name
	}
	return ""
}

func getFuncMetadata(d *ast.FuncDecl) (*FuncMetadata, error) {
	cmmt := d.Doc.Text()
	return NewFuncMetadata(cmmt)
}

func ProcessFile(filepath string) (*Contract, error) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, filepath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	// spew.Fdump(out, f)

	cont, err := parseStructName(f)
	if err != nil {
		return nil, err
	}

	if err = parseFunctions(cont, f); err != nil {
		return nil, err
	}

	/*fmt.Println(cont.Name)
	for _, m := range cont.Functions {
		m.Print()
	}*/
	return cont, nil
}

func realMain() error {
	_, err := ProcessFile("templates/main.go")
	return err
}
