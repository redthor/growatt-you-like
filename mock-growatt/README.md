## Mock Growatt

A simple TCP application to emit messages similar to the Growatt
server.

```shell
# For Help
go run main.go -h

# Run with defaults
go run main.go
```

### Test Messages
Any files with `*.txt` suffixes will be read from the `./messages` folder. Each line is a message.