package fizzbuzz_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExampleSuite struct {
	suite.Suite
}

func TestExampleSuite(t *testing.T) {
	suite.Run(t, new(ExampleSuite))
}

func (s *ExampleSuite) TestBooleans() {
 	s.Equal(true, true)
}
