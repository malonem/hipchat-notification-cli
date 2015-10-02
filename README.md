# hipchat-notification-cli
This cli was created to make it easy to send messages to HipChat rooms using the command line.

# Example

Sending 'test' as a message using the `-message` parameter.
```bash
./hipchat-notification-cli -room=<room> -token=<token> -message=test
```

Sending 'test' passing message through a `pipe`.
```bash
echo "test" | ./hipchat-notification-cli -room=<room> -token=<token>
```
