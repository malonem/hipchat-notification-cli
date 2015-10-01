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
	message := flag.String("message", "", "message to send")

	flag.Parse()

	buf := bytes.NewBuffer(nil)
	stdin := os.Stdin
	if stdin != nil {
		pipeFile, err := stdin.Stat()
		if err != nil {
			panic(err)
		}
		if pipeFile.Mode()&os.ModeNamedPipe != 0 {
			io.Copy(buf, stdin)
			stdin.Close()
		}
		stdin_message := string(buf.Bytes())
		if len(stdin_message) > 0 {
			message = &stdin_message
		}
	}

	if len(*token) > 0 && len(*room) > 0 && len(*message) > 0 {
		client := hipchat.NewClient(*token)
		notifRq := &hipchat.NotificationRequest{Message: *message, Color: *color, Notify: *notify, MessageFormat: *format}

		resp, err := client.Room.Notification(*room, notifRq)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during room notification %q\n", err)
			fmt.Fprintf(os.Stderr, "Server returns %+v\n", resp)
			return
		}
	} else {
		fmt.Fprintln(os.Stderr, "Message, Token, and Room must be specified.")
	}

}
