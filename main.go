package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/tbruyelle/hipchat-go/hipchat"
	"io"
	"os"
)

var (
	token   = flag.String("token", "", "authentication token")
	room    = flag.String("room", "", "room name")
	color   = flag.String("color", "yellow", "message background color: green, red, purple, gray, random")
	notify  = flag.Bool("notify", false, "notify users")
	format  = flag.String("format", "text", "format rendered: text or html")
	message = flag.String("message", "", "message to send")
)

func readMessageFromStdin() (*string, error) {
	buf := bytes.NewBuffer(nil)
	stdin := os.Stdin
	if stdin != nil {
		pipeFile, err := stdin.Stat()
		if err != nil {
			return nil, err
		}
		if pipeFile.Mode()&os.ModeNamedPipe != 0 {
			io.Copy(buf, stdin)
			stdin.Close()
		}
		stdinMessage := string(buf.Bytes())
		if len(stdinMessage) > 0 {
			return &stdinMessage, nil
		}
	}
	return nil, nil
}

func isValidRequest(t *string, r *string, m *string) bool {
	return len(*t) > 0 && len(*r) > 0 && len(*m) > 0
}

func main() {

	flag.Parse()

	notification, _ := readMessageFromStdin()
	if notification == nil {
		notification = message
	}

	if isValidRequest(token, room, notification) {
		client := hipchat.NewClient(*token)

		request := hipchat.NotificationRequest{
			Message:       *notification,
			Color:         *color,
			Notify:        *notify,
			MessageFormat: *format,
		}

		response, err := client.Room.Notification(*room, &request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during room notification %q\n", err)
			fmt.Fprintf(os.Stderr, "Server returns %+v\n", response)
			return
		}
	} else {
		fmt.Fprintln(os.Stderr, "Message, Token, and Room must be specified.")
	}
}
