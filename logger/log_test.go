package log

import (
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

func TestNewLogger(t *testing.T) {
	uuid := 0

	stubUuidMaker := func() string {
		uuid += 1
		return strconv.Itoa(uuid)
	}

	Convey("GIVEN we instantiate a new log", t, func() {
		sut := NewLog("fake-app", stubUuidMaker)

		Convey("AND the application name is set", func() {
			So(sut.ApplicationName, ShouldEqual, "fake-app")
		})
	})
}

func TestLogger_NewContext(t *testing.T) {
	uuid := 0

	stubUuidMaker := func() string {
		uuid += 1
		return strconv.Itoa(uuid)
	}

	Convey("GIVEN we instantiate a new log", t, func() {
		sut := NewLog("some-app", stubUuidMaker)

		Convey("WHEN we call the OpenContext() method", func() {
			sut.OpenContext()

			Convey("THEN a LogContext is added to the log Calls field with the correct fields", func() {
				So(sut.Calls, ShouldHaveLength, 1)
				So(sut.Calls[0].Order, ShouldEqual, 0)
				So(sut.Calls[0].Calls, ShouldHaveLength, 0)

				Convey("AND we call the OpenContext() method again", func() {
					sut.OpenContext()
					So(sut.Calls[0].Calls, ShouldHaveLength, 1)
					So(sut.Calls[0].Calls[0].Order, ShouldEqual, 0)
					So(sut.Calls[0].Calls[0].Calls, ShouldHaveLength, 0)
				})
			})
		})
	})
}
