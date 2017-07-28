package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)


// ------------------- Structs for YAML Config Files ----------------------------------

// For env_keys_config.yaml
// Not to be manually edited
type EnvKeysConfig struct {
	Environments map[string]EnvKeys `yaml:"environments"`
}

// For env_config.yaml
// To be manually edited by the user
type EnvConfig struct {
	Environments map[string]EnvInfo `yaml:"environments"`
}

// To be used in env_keys_config.yaml
type EnvKeys struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`	// to be encrypted (with the user's password) and stored
}


// To be used in EnvConfig
type EnvInfo struct {
	APIManagerEndpoint     string `yaml:"api_manager_endpoint"`
	RegistrationEndpoint string `yaml:"registration_endpoint"`
	TokenEndpoint string `yaml:"token_endpoint"`
}
// ---------------- End of Structs for YAML Config Files ---------------------------------


// variables
var envConfig EnvConfig
var envKeysConfig EnvKeysConfig

// Validates the configuration file
func (envConfig *EnvConfig) validate() {
	//
}

/**
Load the Environments Configuration file from the config.yaml file. If the file is not there
create a new config.yaml file and add default values
Validates the configuration, if it exists
*/
func LoadEnvConfig(envLocalConfig string) /* EnvConfig */ {
}

// Returns a pointer to EnvConfig
func GetEnvConfig() *EnvConfig {
	if &envConfig == nil {
		HandleErrorAndExit("Env configuration is not available", nil)
	}
	return &envConfig
}

// Returns a pointer to EnvKeysConfig
func GetEnvKeysConfig() *EnvKeysConfig{
	if &envKeysConfig == nil {
		HandleErrorAndExit("EnvKeys configuration is not available", nil)
	}
	return &envKeysConfig
}

// Persists the given Env configuration
func WriteConfigFile(envConfig interface{}, envConfigFilePath string) {
	data, err := yaml.Marshal(&envConfig)
	if err != nil {
		HandleErrorAndExit("Unable to create Env Configuration.", err)
	}

	err = ioutil.WriteFile(envConfigFilePath, data, 0644)
	if err != nil {
		HandleErrorAndExit("Unable to create Env Configuration.", err)
	}
}

/*
env_keys_config.yaml (Programmatically edited)
===============
environments:
	dev:
		client_id: xxxxxxxxxx
		client_secret: xxxxxxxxxx
		refresh_token: xxxxxxxxxx

	staging:
		client_id: xxxxxxxxxx
		client_secret: xxxxxxxxxx
		refresh_token: xxxxxxxxxx
 */

/*
env_config.yaml (Manually edited)
===============
environments:
	dev:
		apim_endpoint: xxxxxxxxx
		registration_endpoint: xxxxxxxxxx
		token_endpoint: xxxxxxxxx

	staging:
		apim_endpoint: xxxxxxxxx
		registration_endpoint: xxxxxxxxxx
		token_endpoint: xxxxxxxxx
*/

