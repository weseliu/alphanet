package util

import "hash"
import (
	"crypto/md5"
	"encoding/hex"
)

var md5Ctx hash.Hash
func Md5(content string) string {
	if md5Ctx == nil {
		md5Ctx = md5.New()
	}
	cipherStr := md5Ctx.Sum([]byte(content))
	return hex.EncodeToString(cipherStr)
}