package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrefixSum(t *testing.T) {
	Convey("Test Sum", t, func() {
		Convey("normal test", func() {
			So(PrefixSum(1, 2, 3, 4, 5), ShouldEqual, []int{1, 3, 6, 10, 15})
		})
	})
}
