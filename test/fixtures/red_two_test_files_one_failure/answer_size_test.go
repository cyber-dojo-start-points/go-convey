package hiker

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_answer_is_two_digits(t *testing.T) {
	Convey("The answer is two digits", t, func() {
		So(answer(), ShouldBeGreaterThanOrEqualTo, 10)
	})
}

func Test_answer_is_three_digits(t *testing.T) {
	Convey("The answer is three digits", t, func() {
		So(answer(), ShouldBeGreaterThanOrEqualTo, 100)
	})
}
