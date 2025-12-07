package main

import (
	"fmt"
)

func SnakeCaseToCamelCase(source string) string {

	result := make([]byte, 0, len(source))

	flag := false

	num := len(source) - 1

	for i := 0; i <= num; i++ {

		d := source[i]
		if flag {
			flag = false
			if d >= 'a' && d <= 'z' {
				d = d - 32
			}
		}

		if i < num {
			if d == '_' {
				flag = true
				continue
			}
		}

		result = append(result, d)
	}

	return string(result[:])
}

func main() {
	result := SnakeCaseToCamelCase("ooo_aa")
	fmt.Println(result)
}
