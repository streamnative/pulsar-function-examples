package main

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-logr/logr"
)

type Context struct {
	ctx              context.Context
	stub             ContextServiceClient
	logger           logr.Logger
	tenant           string
	namespace        string
	name             string
	function_id      string
	instance_id      string
	function_version string
	secrets_provider SecretsProvider
	input_topics     []string
	output_topic     string
	user_config      map[string]string
	secrets_map      map[string]string
	message_id       *MessageId
	message          *Record
}

func NewContext(ctx context.Context, tenant, namespace, name, function_id, function_version, instance_id string, input_topics []string,
	output_topic string, user_config map[string]string, secrets_map map[string]string, secrets_provider SecretsProvider,
	stub ContextServiceClient, logger logr.Logger) *Context {
	return &Context{
		ctx:              ctx,
		stub:             stub,
		logger:           logger,
		tenant:           tenant,
		namespace:        namespace,
		name:             name,
		function_id:      function_id,
		function_version: function_version,
		instance_id:      instance_id,
		input_topics:     input_topics,
		output_topic:     output_topic,
		user_config:      user_config,
		secrets_provider: secrets_provider,
		secrets_map:      secrets_map,
	}
}

func (c *Context) SetMessageId(message_id *MessageId) {
	c.message_id = message_id
}

func (c *Context) GetMessage() (*Record, error) {
	if c.message_id != nil {
		return c.message, nil
	}
	res, err := c.stub.CurrentRecord(c.ctx, c.message_id)
	if err != nil {
		return nil, err
	}
	c.message = res
	return res, nil
}

func (c *Context) GetMessageId() *MessageId {
	return c.message_id
}

func (c *Context) GetMessageKey() (string, error) {
	if c.message_id == nil {
		_, err := c.GetMessage()
		if err != nil {
			return "", err
		}
	}
	return c.message.Key, nil
}

func (c *Context) GetMessageEventTime() (int32, error) {
	if c.message_id == nil {
		_, err := c.GetMessage()
		if err != nil {
			return 0, err
		}
	}
	return c.message.EventTimestamp, nil
}

func (c *Context) GetMessageProperties() (map[string]string, error) {
	if c.message_id == nil {
		_, err := c.GetMessage()
		if err != nil {
			return nil, err
		}
	}
	properties := map[string]string{}
	err := json.Unmarshal([]byte(c.message.Properties), &properties)
	if err != nil {
		return nil, err
	}

	return properties, nil
}

func (c *Context) GetMessageTopic() (string, error) {
	if c.message_id == nil {
		_, err := c.GetMessage()
		if err != nil {
			return "", err
		}
	}
	return c.message.TopicName, nil
}

func (c *Context) GetMessagePartitionKey() (string, error) {
	if c.message_id == nil {
		_, err := c.GetMessage()
		if err != nil {
			return "", err
		}
	}
	return c.message.PartitionId, nil
}

func (c *Context) GetTenant() string {
	return c.tenant
}

func (c *Context) GetNamespace() string {
	return c.namespace
}

func (c *Context) GetFunctionName() string {
	return c.name
}

func (c *Context) GetFunctionID() string {
	return c.function_id
}

func (c *Context) GetFunctionVersion() string {
	return c.function_version
}

func (c *Context) GetInstanceID() string {
	return c.instance_id
}

func (c *Context) GetLogger() logr.Logger {
	return c.logger
}

func (c *Context) GetUserConfigValue(key string) string {
	if v, ok := c.user_config[key]; ok {
		return v
	} else {
		return ""
	}
}

func (c *Context) GetUserConfigMap() map[string]string {
	return c.user_config
}

func (c *Context) GetSecret(secret_name string) (*string, error) {
	if c.secrets_provider == nil {
		return nil, errors.New("secrets provider is nil")
	}
	if v, ok := c.secrets_map[secret_name]; ok {
		secret := c.secrets_provider.provide_secret(secret_name, v)
		return &secret, nil
	} else {
		return nil, errors.New("secret not found")
	}
}

func (c *Context) RecordMetrics(metric_name string, metric_value float64) error {
	_, err := c.stub.RecordMetrics(c.ctx, &MetricData{
		MetricName: metric_name,
		Value:      metric_value,
	})
	return err
}

func (c *Context) GetInputTopics() []string {
	return c.input_topics
}

func (c *Context) GetOutputTopic() string {
	return c.output_topic
}

// func GetOutputSerdeClassname

func (c *Context) Publish(topic string, payload []byte) error {
	_, err := c.stub.Publish(c.ctx, &PulsarMessage{
		Topic:   topic,
		Payload: payload,
	})
	return err
}

func (c *Context) IncrCounter(key string, amount int64) error {
	_, err := c.stub.IncrCounter(c.ctx, &IncrStateKey{
		Key:    key,
		Amount: amount,
	})
	return err
}

func (c *Context) GetCounter(key string) (int64, error) {
	res, err := c.stub.GetCounter(c.ctx, &StateKey{
		Key: key,
	})
	if err != nil {
		return 0, err
	}
	return res.Value, nil
}

func (c *Context) PutState(key string, value []byte) error {
	_, err := c.stub.PutState(c.ctx, &StateKeyValue{
		Key:   key,
		Value: value,
	})
	return err
}

func (c *Context) GetState(key string) ([]byte, error) {
	res, err := c.stub.GetState(c.ctx, &StateKey{
		Key: key,
	})
	if err != nil {
		return nil, err
	}
	return res.Value, nil
}

func (c *Context) DeleteState(key string) error {
	_, err := c.stub.DeleteState(c.ctx, &StateKey{
		Key: key,
	})
	return err
}
