package main

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

func myripemd160() {
	hash := ripemd160.New()
	dataHex := "01020304"
	data, _ := hex.DecodeString(dataHex)
	hash.Write(data)
	var res [20]byte
	hash.Sum(res[:0])
	fmt.Println(hex.EncodeToString(res[:]))
}
