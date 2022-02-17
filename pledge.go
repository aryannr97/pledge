package pledge

import (
	"fmt"
)

// Config defines all the attributes required to set up pledge
type Config struct {
	Issuer         string `yaml:"issuer"`
	PublicKeyPath  string `yaml:"publickeypath"`
	PrivateKeyPath string `yaml:"privatekeypath"`
}

var cfg *Config

// GetConfig returns pledge level configurations
func GetConfig() *Config {
	return cfg
}

// SetConfig sets pledge level configurations
func SetConfig(conf *Config) {
	cfg = conf
}

// Auth is generic interface for authentication
type Auth interface {
	GenerateIdentitiy(interface{}) (interface{}, error)
	VerifyIdentity(args ...interface{}) (interface{}, error)
}

// New provides an instance of implementation of Auth interface according to requested auth method
func New(method string, publisher bool, conf *Config) (Auth, error) {
	SetConfig(conf)
	switch method {
	case AuthMethodJWT:
		return newJWTAuth(publisher), nil
	default:
		return nil, fmt.Errorf(pledgeErrorStr, "Invalid method")
	}
}
