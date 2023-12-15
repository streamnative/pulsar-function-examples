## Usage

Update the `Process` method in the process.go:

```go
func Process(msg []byte, context *Context) []byte {
	// 1. unmarshal []byte to your struct, use any schema you want
	payload := string(msg)

	// 2. do your logic
	logger := context.GetLogger()
	logger.Info(fmt.Sprintf("processing message in function %s/%s/%s", context.GetTenant(), context.GetNamespace(), context.GetFunctionName()))
	for _, word := range strings.Split(payload, " ") {
		context.IncrCounter(word, 1)
		count, _ := context.GetCounter(word)
		logger.Info(fmt.Sprintf("got word: %s for count: %d times", word, count))
	}
	context.Publish("persistent://public/default/test-exec-extra", []byte(payload+"!!!"))
	data := payload + "!"

	// 3. marshal your struct to []byte
	return []byte(data)
}
```

## Build

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o excla *.go
```

Now you can use the `excla` binary to create a function.
