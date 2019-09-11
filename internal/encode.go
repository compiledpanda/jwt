package internal

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/pkg/errors"
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
	Secret    string
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
		alg = jwt.NewHS256([]byte(opt.Secret))
	case "HS384":
		alg = jwt.NewHS384([]byte(opt.Secret))
	case "HS512":
		alg = jwt.NewHS512([]byte(opt.Secret))
	case "RS256":
		// TODO
	case "RS384":
		// TODO
	case "RS512":
		// TODO
	case "ES256":
		// TODO
	case "ES384":
		// TODO
	case "ES512":
		// TODO
	case "PS256":
		// TODO
	case "PS384":
		// TODO
	case "PS512":
		// TODO
	case "EdDSA":
		// TODO
	}

	token, err := jwt.Sign(opt.Payload, alg, header...)
	if err != nil {
		return "", errors.Wrap(err, "Could not sign token")
	}
	return string(token) + "\n", nil
}
