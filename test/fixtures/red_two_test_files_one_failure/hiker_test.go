package hiker

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_life_the_universe_and_everything(t *testing.T) {
	Convey("The answer is life, the universe and everything", t, func() {
		So(answer(), ShouldEqual, 42)
	})
}

func Test_answer_is_even(t *testing.T) {
	Convey("The answer is even", t, func() {
		So(answer()%2, ShouldEqual, 0)
	})
}
