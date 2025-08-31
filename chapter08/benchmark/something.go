package something

import "fmt"

func makeSomething(n int) []string {
	r := make([]string, n)
	for i := range n {
		r[i] = fmt.Sprintf("%05d 何か", i)
	}
	return r
}
