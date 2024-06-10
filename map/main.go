package main

import "fmt"

func main() {
	isExtValid := map[string]bool{
		"jpg": true,
		"png": true,
		"gif": true,
	}
	ext := "jpg"
	if isExtValid[ext] {
		//
	}

	animalWords := map[string]string{
		"dog": "anjing",
		"cat": "kucing",
	}
	fmt.Println(animalWords["dog"])
	animalWords["dog"] = "ubah"
	fmt.Println(animalWords["dog"])

	test1 := map[string][]string{
		"hobbies": {"swimming", "reading"},
	}
	fmt.Println(test1)

	test2 := map[string]map[string]int{
		"test": {
			"testlagi": 123,
		},
	}
	fmt.Println(test2)

	fmt.Println(test2["foo"])

	var bar string
	if test2["foo"] == nil {
		bar = "tidak ada value foo"
	}
	fmt.Println(bar)

	c := car{"hallo": "world"}
	fmt.Println(c.get("hallo123", "tidak ada keyword"))
}

type car map[string]string

func (c car) get(k string, defaultValue string) string {
	if c[k] == "" {
		return defaultValue
	}
	return c[k]
}
