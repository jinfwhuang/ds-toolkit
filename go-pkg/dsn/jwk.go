package dsn

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

// Generate byte representation of a JSON Web Key.
func GenerateJWK() ([]byte, error) {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return RSAToJWK(privatekey)
}

// Convert RSA keypair to a JSON Web Key.
func RSAToJWK(privatekey *rsa.PrivateKey) ([]byte, error) {
	jwkKey, err := jwk.FromRaw(privatekey)
	if err != nil {
		return nil, err
	}
	jwkKey.Set("ext", true)
	jwk, err := json.Marshal(jwkKey)
	if err != nil {
		return nil, err
	}
	return jwk, nil
}
