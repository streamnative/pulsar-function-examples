# Generic Runtime

Besides of Java, Python and Golang language supportted by Pulsar functions,StreamNative also provides a GenericRuntime which can support most of languages in the world, such as NodeJs, Python, Golang and even WASM.

This directory contains examples for the GenericRuntime which you can use directly or learn how to develop your own functions based on the GenericRuntime.

> The GenericRuntime hasn't beed contributed to the upstream yet, so it can only run on StreamNative's cloud: https://console.streamnative.cloud/.
> To create a generic-runtime kind function, you need to pass below arguments to the `pulsar-admin`:
> `--py ${function_file} --custom-runtime-options '{"genericKind": "python|nodejs|executable|wasm"}'`
> the `${function_file}` can be a .py/.js/.wasm or an executable file, we use the `--py` here because Pulsar doesn't support passing a generic-kind function file argument like `--function-file` yet.