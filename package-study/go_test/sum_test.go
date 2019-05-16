package testing101

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := 15
	actual := Sum(numbers)
	if actual != expected {
		t.Errorf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual)
	}
}

// 使用 testing.T 的 Run 方法，它允许我们传递一个要运行的子测试的名称，
// 以及一个用于测试的函数
func TestSum1(t *testing.T) {
	t.Run("[1,2,3,4,5]", testSumFunc([]int{1, 2, 3, 4, 5}, 15))
	t.Run("[1,2,3,4,-5]", testSumFunc([]int{1, 2, 3, 4, -5}, 5))
}

func testSumFunc(numbers []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := Sum(numbers)
		if actual != expected {
			t.Errorf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual)
		}
	}
}

// Go 使用在 ExampleXxx()函数底部的 “Output 注释” 部分来确定预期的输出是什么
// 然后在运行测试时，它将实际输出与注释中的预期输出进行比较，如果不匹配
// 将触发失败的测试
func ExampleSum() {
	numbers := []int{5, 5, 5}
	fmt.Println(Sum(numbers))
	// Output:
	// 15
}
