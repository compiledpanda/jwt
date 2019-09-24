package main

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	type test struct {
		args []string
	}

	tests := []test{
		// No Args
		{args: []string{""}},
		// Encode HMAC
		{args: []string{"encode", "-a", "HS256", "-s", "secret"}},
		{args: []string{"encode", "-a", "HS384", "-s", "secret"}},
		{args: []string{"encode", "-a", "HS512", "-s", "secret"}},
		// Encode RSA
		{args: []string{"encode", "-a", "RS256", "-s", "@./test/rsa_pkcs1_private.pem"}},
		{args: []string{"encode", "-a", "RS256", "-s", "@./test/rsa_pkcs1_private.der"}},
		{args: []string{"encode", "-a", "RS256", "-s", "@./test/rsa_pkcs8_private.pem"}},
		{args: []string{"encode", "-a", "RS256", "-s", "@./test/rsa_pkcs8_private.der"}},
		{args: []string{"encode", "-a", "RS384", "-s", "@./test/rsa_pkcs1_private.pem"}},
		{args: []string{"encode", "-a", "RS512", "-s", "@./test/rsa_pkcs1_private.pem"}},
		{args: []string{"encode", "-a", "PS256", "-s", "@./test/rsa_pkcs1_private.pem"}},
		{args: []string{"encode", "-a", "PS384", "-s", "@./test/rsa_pkcs1_private.pem"}},
		{args: []string{"encode", "-a", "PS512", "-s", "@./test/rsa_pkcs1_private.pem"}},
		// Encode ECDSA
		{args: []string{"encode", "-a", "ES256", "-s", "@./test/ecdsa_ec_private.pem"}},
		{args: []string{"encode", "-a", "ES256", "-s", "@./test/ecdsa_ec_private.der"}},
		{args: []string{"encode", "-a", "ES256", "-s", "@./test/ecdsa_pkcs8_private.pem"}},
		{args: []string{"encode", "-a", "ES256", "-s", "@./test/ecdsa_pkcs8_private.der"}},
		{args: []string{"encode", "-a", "ES384", "-s", "@./test/ecdsa_ec_private.pem"}},
		{args: []string{"encode", "-a", "ES512", "-s", "@./test/ecdsa_ec_private.pem"}},
		// Decode
		{args: []string{"decode", "@./test/hmac_hs256.jwt"}},
		// Validate HMAC
		{args: []string{"validate", "@./test/hmac_hs256.jwt", "-a", "HS256", "-s", "secret"}},
		{args: []string{"validate", "@./test/hmac_hs384.jwt", "-a", "HS384", "-s", "secret"}},
		{args: []string{"validate", "@./test/hmac_hs512.jwt", "-a", "HS512", "-s", "secret"}},
		// Validate RSA
		{args: []string{"validate", "@./test/rsa_rs256.jwt", "-a", "RS256", "-s", "@./test/rsa_pkcs1_public.pem"}},
		{args: []string{"validate", "@./test/rsa_rs256.jwt", "-a", "RS256", "-s", "@./test/rsa_pkcs1_public.der"}},
		{args: []string{"validate", "@./test/rsa_rs256.jwt", "-a", "RS256", "-s", "@./test/rsa_x509_public.pem"}},
		{args: []string{"validate", "@./test/rsa_rs256.jwt", "-a", "RS256", "-s", "@./test/rsa_x509_public.der"}},
		{args: []string{"validate", "@./test/rsa_rs384.jwt", "-a", "RS384", "-s", "@./test/rsa_pkcs1_public.pem"}},
		{args: []string{"validate", "@./test/rsa_rs512.jwt", "-a", "RS512", "-s", "@./test/rsa_pkcs1_public.pem"}},
		{args: []string{"validate", "@./test/rsa_ps256.jwt", "-a", "PS256", "-s", "@./test/rsa_pkcs1_public.pem"}},
		{args: []string{"validate", "@./test/rsa_ps384.jwt", "-a", "PS384", "-s", "@./test/rsa_pkcs1_public.pem"}},
		{args: []string{"validate", "@./test/rsa_ps512.jwt", "-a", "PS512", "-s", "@./test/rsa_pkcs1_public.pem"}},
		// Validate ECDSA
		{args: []string{"validate", "@./test/ecdsa_es256.jwt", "-a", "ES256", "-s", "@./test/ecdsa_x509_public.pem"}},
		{args: []string{"validate", "@./test/ecdsa_es256.jwt", "-a", "ES256", "-s", "@./test/ecdsa_x509_public.der"}},
		{args: []string{"validate", "@./test/ecdsa_es384.jwt", "-a", "ES384", "-s", "@./test/ecdsa_x509_public.pem"}},
		{args: []string{"validate", "@./test/ecdsa_es512.jwt", "-a", "ES512", "-s", "@./test/ecdsa_x509_public.pem"}},
	}

	for _, tt := range tests {
		cmd := exec.Command("go", append([]string{"run", "jwt.go"}, tt.args...)...)
		output, err := cmd.Output()

		// Should have some output and not errored
		assert.Nil(t, err)
		assert.NotEmpty(t, output)
	}

}
