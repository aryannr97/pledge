package pledge

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	gojwt "github.com/dgrijalva/jwt-go"
)

const (
	cutset = "Bearer "
)

// jwtAuth is type for JWT based implementation of Auth interface
type jwtAuth struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

// newJWTAuth retuns a new instance of jwtAuth
func newJWTAuth(publisher bool) *jwtAuth {
	auth := &jwtAuth{}
	auth.loadKeys(publisher)
	return auth
}

// loadKeys loads public/private keys for token generation/verification
// private key is loaded only if publisher value is true
func (auth *jwtAuth) loadKeys(publisher bool) {
	if publisher {
		privKeyBytes, err := ioutil.ReadFile(cfg.PrivateKeyPath)
		if err != nil {
			log.Println(err)
		}

		privKey, err := gojwt.ParseRSAPrivateKeyFromPEM(privKeyBytes)
		if err != nil {
			log.Println(err)
		}

		auth.PrivateKey = privKey
	}

	pubKeyBytes, err := ioutil.ReadFile(cfg.PublicKeyPath)
	if err != nil {
		log.Println(err)
	}

	pubKey, err := gojwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {
		log.Println(err)
	}

	auth.PublicKey = pubKey
}

// GenerateIdentitiy generates JWT token with parameterized data
func (auth *jwtAuth) GenerateIdentitiy(data interface{}) (interface{}, error) {
	claims := make(gojwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["data"] = data
	claims["iss"] = cfg.Issuer

	token, err := gojwt.NewWithClaims(gojwt.SigningMethodRS256, claims).SignedString(auth.PrivateKey)

	return token, err
}

// VerifyIdentity extracts token from Authorization header &  confirms the validity
func (auth *jwtAuth) VerifyIdentity(args ...interface{}) (interface{}, error) {
	var token string
	for _, arg := range args {
		token = arg.(string)
	}

	token = strings.Replace(token, cutset, "", 1)

	jwtToken, err := gojwt.Parse(token, func(t *gojwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*gojwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf(pledgeErrorStr, fmt.Sprintf("unsupported signing method: %s", t.Header["alg"]))
		}

		return auth.PublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(gojwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, fmt.Errorf(pledgeErrorStr, "validate => invalid")
	}

	return claims, nil
}
