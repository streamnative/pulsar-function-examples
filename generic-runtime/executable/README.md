# GenericRuntime for Executable

The support of compiled language is different with the way to support NodeJs and Python in GenericRuntime, since a compiled language program needs to be compiled first before executing, and the compile may take lots of time, the GenericRuntime doesn't support uploading a source code file for compiled language directly but only support users compiling their function locally and then uploading the compiled&executable file as the function file.

The whole command looks like below:

```bash
bin/pulsar-admin functions create --py ${executable-file} --name ${function-name} --classname ${executable-file} --inputs ${input-topic} --output ${output-topic} --custom-runtime-options '{"genericKind": "executable"}'
```

We provided a Golang template as an example to demonstrate how to write your own "executable" kind generic-runtime function. You can also write with other languages you like, just make sure it has same logic with the go-template.

> The `context_grpc.pb.go` and `context.pb.go` in the `templates/go-template` is generated from the `templates/context.proto` file.
