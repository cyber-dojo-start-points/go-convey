package hiker

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_answer_is_three_digits(t *testing.T) {
	Convey("The answer is three digits", t, func() {
		So(answer(), ShouldBeGreaterThanOrEqualTo, 100)
