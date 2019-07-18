package main

import "fmt"

func main() {
	ints := []interface{}{1, 7, 2, 3, 45, 46, 47}
	result := remove(func(i interface{}) bool {
		if i.(int) > 20 {
			return true
		}
		return false
	}, ints)
	fmt.Println(result)

	strigs := []interface{}{"ab", "cde", "fg", "hjkl", "a "}
	result = remove(func(i interface{}) bool {
		if len(i.(string)) < 3 {
			return true
		}
		return false
	}, strigs)
	fmt.Println(result)
}

func remove(f func(i interface{}) bool, slice []interface{}) []interface{} {
	var out []interface{}
	for _, e := range slice {
		if f(e) {
			out = append(out, e)
		}
	}
	return out
}
