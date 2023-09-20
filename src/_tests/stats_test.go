package tests

import (
	"testing"

	"github.com/jeremycruzz/msds301-wk2.git/src/datasets"
	"github.com/montanaflynn/stats"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StatsTestSuite struct {
	suite.Suite
	dataSet map[string][]float64
}

//Driver code to run suite
func(t *testing.T) {
	suite.Run(t,new(StatsTestSuite))
}

//BeforeEach test
func(s *StatsTestSuite) SetupTest(){
	s.dataSet = datasets.AnscombeQuartet
}

func(s *StatsTestSuite) Test1(){
	assert.Equal(1,1)
}


