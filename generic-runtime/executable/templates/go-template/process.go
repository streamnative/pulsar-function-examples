package main

import (
	"fmt"
	"strings"
)

func Process(msg []byte, context *Context) []byte {
	// 1. unmarshal []byte to your struct, use any schema you want
	payload := string(msg)

	// 2. do your logic
	logger := context.GetLogger()
	for _, word := range strings.Split(payload, " ") {
		context.IncrCounter(word, 1)
		count, _ := context.GetCounter(word)
		logger.Info(fmt.Sprintf("got word: %s for %d times", word, count))
	}
	context.Publish("persistent://public/default/test-exec-package-serde-extra", []byte(payload+"!!!"))
	data := payload + "!"

	// 3. marshal your struct to []byte
	return []byte(data)
}
