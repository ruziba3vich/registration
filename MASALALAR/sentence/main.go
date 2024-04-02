package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// Berilgan gapdagi har bir so‘zning birinchi harfini bosh harf bilan yozish dastur ishlab chiqing.
	// Bo'shliqlar yoki tinish belgilari bilan ajratilgan so'zlarni ko'rib chiqing.
	// Bosh harf bilan yozilgan gapni qaytaring.


	sentence := "this is a test sentence, with punctuation! hello world 123"
	fmt.Println("1 -> boshlang'ich holat :", sentence)
	sentence = process(sentence)
	fmt.Println("2 -> Berilgan gapdagi har bir so'zning birinchi harfini bosh harf bilan yozish :", sentence)
	sentence = takeOff(sentence)
	fmt.Println("3 -> Boshqa bo'shliq yoki tinish belgilari olib tashlangan holat :", sentence)
}

func process(s string) string {
	ss := strings.Split(s, " ")
	var sb strings.Builder

	for _, w := range ss {
		if len(w) > 0 {
			new := strings.ToUpper(string(w[0])) + w[1:]
			sb.WriteString(" " + new)
		}
	}
	return sb.String()
}

func takeOff(s string) (result string) {
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			result += string(char)
		}
	}
	return result
}
