package pledge

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{
			"test",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	defer cleanup()
	setup()
	type args struct {
		method    string
		publisher bool
		conf      *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Positive case",
			args{
				AuthMethodJWT,
				true,
				&Config{
					Issuer:         "test",
					PublicKeyPath:  "public.pem",
					PrivateKeyPath: "private.pem",
				},
			},
			false,
		},
		{
			"Negative case 1",
			args{
				"jwt",
				true,
				&Config{
					Issuer:         "test",
					PublicKeyPath:  "public.pem",
					PrivateKeyPath: "private.pem",
				},
			},
			true,
		},
		{
			"Negative case 2",
			args{
				AuthMethodJWT,
				true,
				&Config{
					Issuer:         "test",
					PublicKeyPath:  "publicc.pem",
					PrivateKeyPath: "private.pem",
				},
			},
			false,
		},
		{
			"Negative case 3",
			args{
				AuthMethodJWT,
				true,
				&Config{
					Issuer:         "test",
					PublicKeyPath:  "public.pem",
					PrivateKeyPath: "privatee.pem",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.method, tt.args.publisher, tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func setup() {
	// generate key
	privatekey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publickey := &privatekey.PublicKey

	// dump private key to file
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privatePem, _ := os.Create("private.pem")
	// nolint
	pem.Encode(privatePem, privateKeyBlock)

	// dump public key to file
	publicKeyBytes, _ := x509.MarshalPKIXPublicKey(publickey)
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem, _ := os.Create("public.pem")
	// nolint
	pem.Encode(publicPem, publicKeyBlock)

	SetConfig(&Config{
		Issuer:         "Test",
		PublicKeyPath:  "public.pem",
		PrivateKeyPath: "private.pem",
	})
}

func cleanup() {
	os.Remove("public.pem")
	os.Remove("private.pem")
	cfg = nil
}
