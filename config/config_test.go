package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	c := Load()

	ast := assert.New(t)

	ast.NotNil(c)
	ast.Equal("127.0.0.1", c.PostgresHost)
}
