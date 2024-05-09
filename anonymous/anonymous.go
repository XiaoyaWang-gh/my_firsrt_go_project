package anonymous

import (
	"fmt"
	"math"
)

func init() {
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(361))
}
