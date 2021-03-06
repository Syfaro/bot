package bot

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	DefaultChannel = "#go-bot"
	DefaultUser    = &User{Nick: "user123"}
	DefaultCommand = "command"
	DefaultRawArgs = "arg1  arg2"
	DefaultArgs    = []string{
		"arg1",
		"arg2",
	}
)

func TestPaser(t *testing.T) {
	Convey("When given a message", t, func() {
		Convey("When the message is empty", func() {
			cmd := parse("", DefaultChannel, DefaultUser)

			So(cmd, ShouldBeNil)
		})

		Convey("When the message doesn't have the prefix", func() {
			Message := "regular message"
			cmd := parse(Message, DefaultChannel, DefaultUser)

			So(cmd, ShouldBeNil)
		})

		Convey("When the message is only the prefix", func() {
			cmd := parse(CmdPrefix, DefaultChannel, DefaultUser)

			So(cmd, ShouldBeNil)
		})

		Convey("When the message is valid command", func() {
			msg := fmt.Sprintf("%v%v", CmdPrefix, DefaultCommand)
			cmd := parse(msg, DefaultChannel, DefaultUser)

			So(cmd, ShouldNotBeNil)
			So(cmd.Command, ShouldEqual, DefaultCommand)
			So(cmd.Channel, ShouldEqual, DefaultChannel)
		})

		Convey("When the message is a command with args", func() {
			msg := fmt.Sprintf("%v%v %v", CmdPrefix, DefaultCommand, DefaultRawArgs)
			cmd := parse(msg, DefaultChannel, DefaultUser)

			So(cmd, ShouldNotBeNil)
			So(cmd.Command, ShouldEqual, DefaultCommand)
			So(cmd.Channel, ShouldEqual, DefaultChannel)
			So(cmd.Args, ShouldResemble, DefaultArgs)
			So(cmd.RawArgs, ShouldEqual, DefaultRawArgs)
		})

		Convey("When the message has extra spaces", func() {
			msg := fmt.Sprintf(" %v %v %v  %v  ", CmdPrefix, DefaultCommand, DefaultArgs[0], DefaultArgs[1])
			cmd := parse(msg, DefaultChannel, DefaultUser)

			So(cmd, ShouldNotBeNil)
			So(cmd.Command, ShouldEqual, DefaultCommand)
			So(cmd.Channel, ShouldEqual, DefaultChannel)
			So(cmd.Args, ShouldResemble, DefaultArgs)
			So(cmd.RawArgs, ShouldEqual, DefaultRawArgs)
		})
	})
}
