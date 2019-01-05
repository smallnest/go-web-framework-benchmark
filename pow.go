package main

import (
	"crypto/sha256"
	"math/big"
	"runtime"
	"strconv"
)

func pow(targetBits int) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for {
		data := "hello world " + strconv.Itoa(nonce)
		hash = sha256.Sum256([]byte(data))
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(target) == -1 {
			break
		} else {
			nonce++
		}

		if nonce%100 == 0 {
			runtime.Gosched()
		}
	}
}
