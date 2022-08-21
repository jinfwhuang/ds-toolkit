package dsn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateWallet(t *testing.T) {
	b, err := GenerateJWK()
	assert.NoError(t, err)
	assert.NotNil(t, b)
}
