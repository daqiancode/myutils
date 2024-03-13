package cryptos_test

import (
	"testing"

	"github.com/daqiancode/myutils/cryptos"
	"github.com/stretchr/testify/assert"
)

func TestRSA(t *testing.T) {
	rsa := cryptos.RSA{}
	pri, pub, err := rsa.GenerateKeyPairPem(1024)
	assert.Nil(t, err)
	priKey, err := rsa.ParsePrivateKeyFromPEM(pri)
	assert.Nil(t, err)
	pubKey, err := rsa.ParsePublicKeyFromPEM(pub)
	assert.Nil(t, err)
	origData := []byte("hello world")
	encryptData, err := rsa.OAEPEncrypt(origData, pubKey)
	assert.Nil(t, err)
	decryptData, err := rsa.OAEPDecrypt(encryptData, priKey)
	assert.Nil(t, err)
	assert.Equal(t, origData, decryptData)
}
