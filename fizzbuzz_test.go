package fizzbuzz_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type FizzBuzzSuite struct {
	suite.Suite
}

func TestFizzBuzzSuite(t *testing.T) {
	suite.Run(t, new(FizzBuzzSuite))
}

func (s *FizzBuzzSuite) TestPlaceHolder() {
 	s.True(true)
}
