package main

import (
	"fmt"

	"github.com/jeremycruzz/msds301-wk2.git/src/datasets"
	"github.com/jeremycruzz/msds301-wk2.git/src/regression"
)

func main() {
	dataSet := datasets.AnscombeQuartet
	for i := range dataSet {
		m, b, err := regression.LinearCoeffs(dataSet[i])
		if err != nil {
			panic("ah")
		}
		fmt.Printf("m=%.5f b=%.5f\n", m, b)
	}
}
