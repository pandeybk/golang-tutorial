package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test1(t *testing.T) {
	Convey("Test main", t, func() {
		Convey("returnSomething can return a Something", func() {
			ret := returnSomething()
			So(ret, ShouldEqual, "Something1")
		})

		Convey("return sum of 2 integer", func() {
			ret := add(2, 4)
			So(ret, ShouldEqual, 6.0)
		})
	})
}
