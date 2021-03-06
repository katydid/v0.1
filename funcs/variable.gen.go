// Code generated by funcs-gen.
// DO NOT EDIT!

package funcs

import (
	"github.com/katydid/katydid/serialize"
)

type varFloat64 struct {
	Dec serialize.Decoder
	Thrower
}

var _ Decoder = &varFloat64{}
var _ Variable = &varFloat64{}

func (this *varFloat64) Eval() float64 {
	v, err := this.Dec.Float64()
	if err != nil {
		return this.ThrowFloat64(err)
	}
	return v
}

func (this *varFloat64) IsVariable() {}

func (this *varFloat64) SetDecoder(dec serialize.Decoder) {
	this.Dec = dec
}

func (this *varFloat64) String() string {
	return "varFloat64"
}

func NewFloat64Variable() *varFloat64 {
	return &varFloat64{}
}

type varFloat32 struct {
	Dec serialize.Decoder
	Thrower
}

var _ Decoder = &varFloat32{}
var _ Variable = &varFloat32{}

func (this *varFloat32) Eval() float32 {
	v, err := this.Dec.Float32()
	if err != nil {
		return this.ThrowFloat32(err)
	}
	return v
}

func (this *varFloat32) IsVariable() {}

func (this *varFloat32) SetDecoder(dec serialize.Decoder) {
	this.Dec = dec
}

func (this *varFloat32) String() string {
	return "varFloat32"
}

func NewFloat32Variable() *varFloat32 {
	return &varFloat32{}
}

type varInt64 struct {
	Dec serialize.Decoder
	Thrower
}

var _ Decoder = &varInt64{}
var _ Variable = &varInt64{}

func (this *varInt64) Eval() int64 {
	v, err := this.Dec.Int64()
	if err != nil {
		return this.ThrowInt64(err)
	}
	return v
}

func (this *varInt64) IsVariable() {}

func (this *varInt64) SetDecoder(dec serialize.Decoder) {
	this.Dec = dec
}

func (this *varInt64) String() string {
	return "varInt64"
}

func NewInt64Variable() *varInt64 {
	return &varInt64{}
}

type varUint64 struct {
	Dec serialize.Decoder
	Thrower
}

var _ Decoder = &varUint64{}
var _ Variable = &varUint64{}

func (this *varUint64) Eval() uint64 {
	v, err := this.Dec.Uint64()
	if err != nil {
		return this.ThrowUint64(err)
	}
	return v
}

func (this *varUint64) IsVariable() {}

func (this *varUint64) SetDecoder(dec serialize.Decoder) {
	this.Dec = dec
}

func (this *varUint64) String() string {
	return "varUint64"
}

func NewUint64Variable() *varUint64 {
	return &varUint64{}
}

type varInt32 struct {
	Dec serialize.Decoder
	Thrower
}

var _ Decoder = &varInt32{}
var _ Variable = &varInt32{}

func (this *varInt32) Eval() int32 {
	v, err := this.Dec.Int32()
	if err != nil {
		return this.ThrowInt32(err)
	}
	return v
}

func (this *varInt32) IsVariable() {}

func (this *varInt32) SetDecoder(dec serialize.Decoder) {
	this.Dec = dec
}

func (this *varInt32) String() string {
	return "varInt32"
}

func NewInt32Variable() *varInt32 {
	return &varInt32{}
}

type varUint32 struct {
	Dec serialize.Decoder
	Thrower
}

var _ Decoder = &varUint32{}
var _ Variable = &varUint32{}

func (this *varUint32) Eval() uint32 {
	v, err := this.Dec.Uint32()
	if err != nil {
		return this.ThrowUint32(err)
	}
	return v
}

func (this *varUint32) IsVariable() {}

func (this *varUint32) SetDecoder(dec serialize.Decoder) {
	this.Dec = dec
}

func (this *varUint32) String() string {
	return "varUint32"
}

func NewUint32Variable() *varUint32 {
	return &varUint32{}
}

type varBool struct {
	Dec serialize.Decoder
	Thrower
}

var _ Decoder = &varBool{}
var _ Variable = &varBool{}

func (this *varBool) Eval() bool {
	v, err := this.Dec.Bool()
	if err != nil {
		return this.ThrowBool(err)
	}
	return v
}

func (this *varBool) IsVariable() {}

func (this *varBool) SetDecoder(dec serialize.Decoder) {
	this.Dec = dec
}

func (this *varBool) String() string {
	return "varBool"
}

func NewBoolVariable() *varBool {
	return &varBool{}
}

type varString struct {
	Dec serialize.Decoder
	Thrower
}

var _ Decoder = &varString{}
var _ Variable = &varString{}

func (this *varString) Eval() string {
	v, err := this.Dec.String()
	if err != nil {
		return this.ThrowString(err)
	}
	return v
}

func (this *varString) IsVariable() {}

func (this *varString) SetDecoder(dec serialize.Decoder) {
	this.Dec = dec
}

func (this *varString) String() string {
	return "varString"
}

func NewStringVariable() *varString {
	return &varString{}
}

type varBytes struct {
	Dec serialize.Decoder
	Thrower
}

var _ Decoder = &varBytes{}
var _ Variable = &varBytes{}

func (this *varBytes) Eval() []byte {
	v, err := this.Dec.Bytes()
	if err != nil {
		return this.ThrowBytes(err)
	}
	return v
}

func (this *varBytes) IsVariable() {}

func (this *varBytes) SetDecoder(dec serialize.Decoder) {
	this.Dec = dec
}

func (this *varBytes) String() string {
	return "varBytes"
}

func NewBytesVariable() *varBytes {
	return &varBytes{}
}
