package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//const s = "สวัสดี"
	const s = "你好世界！"
	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Println("%#U starts at %d\n", runeValue, idx)
	}
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(runeValue rune) {
	if runeValue == 't' {
		fmt.Println("found tee")
	} else if runeValue == 'ส' {
		fmt.Println("found so sua")
	}
}
