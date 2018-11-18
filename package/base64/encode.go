package main

import (
	"encoding/base64"
	"fmt"
)

// 为了方便，声明该函数，省去错误处理
func mustDecode(enc *base64.Encoding, str string) (string, error) {
	data, err := enc.DecodeString(str)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func testEncode(enc *base64.Encoding, str string) error {
	encStr := enc.EncodeToString([]byte(str))
	fmt.Println(encStr)

	decStr, err := mustDecode(enc, encStr)
	if err != nil {
		return err
	}

	if decStr != str {
		fmt.Println("UnEqual!")
	}
	return nil
}

func encode() {
	const testStr = "Go语言编程"

	// 测试StdEncoding，注意打印结果里的/为URL中的特殊字符，最后有一个padding
	testEncode(base64.StdEncoding, testStr) // 打印：R2/or63oqIDnvJbnqIs=

	// 测试URLEncoding，可以看到/被替换为_
	testEncode(base64.URLEncoding, testStr) // 打印：R2_or63oqIDnvJbnqIs=

	// 测试RawStdEncoding，可以看到去掉了padding
	testEncode(base64.RawStdEncoding, testStr) // 打印：R2/or63oqIDnvJbnqIs

	// 测试RawURLEncoding，可以看到/被替换Wie_，并且却掉了padding
	testEncode(base64.RawURLEncoding, testStr) // 打印：R2_or63oqIDnvJbnqIs
}
