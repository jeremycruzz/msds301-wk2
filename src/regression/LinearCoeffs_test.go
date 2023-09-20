package regression_test

import (
	"testing"

	"github.com/jeremycruzz/msds301-wk2.git/src/datasets"
	"github.com/jeremycruzz/msds301-wk2.git/src/regression"
	"github.com/montanaflynn/stats"
	"github.com/stretchr/testify/suite"
)

type StatsTestSuite struct {
	suite.Suite
	dataSet map[int][]stats.Coordinate
	delta   float64
}

// Driver code to run suite
func TestStatsSuite(t *testing.T) {
	suite.Run(t, new(StatsTestSuite))
}

// "BeforeEach"
func (s *StatsTestSuite) SetupTest() {
	s.dataSet = datasets.AnscombeQuartet
	s.delta = .001
}

func (s *StatsTestSuite) TestSet1() {
	m, b, err := regression.LinearCoeffs(s.dataSet[1])

	s.Nil(err)
	s.InDelta(m, 0.5001, s.delta)
	s.InDelta(b, 3.0001, s.delta)
}

func (s *StatsTestSuite) TestSet2() {
	m, b, err := regression.LinearCoeffs(s.dataSet[2])

	s.Nil(err)
	s.InDelta(m, 0.5000, s.delta)
	s.InDelta(b, 3.0009, s.delta)
}

func (s *StatsTestSuite) TestSet3() {
	m, b, err := regression.LinearCoeffs(s.dataSet[3])

	s.Nil(err)
	s.InDelta(m, 0.4997, s.delta)
	s.InDelta(b, 3.0025, s.delta)
}

func (s *StatsTestSuite) TestSet4() {
	m, b, err := regression.LinearCoeffs(s.dataSet[4])

	s.Nil(err)
	s.InDelta(m, 0.4999, s.delta)
	s.InDelta(b, 3.0017, s.delta)
}
