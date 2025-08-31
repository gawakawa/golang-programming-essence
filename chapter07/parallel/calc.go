package calc

import "time"

func Add(a, b int) int {
	result := a + b

	time.Sleep(3 * time.Second)
	return result
}
