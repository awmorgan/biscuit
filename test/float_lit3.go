// errorcheck

// Check flagging of invalid conversion of constant to float32/float64 near min/max boundaries.

// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

var x = []interface{}{
	float32(-340282356779733661637539395458142568448), // ERROR "constant -3\.40282e\+38 overflows float32"
	float32(-340282356779733661637539395458142568447),
	float32(-340282326356119256160033759537265639424),
	float32(340282326356119256160033759537265639424),
	float32(340282356779733661637539395458142568447),
	float32(340282356779733661637539395458142568448), // ERROR "constant 3\.40282e\+38 overflows float32"
	-1e1000, // ERROR "constant -1\.00000e\+1000 overflows float64"
	float64(-1.797693134862315907937289714053e+308), // ERROR "constant -1\.79769e\+308 overflows float64"
	float64(-1.797693134862315807937289714053e+308),
	float64(-1.797693134862315708145274237317e+308),
	float64(-1.797693134862315608353258760581e+308),
	float64(1.797693134862315608353258760581e+308),
	float64(1.797693134862315708145274237317e+308),
	float64(1.797693134862315807937289714053e+308),
	float64(1.797693134862315907937289714053e+308), // ERROR "constant 1\.79769e\+308 overflows float64"
	1e1000, // ERROR "constant 1\.00000e\+1000 overflows float64"
}