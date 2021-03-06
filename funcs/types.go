//  Copyright 2013 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package funcs

type Float64 interface {
	Eval() float64
}

type Float32 interface {
	Eval() float32
}

type Int64 interface {
	Eval() int64
}

type Uint64 interface {
	Eval() uint64
}

type Int32 interface {
	Eval() int32
}

type Bool interface {
	Eval() bool
}

type String interface {
	Eval() string
}

type Bytes interface {
	Eval() []byte
}

type Uint32 interface {
	Eval() uint32
}

type Float64s interface {
	Eval() []float64
}

type Float32s interface {
	Eval() []float32
}

type Int64s interface {
	Eval() []int64
}

type Uint64s interface {
	Eval() []uint64
}

type Int32s interface {
	Eval() []int32
}

type Bools interface {
	Eval() []bool
}

type Strings interface {
	Eval() []string
}

type ListOfBytes interface {
	Eval() [][]byte
}

type Uint32s interface {
	Eval() []uint32
}
