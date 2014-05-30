// Package gocheckeres contains additional Checker implementations.
package gocheckers

import "launchpad.net/gocheck"

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

func (checker *greaterThanChecker) Check(params []interface{}, names []string) (result bool, error string) {
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
