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

func Test_digits_of_the_answer_sum_to_six(t *testing.T) {
	Convey("The digits of the answer sum to six", t, func() {
		So(answerDigitSum(), ShouldEqual, 6)
	})
}
