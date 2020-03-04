package lib

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLib(t *testing.T) {
	Convey("The library", t, func() {

		Convey("Adds one and two", func() {
			So(Add(1, 2), ShouldEqual, 3)
		})

		Convey("Adds two and three", func() {
			So(Add(2, 3), ShouldEqual, 5)
		})

		Convey("Adds nine and ten", func() {
			So(Add(9, 10), ShouldEqual, 19)
		})

	})
}
