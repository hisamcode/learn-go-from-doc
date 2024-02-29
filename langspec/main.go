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

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2ffb", b)
}

func main() {
	// i := 4
	// So n << x is "n times 2, x times". And y >> z is "y divided by 2, z times".
	// For example, 1 << 5 is "1 times 2, 5 times" or 32. And 32 >> 5 is "32 divided by 2, 5 times" or 1.
	// 1 << 5 = 1*2, 5x is 1 * 2 = 2, 2 * 2 = 4, 4 * 2 = 8, 8 * 2 = 16, 16 * 2 = 32
	// 3 << 3 = 3 * 2, 3x is 3 * 2 = 6, 6 * 2 = 12, 12 * 2 = 24
	// 5 << 4 = 5 * 2, 4x is 5 * 2 = 10, 10 * 2 = 20, 20 * 2 = 40, 40 * 2 = 80
	fmt.Println(1 << 30)
	fmt.Println(ByteSize(1e13))
	// t := ""
	// for i := 0; i < 20; i++ {
	// 	if i == 10 {
	// 		fmt.Println("10")
	// 		continue
	// 	}
	// 	t += fmt.Sprintf("%v", i)
	// }
	// fmt.Println(t)
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
