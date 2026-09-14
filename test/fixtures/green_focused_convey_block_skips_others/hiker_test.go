package hiker

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Focus has to reach the root to take effect: a FocusConvey nested inside a
// plain Convey runs its siblings anyway. Rooted here, only the focused child
// runs and the sibling asserting three digits never does, so the suite passes
// while saying nothing about it. The reporter prints one dot per leaf it ran
// and no word at all about the leaf it skipped, so the assertion count is the
// only trace, and this output cannot be told from a suite of one passing
// block.
func Test_life_the_universe_and_everything(t *testing.T) {
	FocusConvey("Given the answer", t, func() {
		FocusConvey("It is life, the universe and everything", func() {
			So(answer(), ShouldEqual, 42)
		})
		Convey("It is three digits", func() {
			So(answer(), ShouldBeGreaterThanOrEqualTo, 100)
		})
	})
}
