package hashcash

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"math"
)

const (
	defaultTokenSize = 16
	defaultNonceSize = 8
	maxTargetBits    = 24
)

func newToken(targetBits uint64) []byte {
	buf := make([]byte, defaultTokenSize)
	target := uint64(1) << (64 - targetBits)

	binary.BigEndian.PutUint64(buf[:8], target)
	_, _ = rand.Read(buf[8:])

	return buf
}

func hash(data, nonce []byte) []byte {
	h := sha256.New()
	h.Write(data)
	h.Write(nonce)
	return h.Sum(nil)
}

func verify(token, nonce []byte) bool {
	h := hash(token, nonce)
	return bytes.Compare(h, token) < 0
}

func solve(token []byte) []byte {
	nonce := make([]byte, defaultNonceSize)

	for i := uint64(0); i < math.MaxUint64; i++ {
		binary.BigEndian.PutUint64(nonce, i)
		if verify(token, nonce) {
			return nonce
		}
	}

	return nil
}
