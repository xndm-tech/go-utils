package clock

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDay(t *testing.T) {
	Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
		So(GetTodayFormat(), ShouldEqual, "2020-09-15")
	})

	Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
		So(GetYesterdayFormat(), ShouldEqual, "2020-09-14")
	})

	Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
		So(GetWeekAgoFormat(), ShouldEqual, "2020-09-08")
	})

	Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
		So(GetMonthAgoFormat(), ShouldEqual, "2020-08-16")
	})
}
