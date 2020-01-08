package main

func prod(v1, v2 int, c chan int) {
	c <- v1 * v2
}

func product(nums []int) int {
	l := len(nums)
	if l == 0 || l%4 != 0 {
		return 0
	}
	c := make(chan int)
	var a, b int = 1, 1
	for i := 0; i < l; i += 4 {
		go prod(nums[i], nums[i+1], c)
		go prod(nums[i+2], nums[i+3], c)
		a *= <-c
		b *= <-c
	}
	return a * b
}

func singleProduct(nums []int) int {
	l := len(nums)
	if l == 0 || l%4 != 0 {
		return 0
	}

	var ans int = 1

	for _, v := range nums {
		ans *= v
	}
	return ans
}

func main() {}
