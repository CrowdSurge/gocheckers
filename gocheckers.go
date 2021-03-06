// Copyright (c) 2012, CrowdSurge Ltd.
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:

// Redistributions of source code must retain the above copyright notice, this
// list of conditions and the following disclaimer.

// Redistributions in binary form must reproduce the above copyright notice, this
// list of conditions and the following disclaimer in the documentation and/or
// other materials provided with the distribution.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Package gocheckeres contains additional Checker implementations.
package gocheckers

import (
	gocheck "gopkg.in/check.v1"
)

// The GreaterThan checker verifies that the obtained value is strictly
// greater than provided length.
//
// For example:
//
//     c.Assert(len(list), GreaterThan, 5)
//
var GreaterThan gocheck.Checker = &greaterThanChecker{
	&gocheck.CheckerInfo{Name: "GreaterThan", Params: []string{"obtained", "n"}},
}

type greaterThanChecker struct {
	*gocheck.CheckerInfo
}

func (checker *greaterThanChecker) Check(params []interface{}, names []string) (result bool, err string) {
	v, ok := params[0].(int)
	if !ok {
		return false, "val must be an int"
	}

	n, ok := params[1].(int)
	if !ok {
		return false, "n must be an int"
	}
	return v > n, ""
}

// The LessThan checker verifies that the obtained value is strictly
// less than provided value.
//
// For example:
//
// 		c.Check(len(list), LessThan, 5)
//
var LessThan gocheck.Checker = &lessThanChecker{
	&gocheck.CheckerInfo{Name: "LessThan", Params: []string{"obtained", "n"}},
}

type lessThanChecker struct {
	*gocheck.CheckerInfo
}

func (checker *lessThanChecker) Check(params []interface{}, names []string) (result bool, err string) {
	gt, err := GreaterThan.Check(params, names)
	if err != "" {
		return gt, err
	}

	eq, err := gocheck.Equals.Check(params, names)
	if err != "" {
		return eq, err
	}

	return !(gt || eq), err
}
