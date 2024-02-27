package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

type path []byte

// ini ngubah lenght nya
func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

// sedangkan ini ngubah value nya
func (p path) ToUpper() {
	for i, b := range p {
		if 'a' <= b && b <= 'z' {
			p[i] = b + 'A' - 'a'
		}
	}
}

// jadi kalo misalkan mau ubah yang berurusan dengan length, pakai pointer
// kalo itu engga bisa, karena the function is passed a copy of the slice header, not original.
// sedangkan untuk value nya ga pake pointer karena slice nge pointer ke array
// tapi walapun pake pointer juga yang ToUpper masih bisa
func SlicePointer() {
	pathName := path("/usr/bin/tso") // di slice header length nya ada 11
	pathName.TruncateAtFinalSlash()  // sekarang slice header nya ada 8
	fmt.Printf("%s\n", pathName)
	pathName.ToUpper() // ini ubah valuenya
	fmt.Printf("%s\n", pathName)

}

func main() {
	t := ""
	for i := 0; i < 20; i++ {
		if i == 10 {
			fmt.Println("10")
			continue
		}
		t += fmt.Sprintf("%v", i)
	}
	fmt.Println(t)
}

func TestSlice() {

	var buffer [256]byte
	fmt.Println(buffer)
	buffer[100] = 'a'
	buffer[101] = 75
	buffer[149] = 75
	buffer[150] = 75
	buffer[151] = 75
	fmt.Println(buffer)

	var slice []byte = buffer[100:150]
	shortForm := buffer[100:150]
	fmt.Println(slice, len(slice), shortForm)

	slice[2] = 55
	slice = slice[0 : len(slice)-1]
	fmt.Println(slice, len(slice), "Z")

	buffer[100] = 'b'
	fmt.Println(slice)
}

func TestByte() {

	b := []byte("test")
	var bb byte = 't'
	fmt.Println(b, bb)
	for k, v := range b {
		fmt.Println(k, string(v), b[k])
	}
}

func TestRune() {
	r := "ಚa haha"
	rr := []rune{97, 97, 3226}

	fmt.Println(rr, len(rr), string(rr))
	// harus di convert dulu ke rune baru kebaca, kalo ga diconvert malah jadi multiple byte
	rrr := []rune(r)
	for i := 0; i < len(rrr); i++ {
		fmt.Println(rrr[i], string(rrr[i]))
	}
	// pake range ga perlu di convert ke rune terlebih dahulu.
	for i, v := range r {
		fmt.Printf("iterate %v, value %v\n", i, v)
	}

	t := "test ಚ 🇦🇽"
	fmt.Println(t, len(t), len(string(t)), utf8.RuneCountInString(t))
	for i, v := range t {
		fmt.Printf("%v %v %v\n", v, t[i], string(v))
	}
}
