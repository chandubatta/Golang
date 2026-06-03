package strings

import (
	"os"
	"strings"
)

// NOTE : Using byte slice instantiate staring, Because strings in the go are sequence of bytes
// ///////////////////////////////////////////////////////
// Creating
// //////////////////////////////////////////////////////
var s string = "hellos"
var s2 string //Empty string

func ReadBytes() string {
	data, err := os.ReadFile("file.txt")
	return string(data)
}

// creating rune
var r1 rune = '@'
var r2 rune //0 value rune

// ///////////////////////////////////////////////////////
// Common Operations
// //////////////////////////////////////////////////////
func length() {
	s := "hello"
	l := len(s) //5  // The leangth function returns number of bytes
	// Not a number of characters and runes.
	s2 := "he`llo" //6 //e upside ` total take 2 bytes`

	runes := []rune(s2)
	l2 := len(runes) //5
}

func Compare() {
	s1 := "hello"
	s2 := "world"

	equal := s1 == s2

	//0 if equal
	//1 if s1 > s2
	//-1 if s1 < s2
	val := strings.Compare(s1, s2)

	equals := strings.EqualFold(s1, s2)

	less := s1 < s2
}

func Concadinating() {
	s1 := "hello"
	s2 := "world"

	s3 := s1 + " " + s2

	b := strings.Builder{}
	b.Grow(1024)
	b.WriteString(s1)
	b.WriteString(" ")
	b.WriteString(s2)
	b.String()
	// If you want to be concadinate the more than 2 strings use Builder
}

// Access a Index in string
func AccessIndex() {
	s := "hello"
	b1 := s[2] //"l"

	s2 := "he`llo"
	b2 := s2[1] //First Half e`

	runes := []rune(s2)
	r:= runes[1]  //e

}
