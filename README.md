# Go Checkers

These are additional gocheck Checker implementations that can be used in addition to the ones provided by [gocheck](http://labix.org/gocheck).

## Current state: deprecated

Only used by v1 go services, which are currently being made deprecated. We don't think this brings enough value to keep around.

## Installation

    $ go get gopkg.in/check.v1
    $ go get -u "github.com/CrowdSurge/gocheckers"
    
## Basic Example

    package main

    import (
    	"testing"    

    	chk "github.com/CrowdSurge/gocheckers"
    	. "launchpad.net/gocheck"
    )

    func Test(t *testing.T) { TestingT(t) }
    
    type TestSuite struct{}
    
    var _ = Suite(&TestSuite{})

    func (suite *TestSuite) TestGetTemplatesSuccessfully(c *C) {
        templates, err := GetTemplates("Testing")
    	c.Assert(err, IsNil)
    	c.Check(len(templates), chk.GreaterThan, 0)
    }

