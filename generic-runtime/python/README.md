# GenericRuntime for Python

The usage of Python kind generic-runtime functions is exactly the same with the usage of upstream python functions.

The whole command looks like below:

```bash
bin/pulsar-admin functions create --py ${python-file} --name ${function-name} --classname ${class-name} --inputs ${input-topic} --output ${output-topic} --custom-runtime-options '{"genericKind": "python"}'
```

1. exclamation.py

This function file contains basic usage of Pulsar functions, including `log`, `state_store` and `publish`.

2. check_json.py

This function file shows how to use json schema in a Python function, you need to specify the schema to `json`, and the `typeClassName` to `checkjson.Student` when creating function.

3. check_avro.py

This function file shows how to use avro schema in a Python function, you need to specify the schema to `avro`, and the `typeClassName` to `check_avro.Student` when creating function.