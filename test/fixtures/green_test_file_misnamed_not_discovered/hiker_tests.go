package hiker

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// go test collects test functions from files whose names end _test.go. This
// file ends _tests.go, so it is compiled as ordinary source and this function
// is never called. A convey block only runs when its test function calls it,
// so nothing here reaches the reporter. It asserts three digits, which would
// fail if it ran.
func Test_answer_is_three_digits(t *testing.T) {
	Convey("The answer is three digits", t, func() {
		So(answer(), ShouldBeGreaterThanOrEqualTo, 100)
	})
}
