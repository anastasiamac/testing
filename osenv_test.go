// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package testing_test

import (
	"os"

	gc "gopkg.in/check.v1"

	"github.com/juju/testing"
)

type osEnvSuite struct {
	osEnvSuite testing.OsEnvSuite
}

var _ = gc.Suite(&osEnvSuite{})

func (s *osEnvSuite) SetUpSuite(c *gc.C) {
	s.osEnvSuite = testing.OsEnvSuite{}
}

func (s *osEnvSuite) TestOriginalEnvironment(c *gc.C) {
	// The original environment is properly cleaned and restored.
	err := os.Setenv("TESTING_OSENV_ORIGINAL", "original-value")
	c.Assert(err, gc.IsNil)
	s.osEnvSuite.SetUpSuite(c)
	c.Assert(os.Getenv("TESTING_OSENV_ORIGINAL"), gc.Equals, "")
	s.osEnvSuite.TearDownSuite(c)
	// The environment has been restored.
	c.Assert(os.Getenv("TESTING_OSENV_ORIGINAL"), gc.Equals, "original-value")
}

func (s *osEnvSuite) TestTestingEnvironment(c *gc.C) {
	// Environment variables set up by tests are properly removed.
	s.osEnvSuite.SetUpSuite(c)
	s.osEnvSuite.SetUpTest(c)
	err := os.Setenv("TESTING_OSENV_NEW", "new-value")
	c.Assert(err, gc.IsNil)
	s.osEnvSuite.TearDownTest(c)
	s.osEnvSuite.TearDownSuite(c)
	c.Assert(os.Getenv("TESTING_OSENV_NEW"), gc.Equals, "")
}
