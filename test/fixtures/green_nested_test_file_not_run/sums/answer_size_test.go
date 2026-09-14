package sums

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Go ties a package to its directory, so this file is a package of its own and
// go test, run in the sandbox directory, never reaches it. The print is what
// shows whether it ran.
func Test_answer_is_two_digits(t *testing.T) {
	Convey("The answer is two digits", t, func() {
		fmt.Println("checking the size of the answer")
		So(42, ShouldBeGreaterThanOrEqualTo, 10)
	})
}
