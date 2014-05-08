// Code generated by funcs-gen.
// DO NOT EDIT!

package funcs

type listOfFloat64 struct {
	List []Float64
}

func NewListOfFloat64(v []Float64) Float64s {
	return &listOfFloat64{v}
}

func (this *listOfFloat64) Eval(buf []byte) []float64 {
	res := make([]float64, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval(buf)
	}
	return res
}

type listOfFloat32 struct {
	List []Float32
}

func NewListOfFloat32(v []Float32) Float32s {
	return &listOfFloat32{v}
}

func (this *listOfFloat32) Eval(buf []byte) []float32 {
	res := make([]float32, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval(buf)
	}
	return res
}

type listOfInt64 struct {
	List []Int64
}

func NewListOfInt64(v []Int64) Int64s {
	return &listOfInt64{v}
}

func (this *listOfInt64) Eval(buf []byte) []int64 {
	res := make([]int64, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval(buf)
	}
	return res
}

type listOfUint64 struct {
	List []Uint64
}

func NewListOfUint64(v []Uint64) Uint64s {
	return &listOfUint64{v}
}

func (this *listOfUint64) Eval(buf []byte) []uint64 {
	res := make([]uint64, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval(buf)
	}
	return res
}

type listOfInt32 struct {
	List []Int32
}

func NewListOfInt32(v []Int32) Int32s {
	return &listOfInt32{v}
}

func (this *listOfInt32) Eval(buf []byte) []int32 {
	res := make([]int32, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval(buf)
	}
	return res
}

type listOfBool struct {
	List []Bool
}

func NewListOfBool(v []Bool) Bools {
	return &listOfBool{v}
}

func (this *listOfBool) Eval(buf []byte) []bool {
	res := make([]bool, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval(buf)
	}
	return res
}

type listOfString struct {
	List []String
}

func NewListOfString(v []String) Strings {
	return &listOfString{v}
}

func (this *listOfString) Eval(buf []byte) []string {
	res := make([]string, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval(buf)
	}
	return res
}

type listOfBytes struct {
	List []Bytes
}

func NewListOfBytes(v []Bytes) ListOfBytes {
	return &listOfBytes{v}
}

func (this *listOfBytes) Eval(buf []byte) [][]byte {
	res := make([][]byte, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval(buf)
	}
	return res
}

type listOfUint32 struct {
	List []Uint32
}

func NewListOfUint32(v []Uint32) Uint32s {
	return &listOfUint32{v}
}

func (this *listOfUint32) Eval(buf []byte) []uint32 {
	res := make([]uint32, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval(buf)
	}
	return res
}
