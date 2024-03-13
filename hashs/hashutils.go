package hashs

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/bcrypt"
)

func Md5Str(text string) string {
	return Md5Bytes([]byte(text))
}

func Md5Bytes(bs []byte) string {
	hash := md5.Sum(bs)
	return hex.EncodeToString(hash[:])
}
func Md5Reader(r io.Reader) string {
	hash := md5.New()
	io.Copy(hash, r)
	return hex.EncodeToString(hash.Sum(nil))
}

func Sha1Bytes(bs []byte) string {
	hash := sha1.Sum(bs)
	return hex.EncodeToString(hash[:])
}

func Sha1Str(text string) string {
	return Sha1Bytes([]byte(text))
}
func Sha1Reader(r io.Reader) string {
	hash := sha1.New()
	io.Copy(hash, r)
	return hex.EncodeToString(hash.Sum(nil))
}
func Sha2Bytes(bs []byte) string {
	hash := sha256.Sum256(bs)
	return hex.EncodeToString(hash[:])
}

func Sha2Str(text string) string {
	return Sha2Bytes([]byte(text))
}
func Sha2Reader(r io.Reader) string {
	hash := sha256.New()
	io.Copy(hash, r)
	return hex.EncodeToString(hash.Sum(nil))
}

func HMAC(bs, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	return mac.Sum(bs)
}

func Bcrypt(password string, cost int) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func BcryptCompare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
