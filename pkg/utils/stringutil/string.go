package stringutil

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"hash/crc32"
	"math/rand"

	"github.com/google/uuid"
	"github.com/spf13/cast"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func MD5(message []byte) string {
	h := md5.New()
	h.Write(message) // 需要加密的字符串为 123456
	return hex.EncodeToString(h.Sum(nil))
}

func CRC32(message []byte) int {
	return cast.ToInt(crc32.ChecksumIEEE(message))
}

func GenRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RemoveDuplicateElement(arr []string) []string {
	result := make([]string, 0, len(arr))
	tmpSet := map[string]struct{}{}
	for _, item := range arr {
		if _, ok := tmpSet[item]; !ok {
			tmpSet[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// GenUUID UUID Version 4
func GenUUID() string {
	return uuid.New().String()
}

func Base64Encode(str2Encode string) string {
	return base64.StdEncoding.EncodeToString([]byte(str2Encode))
}

func Base64Decode(str2Decode string) string {
	bts, _ := base64.StdEncoding.DecodeString(str2Decode)
	return string(bts)
}

func Sha1Encode(str2Encode string) string {
	arr := sha1.Sum([]byte(str2Encode))
	return hex.EncodeToString(arr[:])
}
