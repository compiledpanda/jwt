package internal

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/pkg/errors"
)

// ValidateOptions specifies options for encoding
type ValidateOptions struct {
	Issuer     string
	Subject    string
	Audience   []string
	Expiration time.Time
	NotBefore  time.Time
	IssuedAt   time.Time
	Jti        string
	Algorithm  string
	Secret     []byte
}

// Validate jwt
func Validate(token []byte, opt ValidateOptions) (string, error) {
	// Validators
	validators := []jwt.Validator{}
	if opt.Issuer != "" {
		validators = append(validators, jwt.IssuerValidator(opt.Issuer))
	}
	if opt.Subject != "" {
		validators = append(validators, jwt.SubjectValidator(opt.Subject))
	}
	if len(opt.Audience) > 0 {
		validators = append(validators, jwt.AudienceValidator(opt.Audience))
	}
	if !opt.Expiration.IsZero() {
		validators = append(validators, jwt.ExpirationTimeValidator(opt.Expiration))
	}
	if !opt.NotBefore.IsZero() {
		validators = append(validators, jwt.NotBeforeValidator(opt.NotBefore))
	}
	if !opt.IssuedAt.IsZero() {
		validators = append(validators, jwt.IssuedAtValidator(opt.IssuedAt))
	}
	if opt.Jti != "" {
		validators = append(validators, jwt.IDValidator(opt.Jti))
	}

	// Algorithm
	var alg jwt.Algorithm
	switch opt.Algorithm {
	case "HS256":
		alg = jwt.NewHS256(opt.Secret)
	case "HS384":
		alg = jwt.NewHS384(opt.Secret)
	case "HS512":
		alg = jwt.NewHS512(opt.Secret)
	case "RS256":
		key, err := parseRSAPublicKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewRS256(jwt.RSAPublicKey(key))
	case "RS384":
		key, err := parseRSAPublicKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewRS384(jwt.RSAPublicKey(key))
	case "RS512":
		key, err := parseRSAPublicKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewRS512(jwt.RSAPublicKey(key))
	case "ES256":
		key, err := parseECDSAPublicKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewES256(jwt.ECDSAPublicKey(key))
	case "ES384":
		key, err := parseECDSAPublicKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewES384(jwt.ECDSAPublicKey(key))
	case "ES512":
		key, err := parseECDSAPublicKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewES512(jwt.ECDSAPublicKey(key))
	case "PS256":
		key, err := parseRSAPublicKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewPS256(jwt.RSAPublicKey(key))
	case "PS384":
		key, err := parseRSAPublicKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewPS384(jwt.RSAPublicKey(key))
	case "PS512":
		key, err := parseRSAPublicKey(opt.Secret)
		if err != nil {
			return "", err
		}
		alg = jwt.NewPS512(jwt.RSAPublicKey(key))
	case "EdDSA":
		// TODO wait for 1.13 support
	}

	// Validate
	p := jwt.Payload{}
	v := jwt.ValidatePayload(&p, validators...)
	header, err := jwt.Verify(token, alg, &p, v)
	if err != nil {
		return "", errors.Wrap(err, "Verification Failed")
	}
	if opt.Algorithm != header.Algorithm {
		return "", errors.New("Algorithm Mismatch: token - " + header.Algorithm + ", requested - " + opt.Algorithm)
	}
	return "VALID\n", nil
}

func parseRSAPublicKey(key []byte) (*rsa.PublicKey, error) {
	bytes := dePem(key)

	// Try PKCS1
	pk, err := x509.ParsePKCS1PublicKey(bytes)
	if err == nil {
		return pk, nil
	}

	// Try PKIX -> RSA
	k, err := x509.ParsePKIXPublicKey(bytes)
	if err == nil {
		pk, ok := k.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("PKIX contained a non-RSA key")
		}
		return pk, nil
	}

	return nil, errors.New("Unknown RSA format")
}

func parseECDSAPublicKey(key []byte) (*ecdsa.PublicKey, error) {
	bytes := dePem(key)

	// Try PKIX -> ECDSA
	k, err := x509.ParsePKIXPublicKey(bytes)
	if err == nil {
		pk, ok := k.(*ecdsa.PublicKey)
		if !ok {
			return nil, errors.New("PKIX contained a non-ECDSA key")
		}
		return pk, nil
	}

	return nil, errors.New("Unknown ECDSA format")
}
