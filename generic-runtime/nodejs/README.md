# GenericRuntime for NodeJs

To create a NodeJs kind generic-runtime function, you need to specify the `--classname ${nodejs_file_name}`, such as `--classname exclamation`.

The whole command looks like below:

```bash
bin/pulsar-admin functions create --py ${nodejs-file} --name ${function-name} --classname ${nodejs-file-with-out-.js-suffix} --inputs ${input-topic} --output ${output-topic} --custom-runtime-options '{"genericKind": "nodejs"}'
```

1. exclamation.js

This function file contains basic usage of Pulsar functions, including `log`, `state_store` and `publish`.

2. check_json.js

This function file shows how to use json schema in a NodeJs function, you need to specify the schema to `json` when creating function.

3. check_avro.js

This function file shows how to use avro schema in a NodeJs function, you need to specify the schema to `avro`, and the `typeClassName` to `check_avro.definitions` when creating function.