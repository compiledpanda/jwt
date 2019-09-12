package internal

import (
	"encoding/pem"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDePem(t *testing.T) {
	bytes := dePem(pem.EncodeToMemory(&pem.Block{
		Type:  "MESSAGE",
		Bytes: []byte("test"),
	}))

	// Should extract bytes
	assert.Equal(t, bytes, []byte("test"))

	bytes = dePem(pem.EncodeToMemory(&pem.Block{
		Type:  "MESSAGE",
		Bytes: nil,
	}))

	// Should return empty when passed nil
	assert.Equal(t, bytes, []byte{})

	// Should not try to decode and pass bytes back
	assert.Equal(t, []byte("test"), []byte("test"))
}
