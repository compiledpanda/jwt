package internal

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

// Algorithms contains the list of accepted algorithms
var Algorithms = []string{"HS256", "HS384", "HS512", "RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512", "EdDSA"}

// ValidAlgorithm returns true for valid algorithm
func ValidAlgorithm(algorithm string) bool {
	for _, a := range Algorithms {
		if a == algorithm {
			return true
		}
	}
	return false
}

// EncodeOptions specifies options for encoding
type EncodeOptions struct {
	Header    map[string]string
	Payload   map[string]interface{}
	Algorithm string
	Secret    []byte
}

// Encode jwt
func Encode(opt EncodeOptions) (str string, err error) {

	header := []jwt.SignOption{}
	for k, v := range opt.Header {
		switch k {
		case "kid":
			header = append(header, jwt.KeyID(v))
		case "cty":
			header = append(header, jwt.ContentType(v))
		}
	}

	var alg jwt.Algorithm
	switch opt.Algorithm {
	case "HS256":
		alg = jwt.NewHS256(opt.Secret)
	case "HS384":
		alg = jwt.NewHS384(opt.Secret)
	case "HS512":
		alg = jwt.NewHS512(opt.Secret)
	case "RS256":
		key, err := parseRSAPrivateKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewRS256(jwt.RSAPrivateKey(key))
	case "RS384":
		key, err := parseRSAPrivateKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewRS384(jwt.RSAPrivateKey(key))
	case "RS512":
		key, err := parseRSAPrivateKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewRS512(jwt.RSAPrivateKey(key))
	case "ES256":
		key, err := parseECDSAPrivateKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewES256(jwt.ECDSAPrivateKey(key))
	case "ES384":
		key, err := parseECDSAPrivateKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewES384(jwt.ECDSAPrivateKey(key))
	case "ES512":
		key, err := parseECDSAPrivateKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewES512(jwt.ECDSAPrivateKey(key))
	case "PS256":
		key, err := parseRSAPrivateKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewPS256(jwt.RSAPrivateKey(key))
	case "PS384":
		key, err := parseRSAPrivateKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewPS384(jwt.RSAPrivateKey(key))
	case "PS512":
		key, err := parseRSAPrivateKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewPS512(jwt.RSAPrivateKey(key))
	case "EdDSA":
		// TODO wait for 1.13 support
	}

	token, err := jwt.Sign(opt.Payload, alg, header...)
	if err != nil {
		return "", errors.Wrap(err, "Could not sign token")
	}
	return string(token) + "\n", nil
}

func parseRSAPrivateKey(key []byte) (*rsa.PrivateKey, error) {
	bytes := dePem(key)

	// Try PKCS8 -> RSA
	k, err := x509.ParsePKCS8PrivateKey(bytes)
	if err == nil {
		pk, ok := k.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("PKCS8 contained a non-RSA key")
		}
		return pk, nil
	}

	// Try PKCS1
	pk, err := x509.ParsePKCS1PrivateKey(bytes)
	if err == nil {
		return pk, nil
	}

	// Try OPENSSH
	sk, err := ssh.ParseRawPrivateKey([]byte(key))
	if err == nil {
		pk, ok := sk.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("OPENSSH contained a non-RSA key")
		}
		return pk, nil
	}

	return nil, errors.New("Unknown RSA format")
}

func parseECDSAPrivateKey(key []byte) (*ecdsa.PrivateKey, error) {
	bytes := dePem(key)

	// Try ECPrivateKey
	pk, err := x509.ParseECPrivateKey(bytes)
	if err == nil {
		return pk, nil
	}

	// Try PKCS8 -> ECDSA
	k, err := x509.ParsePKCS8PrivateKey(bytes)
	if err == nil {
		pk, ok := k.(*ecdsa.PrivateKey)
		if !ok {
			return nil, errors.New("PKCS8 contained a non-ECDSA key")
		}
		return pk, nil
	}

	return nil, errors.New("Unknown ECDSA format")
}
