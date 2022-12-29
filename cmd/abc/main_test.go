package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbc(t *testing.T) {
	assert.Equal(t, "abc", Abc())
}

func TestBcd(t *testing.T) {
	assert.Equal(t, "bcd", Bcd())
}
