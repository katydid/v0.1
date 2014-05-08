// Code generated by funcs-gen.
// DO NOT EDIT!

package funcs

type rangeFloat64s struct {
	List  Float64s
	First Int64
	Last  Int64
}

func (this *rangeFloat64s) Eval(buf []byte) []float64 {
	list := this.List.Eval(buf)
	first := int(this.First.Eval(buf))
	last := int(this.Last.Eval(buf))
	first = first % len(list)
	if last > len(list) {
		last = last % len(list)
	}
	if first > last {
		first = last
	}
	return list[first:last]
}

func init() {
	Register("range", new(rangeFloat64s))
}

type rangeFloat32s struct {
	List  Float32s
	First Int64
	Last  Int64
}

func (this *rangeFloat32s) Eval(buf []byte) []float32 {
	list := this.List.Eval(buf)
	first := int(this.First.Eval(buf))
	last := int(this.Last.Eval(buf))
	first = first % len(list)
	if last > len(list) {
		last = last % len(list)
	}
	if first > last {
		first = last
	}
	return list[first:last]
}

func init() {
	Register("range", new(rangeFloat32s))
}

type rangeInt64s struct {
	List  Int64s
	First Int64
	Last  Int64
}

func (this *rangeInt64s) Eval(buf []byte) []int64 {
	list := this.List.Eval(buf)
	first := int(this.First.Eval(buf))
	last := int(this.Last.Eval(buf))
	first = first % len(list)
	if last > len(list) {
		last = last % len(list)
	}
	if first > last {
		first = last
	}
	return list[first:last]
}

func init() {
	Register("range", new(rangeInt64s))
}

type rangeUint64s struct {
	List  Uint64s
	First Int64
	Last  Int64
}

func (this *rangeUint64s) Eval(buf []byte) []uint64 {
	list := this.List.Eval(buf)
	first := int(this.First.Eval(buf))
	last := int(this.Last.Eval(buf))
	first = first % len(list)
	if last > len(list) {
		last = last % len(list)
	}
	if first > last {
		first = last
	}
	return list[first:last]
}

func init() {
	Register("range", new(rangeUint64s))
}

type rangeInt32s struct {
	List  Int32s
	First Int64
	Last  Int64
}

func (this *rangeInt32s) Eval(buf []byte) []int32 {
	list := this.List.Eval(buf)
	first := int(this.First.Eval(buf))
	last := int(this.Last.Eval(buf))
	first = first % len(list)
	if last > len(list) {
		last = last % len(list)
	}
	if first > last {
		first = last
	}
	return list[first:last]
}

func init() {
	Register("range", new(rangeInt32s))
}

type rangeUint32s struct {
	List  Uint32s
	First Int64
	Last  Int64
}

func (this *rangeUint32s) Eval(buf []byte) []uint32 {
	list := this.List.Eval(buf)
	first := int(this.First.Eval(buf))
	last := int(this.Last.Eval(buf))
	first = first % len(list)
	if last > len(list) {
		last = last % len(list)
	}
	if first > last {
		first = last
	}
	return list[first:last]
}

func init() {
	Register("range", new(rangeUint32s))
}

type rangeBools struct {
	List  Bools
	First Int64
	Last  Int64
}

func (this *rangeBools) Eval(buf []byte) []bool {
	list := this.List.Eval(buf)
	first := int(this.First.Eval(buf))
	last := int(this.Last.Eval(buf))
	first = first % len(list)
	if last > len(list) {
		last = last % len(list)
	}
	if first > last {
		first = last
	}
	return list[first:last]
}

func init() {
	Register("range", new(rangeBools))
}

type rangeStrings struct {
	List  Strings
	First Int64
	Last  Int64
}

func (this *rangeStrings) Eval(buf []byte) []string {
	list := this.List.Eval(buf)
	first := int(this.First.Eval(buf))
	last := int(this.Last.Eval(buf))
	first = first % len(list)
	if last > len(list) {
		last = last % len(list)
	}
	if first > last {
		first = last
	}
	return list[first:last]
}

func init() {
	Register("range", new(rangeStrings))
}

type rangeListOfBytes struct {
	List  ListOfBytes
	First Int64
	Last  Int64
}

func (this *rangeListOfBytes) Eval(buf []byte) [][]byte {
	list := this.List.Eval(buf)
	first := int(this.First.Eval(buf))
	last := int(this.Last.Eval(buf))
	first = first % len(list)
	if last > len(list) {
		last = last % len(list)
	}
	if first > last {
		first = last
	}
	return list[first:last]
}

func init() {
	Register("range", new(rangeListOfBytes))
}
