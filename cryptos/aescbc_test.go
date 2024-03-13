package cryptos_test

import (
	"testing"

	"github.com/daqiancode/myutils/cryptos"
	"github.com/stretchr/testify/assert"
)

func TestAESCBC(t *testing.T) {
	aseCBC := cryptos.NewAESCBCMd5Key("hello")
	raw := "world"
	encrypted, err := aseCBC.EncryptStr(raw)
	assert.Nil(t, err)
	decrypted, err := aseCBC.DecryptStr(encrypted)
	assert.Nil(t, err)
	assert.Equal(t, decrypted, raw)
}
