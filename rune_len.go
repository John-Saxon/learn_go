package main

import "fmt"

// import "unicode/utf8"
import "strconv"
import "errors"
import "strings"

// golang case

func GetSeparator(s string) rune {
	var sep string
	s = `'` + s + `'`
	sep, _ = strconv.Unquote(s) // please handle the err properly
	fmt.Println("[]rune(sep):", []rune(sep))
	return ([]rune(sep))[0]
}

func main() {
	// s := "\t"
	// fmt.Println(len(s))
	// fmt.Println([]rune(s))
	// fmt.Println(len([]rune(s)))

	// s = `\t`
	// fmt.Println(len(s))
	// fmt.Println([]rune(s))
	// fmt.Println(len([]rune(s)))
	// // fmt.Println(len('\t'))

	// fmt.Println(utf8.DecodeRuneInString(s))
	// fmt.Println(GetSeparator(s))

	// var separator string
	// s = `"\t\t"`
	// _, err := fmt.Sscanf(s, "%q", &separator)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("rune separator", []rune(separator))
	// fmt.Printf("The charater is '%s'.\n", separator)

	// // var separator string
	// s = `"\t,"`
	// _, err = fmt.Sscanf(s, "%q", &separator)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("rune separator", []rune(separator))
	// fmt.Printf("The charater is '%s'.\n", separator)

	s := ` \t `
	r, err := GetSeparator_v3(s)
	if err != nil {
		fmt.Println("Failed to GetSeparator_v1: ", err)
		return
	}
	fmt.Println("del is : ", r)
}

func GetSeparator_v1(s string) (*rune, error) {
	// fmt.Println("s: ", s)
	var sep string
	sq := `'` + s + `'`
	sep, err := strconv.Unquote(sq)
	if err != nil {
		return GetSeparator_v1(s[:len(s)-1])
	}
	rs := []rune(sep)
	if len(rs) > 0 {
		return &rs[0], nil
	}
	return nil, errors.New("No delimiter found!")
}

func GetSeparator_v2(s string) (*rune, error) {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return nil, errors.New("No delimiter found!")
	}
	var sep string
	sq := `'` + s + `'`
	sep, err := strconv.Unquote(sq)
	if err != nil {
		return GetSeparator_v1(s[:len(s)-1])
	}
	rs := []rune(sep)
	if len(rs) > 0 {
		return &rs[0], nil
	}
	return nil, errors.New("No delimiter found!")
}

func GetSeparator_v3(s string) (rune, error) {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return ' ', errors.New("No delimiter found!")
	}
	var sep string
	sq := `'` + s + `'`
	sep, err := strconv.Unquote(sq)
	if err != nil {
		return GetSeparator_v1(s[:len(s)-1])
	}
	rs := []rune(sep)
	if len(rs) > 0 {
		return rs[0], nil
	}
	return '', errors.New("No delimiter found!")
}
