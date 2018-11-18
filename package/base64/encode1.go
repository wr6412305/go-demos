package main

import (
	"encoding/base64"
	"io"
	"log"
	"os"
)

// 流式API多用于在数据传输过程中对数据进行Base64编码或解码
// 比如在读取较大文件时，或者网络传输时

type Buf struct {
	data []byte
	size int
}

func encode1() {
	filename := "test.txt"
	encodeFilename := "test_enc"

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fEnc, err := os.Create(encodeFilename)
	if err != nil {
		log.Fatalln(err)
	}
	defer fEnc.Close()

	w := base64.NewEncoder(base64.StdEncoding, fEnc)
	io.Copy(w, f)
	w.Close()
}
