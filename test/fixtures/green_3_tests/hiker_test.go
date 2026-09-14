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

func Test_answer_is_less_than_fifty(t *testing.T) {
	Convey("The answer is less than fifty", t, func() {
		So(answer(), ShouldBeLessThan, 50)
	})
}

func Test_answer_is_a_multiple_of_seven(t *testing.T) {
	Convey("The answer is a multiple of seven", t, func() {
		So(answer()%7, ShouldEqual, 0)
	})
}
