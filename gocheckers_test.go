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

package gocheckers

import (
	"reflect"
	"testing"

	"launchpad.net/gocheck"
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type CheckersTestSuite struct{}

var _ = gocheck.Suite(&CheckersTestSuite{})

func testInfo(c *gocheck.C, checker gocheck.Checker, name string, paramNames []string) {
	info := checker.Info()
	if info.Name != name {
		c.Fatalf("Got name %s, expected %s", info.Name, name)
	}
	if !reflect.DeepEqual(info.Params, paramNames) {
		c.Fatalf("Got param names %#v, expected %#v", info.Params, paramNames)
	}
}

func testCheck(c *gocheck.C, checker gocheck.Checker, result bool, error string, params ...interface{}) ([]interface{}, []string) {
	info := checker.Info()
	if len(params) != len(info.Params) {
		c.Fatalf("unexpected param count in test; expected %d got %d", len(info.Params), len(params))
	}
	names := append([]string{}, info.Params...)
	result_, error_ := checker.Check(params, names)
	if result_ != result || error_ != error {
		c.Fatalf("%s.Check(%#v) returned (%#v, %#v) rather than (%#v, %#v)",
			info.Name, params, result_, error_, result, error)
	}
	return params, names
}

func (s *CheckersTestSuite) TestGreaterThan(c *gocheck.C) {
	testInfo(c, GreaterThan, "GreaterThan", []string{"obtained", "n"})

	testCheck(c, GreaterThan, true, "", len([]int{1, 2, 3}), 2)
	testCheck(c, GreaterThan, false, "", 2, 6)
	testCheck(c, GreaterThan, false, "val must be an int", "", 0)
	testCheck(c, GreaterThan, false, "val must be an int", []int{0}, 0)
}

func (s *CheckersTestSuite) TestLessThan(c *gocheck.C) {
	testInfo(c, LessThan, "LessThan", []string{"obtained", "n"})

	testCheck(c, LessThan, true, "", len([]int{1, 2, 3}), 5)
	testCheck(c, LessThan, false, "", 2, 0)
	testCheck(c, LessThan, false, "val must be an int", "", 0)
	testCheck(c, LessThan, false, "val must be an int", []int{0}, 0)
}
