package main

import (
	"encoding/hex"
	"fmt"

	"github.com/InWeCrypto/sha3"
)

func InWeCrypto() {
	keccak256 := sha3.NewKeccak256()
	strData := "fbad17237e6618f10528f9af68c72bd500a5c1032036ae9bdce1c3a02687400f4258a8b6ffee60238047a1917708afe38100aa8a03b37bd81203c85a4a8a3a48"
	fmt.Println("data:", strData)
	byteData, _ := hex.DecodeString(strData)

	keccak256.Write(byteData)
	var digest [32]byte
	keccak256.Sum(digest[:0])
	fmt.Println("keccak256:", hex.EncodeToString(digest[:]))

	checksum1 := sha3.Sum256(byteData)
	fmt.Println("sha3-256:", hex.EncodeToString(checksum1[:]))
}
