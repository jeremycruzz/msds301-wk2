package regression

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func LinearCoeffs(points []stats.Coordinate) (m, b float64, err error) {
	coords, err := stats.LinearRegression(points)
	if err != nil {
		return 0, 0, err
	}
	return getCoeffs(coords)
}

func getCoeffs(points []stats.Coordinate) (m, b float64, err error) {
	//Only need 2 coordinates to find slope
	p0 := points[0]
	p1 := points[1]

	//Find p1 where p1.x != p0.x
	for i := 1; i < len(points); i++ {
		if (p0.X-p1.X > .001) || (p0.X-p1.X < -.001) {
			break
		}
		if i+1 == len(points) {
			return 0, 0, fmt.Errorf("all points have the same X value")
		}
		p1 = points[i+1]
	}

	m = (p1.Y - p0.Y) / (p1.X - p0.X)

	//Find intercept with slope and x0,y0
	b = p0.Y - (m * p0.X)

	return m, b, nil
}
