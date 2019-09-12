package main

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	cmd := exec.Command("go", "run", "jwt.go")
	output, err := cmd.Output()

	// Should have some output and not errored
	assert.Nil(t, err)
	assert.NotEmpty(t, output)
}
