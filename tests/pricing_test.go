package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

const targetUR = "http://localshot:8080"

type pricingTestSuite struct {
	suite.Suite
	client *http.Client
}

func (s pricingTestSuite) SetupTest() {
	s.client = http.DefaultClient
}

func (s pricingTestSuite) TearDownTest() {
	s.client.CloseIdleConnections()
}

func TestName(t *testing.T) {
	suite.Run(t, &pricingTestSuite{})
}

func (s pricingTestSuite) TestPricingApi(t *testing.T) {
	s.T().Run("", func(t *testing.T) {
		_, err := s.client.Get(targetUR)
		if err != nil {
			fmt.Println("")
		}
	})
}
