package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("zhaoyunxing")

	assert.NoError(t, err)

	claims, err := ParseToken(token)

	assert.NoError(t, err)

	assert.Equal(t, claims.Username, "zhaoyunxing")

}
