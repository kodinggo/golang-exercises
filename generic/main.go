package main

import (
	"fmt"
)

type numeric interface {
	int | int64 | int32 | float32 | float64 | byte
}

func main() {
	s := cetak("hello")
	i := cetak(123)
	b := cetak(true)

	fmt.Println(s, i, b)

	struct1 := struct {
		ID int
	}{
		ID: 1,
	}
	struct2 := struct {
		ID int
	}{
		ID: 1,
	}

	fmt.Println(isEqual(1, 1))
	fmt.Println(isEqual("blah", "blah"))
	fmt.Println(isEqual(struct1, struct2))

	slice1 := []string{"senin", "selasa"}
	slice2 := []string{"rabu", "kamis"}
	newSlice := mergeSlice(slice1, slice2)
	fmt.Println(newSlice)

	newSliceFloat := mergeSlice([]float32{1.2, 2.4}, []float32{1.2, 2.4})
	fmt.Println(newSliceFloat)

	m := map[int]string{
		1: "satu",
	}
	fmt.Println(findValueMap(m, 1))
}

func cetak[T string | int | bool](arg T) T {
	return arg
}

func printNumber[T numeric](n T) {
	fmt.Println(n)
}

func isEqual[x comparable](arg1, arg2 x) bool {
	return arg1 == arg2
}

func mergeSlice[T comparable](slice1, slice2 []T) []T {
	return append(slice1, slice2...)
}

func findValueMap[v comparable, k comparable](m map[k]v, idx k) v {
	return m[idx]
}
