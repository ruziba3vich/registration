package main

import (
	"fmt"
	"strings"
)

type Dictionary map[string]string

func main() {
	// 1. Foydalanuvchilarga so'zni kiritish va lug'atda mavjud bo'lsa, uning ta'rifini ko'rsating.
	// 2. Lug'atti qiymatlari bo'yicha so'z qidirish (masalan: small berilganda barcha key larni qaytarish)

	dict := Dictionary{
		"apple":      "a fruit that grows on trees and is typically red, yellow, or green",
		"banana":     "a long curved fruit which grows in clusters and has soft pulpy flesh and yellow skin when ripe",
		"orange":     "a round juicy citrus fruit with a tough bright reddish-yellow rind",
		"grape":      "a small oval fruit with smooth skin and juicy flesh, typically eaten raw and used for making wine",
		"pineapple":  "a large tropical fruit with a tough orange or green skin, sweet yellow flesh, and a fibrous core",
		"watermelon": "a large melon with a hard, green rind and sweet, juicy, usually red flesh",
		"strawberry": "a sweet soft red fruit with a seed-studded surface",
		"blueberry":  "a small sweet fruit, typically dark blue, growing on low bushes",
		"peach":      "a round stone fruit with juicy yellow flesh and downy pinkish-yellow skin",
		"pear":       "a yellowish- or brownish-green edible fruit that is typically narrower at the stalk and wider toward the base",
		"kiwi":       "a brown, hairy fruit with bright green flesh and black seeds",
		"mango":      "a tropical fruit with juicy flesh, distinctive aroma, and sweet flavor",
		"avocado":    "a pear-shaped fruit with a rough, dark green or black skin, smooth flesh, and large stone",
		"lemon":      "a yellow, oval citrus fruit with thick skin and fragrant, acidic juice",
		"lime":       "a rounded citrus fruit with green skin, juicy flesh, and a sour taste",
		"coconut":    "a large oval brown seed of a tropical palm, consisting of a hard shell lined with edible white flesh and containing a clear liquid",
		"papaya":     "a tropical fruit shaped like a large melon, with sweet orange flesh and many black seeds",
		"melon":      "a large round fruit with sweet juicy flesh and a thick rind, typically eaten sliced",
		"fig":        "a soft pear-shaped fruit with sweet dark flesh and many small seeds, eaten fresh or dried",
		"apricot":    "a small, round, yellow or orange fruit with a slightly sour taste, eaten fresh or dried",
		"plum":       "a round or oval fruit with a smooth skin and a flattish pointed stone, sweet or tart in flavor",
		"cherry":     "a small, round stone fruit that is typically bright or dark red",
	}

	status := 1

	for status == 1 {
		var option int
		fmt.Print("1 -> So'z orqali key oxtarish : ")
		fmt.Print("2 -> Key orqali ta'rif oxtarish : ")
		fmt.Scan(&option)

		if option == 1 {
			var word string
			fmt.Print("Qaysi so'z orqali key qidirmoqchisiz : ")
			fmt.Scan(&word)
			keys, descs, found := GetKeysByWord(word, dict)

			if found {
				fmt.Println("| ------- >  Siz kiritgan so'zni o'z ichiga olgan so'zlar : ")
				for i := range keys {
					fmt.Println("------------>   ", keys[i], ":", descs[i])
				}
			} else {
				fmt.Println("Siz kiritgan so'z bo'yicha hech qanday ma'lumot topilmadi !")
			}
		} else {
			var key string
			fmt.Print("Qaysi keyga ta'rif oxtaryapsiz : ")
			fmt.Scan(&key)
			desc, ok := SearchByKey(key, dict)

			if ok {
				fmt.Println("Siz qidirgan so'zning lug'atdagi ma'nosi :", desc)
			} else {
				fmt.Println("Siz kiritgan so'zning ma'nosi topilmadi . . .")
			}
		}
		fmt.Print("Yana qandaydir amal bajarasizmi -> [1] Ha / [2] Yo'q : ")
		fmt.Scan(&status)
	}
	fmt.Println("Dasturdan foydalanganingiz uchun raxmat !")
}

func SearchByKey(key string, dict Dictionary) (string, bool) {
	desc, ok := dict[key]

	if ok {
		return desc, true
	} else {
		return "", false
	}
}

func GetKeysByWord(word string, dict Dictionary) (keys []string, descs []string, found bool) {
	for key, val := range dict {
		if process(val, word) {
			keys = append(keys, key)
			descs = append(descs, val)
			found = true
		}
	}

	return keys, descs, found
}

func process(sentence string, word string) bool {
	ss := strings.Split(sentence, " ")
	for _, w := range ss {
		if w == word {
			return true
		}
	}
	return false
}
