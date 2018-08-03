package calc

// 函数的第一个字母必须大写，然后才能被别的包使用，
// 若是小写，则别的包不可见
func Add(a int, b int, c chan int){
	sum := a + b
	c <- sum
}