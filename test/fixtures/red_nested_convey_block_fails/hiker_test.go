package hiker

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Convey blocks nest, and each leaf is run in its own pass over the tree, so
// this one test function makes two runs and two assertions. The reporter names
// neither block: it prints one dot or x per leaf and locates a failure by file
// and line only, so the count of assertions is what says how deep the tree ran.
func Test_life_the_universe_and_everything(t *testing.T) {
	Convey("Given the answer", t, func() {
		Convey("It is life, the universe and everything", func() {
			So(answer(), ShouldEqual, 42)
		})
		Convey("It is even", func() {
			So(answer()%2, ShouldEqual, 0)
		})
	})
}
