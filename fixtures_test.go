package fixtures

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFixturesSearcherFilter(t *testing.T) {
	Convey("Should return error, if we pass non-existing path to file", t, func() {
		fixture, err := GetFixture("test_fixtures1.yml")
		So(err, ShouldNotBeNil)
		So(fixture, ShouldBeNil)

	})
	Convey("Should return error, if we can't parse file", t, func() {
		fixture, err := GetFixture("test_fixtures2.yml")
		So(err, ShouldNotBeNil)
		So(fixture, ShouldBeNil)

	})
	Convey("Should fetch fixtures", t, func() {
		fixture, err := GetFixture("test_fixtures.yml")
		So(err, ShouldBeNil)
		Convey("Shoud panic if trying to filter with unexisting filter", func() {
			defer func() {
				So(recover(), ShouldNotBeNil)
			}()
			fixture.Filter("id", 90, "some_value").All()
		})
		Convey("Shoud skip fixtures, if it doesn't have filtered field", func() {
			c := fixture.Filter("unexisting_feld", Equal, "some_value").Count()
			So(c, ShouldEqual, 0)
		})
		Convey("Shoud return error if we're trying to fetch first unexisting fixture", func() {
			c, err := fixture.Filter("A", Equal, "unexisting_value").First()
			So(err, ShouldNotBeNil)
			So(c, ShouldBeNil)
		})
		Convey("Shoud fetch first fixture with A == asdasd", func() {
			c, err := fixture.Filter("A", Equal, "asdasd").First()
			So(err, ShouldBeNil)
			So(c["id"], ShouldEqual, "1")
		})
		Convey("Shoud fetch count if fixtures with A == asdasd", func() {
			c := fixture.Filter("A", Equal, "asdasd").Count()
			So(c, ShouldEqual, 2)
		})
		Convey("Shoud fetch fixtures with A == asdasd", func() {
			c := fixture.Filter("A", Equal, "asdasd").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldEqual, "asdasd")
			}
		})

		Convey("Shoud fetch fixtures with A != asdasd", func() {
			c := fixture.Filter("A", NotEqual, "asdasd").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldNotEqual, "asdasd")
			}
		})

		Convey("Shoud fetch fixtures which starts with 'cc'", func() {
			c := fixture.Filter("A", StartsWith, "cc").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldStartWith, "cc")
			}
		})
		Convey("Shoud fetch fixtures which not starts with 'cc'", func() {
			c := fixture.Filter("A", NotStartsWith, "cc").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldNotStartWith, "cc")
			}
		})
		Convey("Shoud fetch fixtures which starts with 'aa'", func() {
			c := fixture.Filter("A", EndsWith, "aa").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldEndWith, "aa")
			}
		})
		Convey("Shoud fetch fixtures which not ends with 'aa'", func() {
			c := fixture.Filter("A", NotEndsWith, "aa").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldNotEndWith, "cc")
			}
		})
		Convey("Shoud fetch fixtures which contains 'wer'", func() {
			c := fixture.Filter("A", Contains, "wer").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldContainSubstring, "wer")
			}
		})
		Convey("Shoud fetch fixtures which not contains 'wer'", func() {
			c := fixture.Filter("A", NotContains, "wer").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldNotContainSubstring, "wer")
			}
		})
	})

}
