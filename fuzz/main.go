package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
	// misal abcd
	// 0 i=0 j=3 b[0] = d b[3] = a
	// 1 i=1 j=2 b[1] = c b[2] = b
}

func main() {
	// input := "The quick brown fox jumped over the lazy dog"
	input := "ڗ"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)

	fmt.Printf("Original: %q\n", input)
	fmt.Printf("reversed: %q, err:%v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err:%v\n", doubleRev, doubleRevErr)

	t := "ڗ"
	fmt.Println([]byte(t))
	fmt.Println(len(t))
}
