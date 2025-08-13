//go:generate stringer -type Fruit fruit.go
package enum

type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)
