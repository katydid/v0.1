// Code generated by funcs-gen.
// DO NOT EDIT!

package funcs

type elemFloat64s struct {
	List  Float64s
	Index Int64
	Thrower
}

func (this *elemFloat64s) Eval() float64 {
	list := this.List.Eval()
	index := int(this.Index.Eval())
	if len(list) == 0 {
		return this.ThrowFloat64(NewRangeCheckErr(index, len(list)))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return this.ThrowFloat64(NewRangeCheckErr(index, len(list)))
	}
	return list[index]
}

func init() {
	Register("elem", new(elemFloat64s))
}

type elemFloat32s struct {
	List  Float32s
	Index Int64
	Thrower
}

func (this *elemFloat32s) Eval() float32 {
	list := this.List.Eval()
	index := int(this.Index.Eval())
	if len(list) == 0 {
		return this.ThrowFloat32(NewRangeCheckErr(index, len(list)))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return this.ThrowFloat32(NewRangeCheckErr(index, len(list)))
	}
	return list[index]
}

func init() {
	Register("elem", new(elemFloat32s))
}

type elemInt64s struct {
	List  Int64s
	Index Int64
	Thrower
}

func (this *elemInt64s) Eval() int64 {
	list := this.List.Eval()
	index := int(this.Index.Eval())
	if len(list) == 0 {
		return this.ThrowInt64(NewRangeCheckErr(index, len(list)))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return this.ThrowInt64(NewRangeCheckErr(index, len(list)))
	}
	return list[index]
}

func init() {
	Register("elem", new(elemInt64s))
}

type elemUint64s struct {
	List  Uint64s
	Index Int64
	Thrower
}

func (this *elemUint64s) Eval() uint64 {
	list := this.List.Eval()
	index := int(this.Index.Eval())
	if len(list) == 0 {
		return this.ThrowUint64(NewRangeCheckErr(index, len(list)))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return this.ThrowUint64(NewRangeCheckErr(index, len(list)))
	}
	return list[index]
}

func init() {
	Register("elem", new(elemUint64s))
}

type elemInt32s struct {
	List  Int32s
	Index Int64
	Thrower
}

func (this *elemInt32s) Eval() int32 {
	list := this.List.Eval()
	index := int(this.Index.Eval())
	if len(list) == 0 {
		return this.ThrowInt32(NewRangeCheckErr(index, len(list)))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return this.ThrowInt32(NewRangeCheckErr(index, len(list)))
	}
	return list[index]
}

func init() {
	Register("elem", new(elemInt32s))
}

type elemUint32s struct {
	List  Uint32s
	Index Int64
	Thrower
}

func (this *elemUint32s) Eval() uint32 {
	list := this.List.Eval()
	index := int(this.Index.Eval())
	if len(list) == 0 {
		return this.ThrowUint32(NewRangeCheckErr(index, len(list)))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return this.ThrowUint32(NewRangeCheckErr(index, len(list)))
	}
	return list[index]
}

func init() {
	Register("elem", new(elemUint32s))
}

type elemBools struct {
	List  Bools
	Index Int64
	Thrower
}

func (this *elemBools) Eval() bool {
	list := this.List.Eval()
	index := int(this.Index.Eval())
	if len(list) == 0 {
		return this.ThrowBool(NewRangeCheckErr(index, len(list)))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return this.ThrowBool(NewRangeCheckErr(index, len(list)))
	}
	return list[index]
}

func init() {
	Register("elem", new(elemBools))
}

type elemStrings struct {
	List  Strings
	Index Int64
	Thrower
}

func (this *elemStrings) Eval() string {
	list := this.List.Eval()
	index := int(this.Index.Eval())
	if len(list) == 0 {
		return this.ThrowString(NewRangeCheckErr(index, len(list)))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return this.ThrowString(NewRangeCheckErr(index, len(list)))
	}
	return list[index]
}

func init() {
	Register("elem", new(elemStrings))
}

type elemListOfBytes struct {
	List  ListOfBytes
	Index Int64
	Thrower
}

func (this *elemListOfBytes) Eval() []byte {
	list := this.List.Eval()
	index := int(this.Index.Eval())
	if len(list) == 0 {
		return this.ThrowBytes(NewRangeCheckErr(index, len(list)))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return this.ThrowBytes(NewRangeCheckErr(index, len(list)))
	}
	return list[index]
}

func init() {
	Register("elem", new(elemListOfBytes))
}
