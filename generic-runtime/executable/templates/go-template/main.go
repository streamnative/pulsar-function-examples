package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
	"google.golang.org/grpc"
)

func newStderrLogger() logr.Logger {
	return funcr.New(func(prefix, args string) {
		if prefix != "" {
			fmt.Fprintf(os.Stderr, "%s:%s\n", prefix, args)
		} else {
			fmt.Fprintln(os.Stderr, args)
		}
	}, funcr.Options{})
}

func newPulsarLogger(ctx context.Context, stub ContextServiceClient, log_topic string) logr.Logger {
	return funcr.New(func(prefix, args string) {
		message := args
		if prefix != "" {
			message = fmt.Sprintf("%s:%s\n", prefix, args)
		}
		stub.Publish(ctx, &PulsarMessage{
			Topic:   log_topic,
			Payload: []byte(message),
		})
	}, funcr.Options{})
}

func main() {
	var tenant, namespace, name string
	var source, sink string
	var instance_id, function_id, function_version string
	var user_config string
	var secrets_map string
	var log_topic string
	flag.StringVar(&tenant, "tenant", "", "tenant of function")
	flag.StringVar(&namespace, "namespace", "", "namespace of function")
	flag.StringVar(&name, "name", "", "name of function")
	flag.StringVar(&source, "source", "", "the source spec(in json format)")
	flag.StringVar(&sink, "sink", "", "the sink spec(in json format)")
	flag.StringVar(&instance_id, "instance_id", "", "the instance id")
	flag.StringVar(&function_id, "function_id", "", "the function id")
	flag.StringVar(&function_version, "function_version", "", "the function version")
	flag.StringVar(&user_config, "user_config", "", "the user config(in json format)")
	flag.StringVar(&secrets_map, "secrets_map", "", "the secrets map(in json format)")
	flag.StringVar(&log_topic, "log_topic", "", "the log topic")
	flag.Parse()

	var secrets_provider_impl SecretsProvider = &EnvironmentBasedSecretsProvider{}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	work_dir := filepath.Dir(ex)
	channel, err := grpc.Dial("unix://"+work_dir+"/context.sock", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer channel.Close()
	stub := NewContextServiceClient(channel)

	user_config_map := make(map[string]string)
	if user_config != "" {
		err = json.Unmarshal([]byte(user_config), &user_config_map)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error unmarshal user configs:", err)
		}
	}

	secrets_map_map := make(map[string]string)
	if secrets_map != "" {
		err = json.Unmarshal([]byte(secrets_map), &secrets_map_map)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error unmarshal secrets map:", err)
		}
	}

	var logger logr.Logger
	if log_topic != "" {
		logger = newPulsarLogger(context.Background(), stub, log_topic)
	} else {
		logger = newStderrLogger()
	}

	ctx := NewContext(context.Background(), tenant, namespace, name, function_id, function_version, instance_id, []string{source}, sink, user_config_map, secrets_map_map, secrets_provider_impl, stub, logger)

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintln(os.Stderr, "Error reading from stdout:", err)
			}
			break
		}
		meta_length := line[0]

		if len(line) < int(meta_length+3) {
			os.Stdout.Write([]byte("error: meta length is too long"))
			os.Stdout.Write([]byte("\n"))
			continue
		}

		meta := strings.Split(string(line[1:meta_length+1]), "@")
		if len(meta) != 2 {
			os.Stdout.Write([]byte("error: meta length is not 2"))
			os.Stdout.Write([]byte("\n"))
			continue
		}
		//msg_id := meta[0]
		//topic := meta[1]
		ctx.SetMessageId(&MessageId{
			Id: meta[0],
		})

		// ignore the last `\n` byte
		msg := line[meta_length+1 : len(line)-1]
		if len(msg) == 0 {
			os.Stdout.Write([]byte("error: msg length is 0"))
			os.Stdout.Write([]byte("\n"))
			continue
		}

		result := Process(msg, ctx)

		result = bytes.Replace(result, []byte("\n"), []byte(""), -1)
		os.Stdout.Write(result)
		os.Stdout.Write([]byte("\n"))
	}
}
