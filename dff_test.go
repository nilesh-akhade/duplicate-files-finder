package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New("/tmp", true))
}

func TestProcess_DirNotFound(t *testing.T) {
	dff := New("/unavailable-dir", true)
	_, err := dff.Find()
	assert.Error(t, err)
}

func TestProcess_WrongDir(t *testing.T) {
	dff := New("\n", true)
	_, err := dff.Find()
	assert.Error(t, err)
}

func TestProcess(t *testing.T) {
	dff := New("/tmp", true)
	_, err := dff.Find()
	assert.NotNil(t, err)
}
