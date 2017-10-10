package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		// true
		strings.Contains("test", "es"),

		// 2
		strings.Count("test", "t"),

		// true
		strings.HasPrefix("test", "te"),

		// true
		strings.HasSuffix("test", "st"),

		// 1
		strings.Index("test", "s"),

		// "a-b"
		strings.Join([]string{"a","b"}, "-"),

		// == "aaaaa"
		strings.Repeat("a", 5),

		// "bbaa"
		strings.Replace("aaaa", "a", "b", 2),

		// []string{a, b, c, d, e, f}
		strings.Split("a,b,c,d,e", ","), 

		// "TEST"
		strings.ToUpper("test"),
	)

	// String to binary data
	arr := []byte("test")
	str := string([]byte{1,2,3,4,5})
	fmt.Println(arr, str)
}