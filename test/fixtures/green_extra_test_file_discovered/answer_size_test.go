package hiker

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// A convey block that passes prints only a dot, and a dot says nothing about
// which file it came from. The count of assertions is what shows this file was
// collected and run, so this block adds a second one.
func Test_answer_is_two_digits(t *testing.T) {
	Convey("The answer is two digits", t, func() {
		So(answer(), ShouldBeGreaterThanOrEqualTo, 10)
		So(answer(), ShouldBeLessThanOrEqualTo, 99)
	})
}
