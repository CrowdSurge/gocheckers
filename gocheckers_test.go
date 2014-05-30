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
}
