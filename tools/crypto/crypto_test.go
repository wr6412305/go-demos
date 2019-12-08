package main

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	src := "你好，世界！ hello world"
	fmt.Println(src)
	dst := Base64Encode(src)
	fmt.Println(dst)

	srcDecode, err := Base64Decode(dst)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(srcDecode)
}

func TestKeccak256(t *testing.T) {
	dataHexs := "123456"
	hash, err := Keccak256(dataHexs)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("hash:", hash)
}

func TestSha3_256(t *testing.T) {
	dataHexs := "123456"
	hash, err := Sha3_256(dataHexs)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("hash:", hash)
}
