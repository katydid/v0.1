// Code generated by protoc-gen-gogo.
// source: asm.proto
// DO NOT EDIT!

/*
	Package ast is a generated protocol buffer package.

	It is generated from these files:
		asm.proto

	It has these top-level messages:
		Rules
		Init
		Transition
		IfExpr
		StateExpr
		Expr
		List
		Function
		Terminal
		Variable
*/
package ast

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"
import types "github.com/awalterschulze/katydid/types"

import fmt "fmt"
import strings "strings"
import code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Rules struct {
	Root             *Init         `protobuf:"bytes,1,opt" json:"Root,omitempty"`
	Init             []*Init       `protobuf:"bytes,2,rep" json:"Init,omitempty"`
	Transition       []*Transition `protobuf:"bytes,3,rep" json:"Transition,omitempty"`
	IfExpr           []*IfExpr     `protobuf:"bytes,4,rep" json:"IfExpr,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Rules) Reset()      { *m = Rules{} }
func (*Rules) ProtoMessage() {}

func (m *Rules) GetRoot() *Init {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *Rules) GetInit() []*Init {
	if m != nil {
		return m.Init
	}
	return nil
}

func (m *Rules) GetTransition() []*Transition {
	if m != nil {
		return m.Transition
	}
	return nil
}

func (m *Rules) GetIfExpr() []*IfExpr {
	if m != nil {
		return m.IfExpr
	}
	return nil
}

type Init struct {
	Package          string `protobuf:"bytes,1,opt" json:"Package"`
	Message          string `protobuf:"bytes,2,opt" json:"Message"`
	State            string `protobuf:"bytes,4,opt" json:"State"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Init) Reset()      { *m = Init{} }
func (*Init) ProtoMessage() {}

func (m *Init) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *Init) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Init) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type Transition struct {
	Src              string `protobuf:"bytes,1,opt" json:"Src"`
	Input            string `protobuf:"bytes,2,opt" json:"Input"`
	Dst              string `protobuf:"bytes,3,opt" json:"Dst"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Transition) Reset()      { *m = Transition{} }
func (*Transition) ProtoMessage() {}

func (m *Transition) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *Transition) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *Transition) GetDst() string {
	if m != nil {
		return m.Dst
	}
	return ""
}

type IfExpr struct {
	Condition        *Expr      `protobuf:"bytes,1,opt" json:"Condition,omitempty"`
	Then             *StateExpr `protobuf:"bytes,2,opt" json:"Then,omitempty"`
	Else             *StateExpr `protobuf:"bytes,3,opt" json:"Else,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *IfExpr) Reset()      { *m = IfExpr{} }
func (*IfExpr) ProtoMessage() {}

func (m *IfExpr) GetCondition() *Expr {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *IfExpr) GetThen() *StateExpr {
	if m != nil {
		return m.Then
	}
	return nil
}

func (m *IfExpr) GetElse() *StateExpr {
	if m != nil {
		return m.Else
	}
	return nil
}

type StateExpr struct {
	State            *string `protobuf:"bytes,1,opt" json:"State,omitempty"`
	IfExpr           *IfExpr `protobuf:"bytes,2,opt" json:"IfExpr,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StateExpr) Reset()      { *m = StateExpr{} }
func (*StateExpr) ProtoMessage() {}

func (m *StateExpr) GetState() string {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ""
}

func (m *StateExpr) GetIfExpr() *IfExpr {
	if m != nil {
		return m.IfExpr
	}
	return nil
}

type Expr struct {
	Terminal         *Terminal `protobuf:"bytes,1,opt" json:"Terminal,omitempty"`
	List             *List     `protobuf:"bytes,2,opt" json:"List,omitempty"`
	Function         *Function `protobuf:"bytes,3,opt" json:"Function,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Expr) Reset()      { *m = Expr{} }
func (*Expr) ProtoMessage() {}

func (m *Expr) GetTerminal() *Terminal {
	if m != nil {
		return m.Terminal
	}
	return nil
}

func (m *Expr) GetList() *List {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *Expr) GetFunction() *Function {
	if m != nil {
		return m.Function
	}
	return nil
}

type List struct {
	Type             *types.Type `protobuf:"varint,1,opt,enum=types.Type" json:"Type,omitempty"`
	Elems            []*Expr     `protobuf:"bytes,2,rep" json:"Elems,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *List) Reset()      { *m = List{} }
func (*List) ProtoMessage() {}

func (m *List) GetType() types.Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return types.ENCODED_DOUBLE
}

func (m *List) GetElems() []*Expr {
	if m != nil {
		return m.Elems
	}
	return nil
}

type Function struct {
	Name             string  `protobuf:"bytes,1,opt" json:"Name"`
	Params           []*Expr `protobuf:"bytes,2,rep" json:"Params,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Function) Reset()      { *m = Function{} }
func (*Function) ProtoMessage() {}

func (m *Function) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Function) GetParams() []*Expr {
	if m != nil {
		return m.Params
	}
	return nil
}

type Terminal struct {
	DoubleValue      *float64  `protobuf:"fixed64,1,opt" json:"DoubleValue,omitempty"`
	FloatValue       *float32  `protobuf:"fixed32,2,opt" json:"FloatValue,omitempty"`
	Int64Value       *int64    `protobuf:"varint,3,opt" json:"Int64Value,omitempty"`
	Uint64Value      *uint64   `protobuf:"varint,4,opt" json:"Uint64Value,omitempty"`
	Int32Value       *int32    `protobuf:"varint,5,opt" json:"Int32Value,omitempty"`
	BoolValue        *bool     `protobuf:"varint,8,opt" json:"BoolValue,omitempty"`
	StringValue      *string   `protobuf:"bytes,9,opt" json:"StringValue,omitempty"`
	BytesValue       []byte    `protobuf:"bytes,12,opt" json:"BytesValue,omitempty"`
	Uint32Value      *uint32   `protobuf:"varint,13,opt" json:"Uint32Value,omitempty"`
	Variable         *Variable `protobuf:"bytes,50,opt" json:"Variable,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Terminal) Reset()      { *m = Terminal{} }
func (*Terminal) ProtoMessage() {}

func (m *Terminal) GetDoubleValue() float64 {
	if m != nil && m.DoubleValue != nil {
		return *m.DoubleValue
	}
	return 0
}

func (m *Terminal) GetFloatValue() float32 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Terminal) GetInt64Value() int64 {
	if m != nil && m.Int64Value != nil {
		return *m.Int64Value
	}
	return 0
}

func (m *Terminal) GetUint64Value() uint64 {
	if m != nil && m.Uint64Value != nil {
		return *m.Uint64Value
	}
	return 0
}

func (m *Terminal) GetInt32Value() int32 {
	if m != nil && m.Int32Value != nil {
		return *m.Int32Value
	}
	return 0
}

func (m *Terminal) GetBoolValue() bool {
	if m != nil && m.BoolValue != nil {
		return *m.BoolValue
	}
	return false
}

func (m *Terminal) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Terminal) GetBytesValue() []byte {
	if m != nil {
		return m.BytesValue
	}
	return nil
}

func (m *Terminal) GetUint32Value() uint32 {
	if m != nil && m.Uint32Value != nil {
		return *m.Uint32Value
	}
	return 0
}

func (m *Terminal) GetVariable() *Variable {
	if m != nil {
		return m.Variable
	}
	return nil
}

type Variable struct {
	Package          string `protobuf:"bytes,1,opt" json:"Package"`
	Message          string `protobuf:"bytes,2,opt" json:"Message"`
	Field            string `protobuf:"bytes,3,opt" json:"Field"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Variable) Reset()      { *m = Variable{} }
func (*Variable) ProtoMessage() {}

func (m *Variable) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *Variable) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Variable) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func init() {
}
func (this *StateExpr) GetValue() interface{} {
	if this.State != nil {
		return this.State
	}
	if this.IfExpr != nil {
		return this.IfExpr
	}
	return nil
}

func (this *StateExpr) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *string:
		this.State = vt
	case *IfExpr:
		this.IfExpr = vt
	default:
		return false
	}
	return true
}
func (this *Expr) GetValue() interface{} {
	if this.Terminal != nil {
		return this.Terminal
	}
	if this.List != nil {
		return this.List
	}
	if this.Function != nil {
		return this.Function
	}
	return nil
}

func (this *Expr) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *Terminal:
		this.Terminal = vt
	case *List:
		this.List = vt
	case *Function:
		this.Function = vt
	default:
		this.Terminal = new(Terminal)
		if set := this.Terminal.SetValue(value); set {
			return true
		}
		this.Terminal = nil
		return false
	}
	return true
}
func (this *Terminal) GetValue() interface{} {
	if this.DoubleValue != nil {
		return this.DoubleValue
	}
	if this.FloatValue != nil {
		return this.FloatValue
	}
	if this.Int64Value != nil {
		return this.Int64Value
	}
	if this.Uint64Value != nil {
		return this.Uint64Value
	}
	if this.Int32Value != nil {
		return this.Int32Value
	}
	if this.BoolValue != nil {
		return this.BoolValue
	}
	if this.StringValue != nil {
		return this.StringValue
	}
	if this.BytesValue != nil {
		return this.BytesValue
	}
	if this.Uint32Value != nil {
		return this.Uint32Value
	}
	if this.Variable != nil {
		return this.Variable
	}
	return nil
}

func (this *Terminal) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *float64:
		this.DoubleValue = vt
	case *float32:
		this.FloatValue = vt
	case *int64:
		this.Int64Value = vt
	case *uint64:
		this.Uint64Value = vt
	case *int32:
		this.Int32Value = vt
	case *bool:
		this.BoolValue = vt
	case *string:
		this.StringValue = vt
	case []byte:
		this.BytesValue = vt
	case *uint32:
		this.Uint32Value = vt
	case *Variable:
		this.Variable = vt
	default:
		return false
	}
	return true
}
func (this *Rules) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Rules{` + `Root:` + fmt.Sprintf("%#v", this.Root), `Init:` + fmt.Sprintf("%#v", this.Init), `Transition:` + fmt.Sprintf("%#v", this.Transition), `IfExpr:` + fmt.Sprintf("%#v", this.IfExpr), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Init) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Init{` + `Package:` + fmt.Sprintf("%#v", this.Package), `Message:` + fmt.Sprintf("%#v", this.Message), `State:` + fmt.Sprintf("%#v", this.State), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Transition) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Transition{` + `Src:` + fmt.Sprintf("%#v", this.Src), `Input:` + fmt.Sprintf("%#v", this.Input), `Dst:` + fmt.Sprintf("%#v", this.Dst), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *IfExpr) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.IfExpr{` + `Condition:` + fmt.Sprintf("%#v", this.Condition), `Then:` + fmt.Sprintf("%#v", this.Then), `Else:` + fmt.Sprintf("%#v", this.Else), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *StateExpr) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.StateExpr{` + `State:` + valueToGoStringAsm(this.State, "string"), `IfExpr:` + fmt.Sprintf("%#v", this.IfExpr), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Expr) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Expr{` + `Terminal:` + fmt.Sprintf("%#v", this.Terminal), `List:` + fmt.Sprintf("%#v", this.List), `Function:` + fmt.Sprintf("%#v", this.Function), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *List) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.List{` + `Type:` + valueToGoStringAsm(this.Type, "ast.types.Type"), `Elems:` + fmt.Sprintf("%#v", this.Elems), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Function) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Function{` + `Name:` + fmt.Sprintf("%#v", this.Name), `Params:` + fmt.Sprintf("%#v", this.Params), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Terminal) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Terminal{` + `DoubleValue:` + valueToGoStringAsm(this.DoubleValue, "float64"), `FloatValue:` + valueToGoStringAsm(this.FloatValue, "float32"), `Int64Value:` + valueToGoStringAsm(this.Int64Value, "int64"), `Uint64Value:` + valueToGoStringAsm(this.Uint64Value, "uint64"), `Int32Value:` + valueToGoStringAsm(this.Int32Value, "int32"), `BoolValue:` + valueToGoStringAsm(this.BoolValue, "bool"), `StringValue:` + valueToGoStringAsm(this.StringValue, "string"), `BytesValue:` + valueToGoStringAsm(this.BytesValue, "byte"), `Uint32Value:` + valueToGoStringAsm(this.Uint32Value, "uint32"), `Variable:` + fmt.Sprintf("%#v", this.Variable), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Variable) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Variable{` + `Package:` + fmt.Sprintf("%#v", this.Package), `Message:` + fmt.Sprintf("%#v", this.Message), `Field:` + fmt.Sprintf("%#v", this.Field), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func valueToGoStringAsm(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionMapToGoStringAsm(e map[int32]code_google_com_p_gogoprotobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func extensionSliceToGoStringAsm(e []code_google_com_p_gogoprotobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "[]proto.Extension{"
	ss := make([]string, len(e))
	for i, ext := range e {
		ss[i] = ext.GoString()
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
