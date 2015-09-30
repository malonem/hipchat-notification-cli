package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/tbruyelle/hipchat-go/hipchat"
	"io"
	"os"
)

func main() {
	token := flag.String("token", "", "authentication token")
	room := flag.String("room", "", "room name")
	color := flag.String("color", "yellow", "message background color: green, red, purple, gray, random")
	notify := flag.Bool("notify", false, "notify users")
	format := flag.String("format", "text", "format rendered: text or html")

	flag.Parse()

	buf := bytes.NewBuffer(nil)
	f := os.Stdin
	var s string
	if f != nil {
		fi, err := f.Stat()
		if err != nil {
			panic(err)
		}
		if fi.Mode()&os.ModeNamedPipe != 0 {
			io.Copy(buf, f)
			f.Close()
		}
		s = string(buf.Bytes())
	}

	if len(*token) > 0 && len(*room) > 0 {
		c := hipchat.NewClient(*token)
		notifRq := &hipchat.NotificationRequest{Message: s, Color: *color, Notify: *notify, MessageFormat: *format}

		resp, err := c.Room.Notification(*room, notifRq)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during room notification %q\n", err)
			fmt.Fprintf(os.Stderr, "Server returns %+v\n", resp)
			return
		}
	}
}
