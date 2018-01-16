package mapReduce

import "fmt"

func length(s string) int {
	return len(s)
}


func sum(a, b int) int {
	return a + b
}

func màp (list []string, fn func(string) int) []int{
	res := make([]int, len(list))
	for i, item := range list {
		res[i] = fn(item)
	}
	return res
}

func reduce (list []int, fn func(int, int) int) (res int) {
	for _, elem := range list {
		res = fn(res, elem)
	}
	return res
}

func main() {
	list := []string{"a", "bcd", "ff", "zxcv"}
	fmt.Println(reduce(màp(list, length), sum))
}


