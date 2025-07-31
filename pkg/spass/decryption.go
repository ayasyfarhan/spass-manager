package spass

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

const (
	iteration_count = 70000
	key_length      = 256 / 8
	salt_bytes      = 20
)

func Decrypt(data_b64 []byte, password string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(string(data_b64))
	if err != nil {
		return nil, err
	}

	salt := data[:salt_bytes]
	block_size := aes.BlockSize
	iv := data[salt_bytes : salt_bytes+block_size]

	data_enc := data[salt_bytes+block_size:]

	key, err := pbkdf2.Key(sha256.New, password, salt, iteration_count, key_length)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	data_dec := make([]byte, len(data_enc))
	mode.CryptBlocks(data_dec, data_enc)

	data_dec, err = removePKCS7Padding(data_dec)
	if err != nil {
		return nil, err
	}

	return data_dec, nil
}

func removePKCS7Padding(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("empty data")
	}

	paddingLen := int(data[len(data)-1])

	if paddingLen == 0 || paddingLen > len(data) {
		return nil, errors.New("invalid padding length (wrong password?)")
	}

	for i := range paddingLen {
		if data[len(data)-1-i] != byte(paddingLen) {
			return nil, errors.New("invalid padding content (wrong password?)")
		}
	}

	return data[:len(data)-paddingLen], nil
}
