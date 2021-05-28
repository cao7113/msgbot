package lark

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestLarkSuite(t *testing.T) {
	suite.Run(t, &LarkSuite{})
}

type LarkSuite struct {
	suite.Suite
}
