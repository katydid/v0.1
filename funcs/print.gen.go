// Code generated by funcs-gen.
// DO NOT EDIT!

package funcs

import (
	"fmt"
)

type printFloat64 struct {
	E Float64
}

func (this *printFloat64) Eval() float64 {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printFloat64) IsVariable() {}

func init() {
	Register("print", new(printFloat64))
}

type printFloat32 struct {
	E Float32
}

func (this *printFloat32) Eval() float32 {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printFloat32) IsVariable() {}

func init() {
	Register("print", new(printFloat32))
}

type printInt64 struct {
	E Int64
}

func (this *printInt64) Eval() int64 {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printInt64) IsVariable() {}

func init() {
	Register("print", new(printInt64))
}

type printUint64 struct {
	E Uint64
}

func (this *printUint64) Eval() uint64 {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printUint64) IsVariable() {}

func init() {
	Register("print", new(printUint64))
}

type printInt32 struct {
	E Int32
}

func (this *printInt32) Eval() int32 {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printInt32) IsVariable() {}

func init() {
	Register("print", new(printInt32))
}

type printUint32 struct {
	E Uint32
}

func (this *printUint32) Eval() uint32 {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printUint32) IsVariable() {}

func init() {
	Register("print", new(printUint32))
}

type printBool struct {
	E Bool
}

func (this *printBool) Eval() bool {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printBool) IsVariable() {}

func init() {
	Register("print", new(printBool))
}

type printString struct {
	E String
}

func (this *printString) Eval() string {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printString) IsVariable() {}

func init() {
	Register("print", new(printString))
}

type printBytes struct {
	E Bytes
}

func (this *printBytes) Eval() []byte {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printBytes) IsVariable() {}

func init() {
	Register("print", new(printBytes))
}

type printFloat64s struct {
	E Float64s
}

func (this *printFloat64s) Eval() []float64 {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printFloat64s) IsVariable() {}

func init() {
	Register("print", new(printFloat64s))
}

type printFloat32s struct {
	E Float32s
}

func (this *printFloat32s) Eval() []float32 {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printFloat32s) IsVariable() {}

func init() {
	Register("print", new(printFloat32s))
}

type printInt64s struct {
	E Int64s
}

func (this *printInt64s) Eval() []int64 {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printInt64s) IsVariable() {}

func init() {
	Register("print", new(printInt64s))
}

type printUint64s struct {
	E Uint64s
}

func (this *printUint64s) Eval() []uint64 {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printUint64s) IsVariable() {}

func init() {
	Register("print", new(printUint64s))
}

type printInt32s struct {
	E Int32s
}

func (this *printInt32s) Eval() []int32 {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printInt32s) IsVariable() {}

func init() {
	Register("print", new(printInt32s))
}

type printUint32s struct {
	E Uint32s
}

func (this *printUint32s) Eval() []uint32 {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printUint32s) IsVariable() {}

func init() {
	Register("print", new(printUint32s))
}

type printBools struct {
	E Bools
}

func (this *printBools) Eval() []bool {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printBools) IsVariable() {}

func init() {
	Register("print", new(printBools))
}

type printStrings struct {
	E Strings
}

func (this *printStrings) Eval() []string {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printStrings) IsVariable() {}

func init() {
	Register("print", new(printStrings))
}

type printListOfBytes struct {
	E ListOfBytes
}

func (this *printListOfBytes) Eval() [][]byte {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *printListOfBytes) IsVariable() {}

func init() {
	Register("print", new(printListOfBytes))
}
