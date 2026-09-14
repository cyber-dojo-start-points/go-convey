package fizzbuzz

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_multiples_of_three_and_five_are_fizzbuzz(t *testing.T) {
	Convey("Multiples of three and five are FizzBuzz", t, func() {
		So(fizzBuzz(15), ShouldEqual, "FizzBuzz")
	})
}
