package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidCep(t *testing.T) {
	address, err := GetLocation("73031031")
	assert.Nil(t, err)
	assert.NotNil(t, address)
}

func TestGetInvalidCep(t *testing.T) {
	address, err := GetLocation("408201100")
	assert.Nil(t, address)
	assert.NotNil(t, err)
}

func TestGetValidWeather(t *testing.T) {
	clima, err := GetWeather("Recife")
	assert.Nil(t, err)
	assert.NotNil(t, clima)
}

func TestGetInvalidWeather(t *testing.T) {
	clima, err := GetWeather("anycity")
	assert.Nil(t, clima)
	assert.NotNil(t, err)
}
