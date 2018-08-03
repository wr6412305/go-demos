// hello.go
// 一个包可以由多个源文件组成，只要他们开头的包声明一样
// 一个包对应生成一个*.a文件，生成的文件名并不是包名+.a,而是目录名+.a
// go install xxx这里对应的并不是包名，而是路径名
// import xxx这里使用的也不是包名，也是路径名
// xxx.SayHello()这里使用的才是包名
package hello_a

import "fmt"

func SayHello(){
	fmt.Println("SayHello()-->Hello")
}
