package aefire

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
)

func MD5Base64(message []byte) (string, error) {
	hasher := md5.New()
	_, err := hasher.Write(message)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(hasher.Sum(nil)), nil
}

func HMacSha1(key, message []byte) (string, error) {
	h := hmac.New(sha1.New, key)
	if _, err := h.Write(message); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
