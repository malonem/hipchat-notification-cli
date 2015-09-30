package main

import (
	"bytes"
	"fmt"
	"github.com/tbruyelle/hipchat-go/hipchat"
	"io"
	"os"
)

func main() {
	buf := bytes.NewBuffer(nil)
	f := os.Stdin
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe != 0 {
		io.Copy(buf, f)
		f.Close()
	}
	s := string(buf.Bytes())

	c := hipchat.NewClient("yCacsX0NsiniGeh0AkL0rmUdSxaLn1Pxz7YrTu00")
	notifRq := &hipchat.NotificationRequest{Message: s, Color: "purple", Notify: false, MessageFormat: "text"}

	resp, err := c.Room.Notification("m-test", notifRq)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during room notification %q\n", err)
		fmt.Fprintf(os.Stderr, "Server returns %+v\n", resp)
		return
	}
}
