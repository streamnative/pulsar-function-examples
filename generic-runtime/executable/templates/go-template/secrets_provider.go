package main

import "os"

type SecretsProvider interface {
	init(config map[string]string)
	provide_secret(secret_name, secret_path string) string
}

type ClearTextSecretsProvider struct {
}

func (s *ClearTextSecretsProvider) init(config map[string]string) {
}

func (s *ClearTextSecretsProvider) provide_secret(secret_name, secret_path string) string {
	return secret_path
}

type EnvironmentBasedSecretsProvider struct{}

func (s *EnvironmentBasedSecretsProvider) init(config map[string]string) {
}

func (s *EnvironmentBasedSecretsProvider) provide_secret(secret_name, secret_path string) string {
	return os.Getenv(secret_name)
}
