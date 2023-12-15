# GenericRuntime for WASM

The support of WASM is still in eayly stage and can be changed fast in the future.

It uses the `wasmedge-bindgen` and `wasmedge-bindgen-macro` to support WASM in Pulsar functions.

## excla

Run `cargo build --target wasm32-wasi --release` in the `excla` dir and you will get a `excla.wasm` under `target/wasm32-wasi/release`.

Then you can create a function with below command:

```bash
bin/pulsar-admin functions create --py ${wasm-file} --name ${function-name} --classname process --inputs ${input-topic} --output ${output-topic} --custom-runtime-options '{"genericKind": "wasm"}'
```

or

```bash
bin/pulsar-admin functions create --py ${wasm-file} --name ${function-name} --classname process_json --inputs ${input-topic} --output ${output-topic} --custom-runtime-options '{"genericKind": "wasm"}'
```