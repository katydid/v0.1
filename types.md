---
layout: page
title: Types
---

Katydid's functions are type safe.
Many functions are allowed to have the same name if their input parameter type signatures are different.
For example the `eq` function exists for multiple types like `eq(string, string) bool` and `eq(int64, int64) bool`.

Katydid does not allow the creation of your own custom types and is limited to a few native types.
These types match the decoded types of the native types existing in protocol buffers:

* double
* float
* int64
* uint64
* int32
* bool
* string
* bytes
* uint32

Katydid also includes list types for each of these the native types:

* []double
* []float
* []int64
* []uint64
* []int32
* []bool
* []string
* []bytes
* []uint32