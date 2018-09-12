package nato

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func PhoneticStringMap(s string) string {
	words := map[string]string{
		"a": "alfa", "b": "bravo", "c": "charlie", "d": "delta", "e": "echo",
		"f": "foxtrot", "g": "golf", "h": "hotel", "i": "india", "j": "juliett",
		"k": "kilo", "l": "lima", "m": "mike", "n": "november", "o": "oscar",
		"p": "papa", "q": "quebec", "r": "romeo", "s": "sierra", "t": "tango",
		"u": "uniform", "v": "victor", "w": "whiskey", "x": "x-ray", "y": "yankee",
		"z": "zulu", "0": "zero", "1": "one", "2": "two", "3": "three",
		"4": "four", "5": "five", "6": "six", "7": "seven", "8": "eight", "9": "nine",
	}

	rc := []string{}
	for _, r := range s {
		if w, ok := words[string(r)]; ok == true {
			rc = append(rc, w)
			continue
		}
		rc = append(rc, string(r))
	}

	return strings.Join(rc, " ")
}

func PhoneticRuneMap(s string) string {
	words := map[rune]string{
		'a': "alfa", 'b': "bravo", 'c': "charlie", 'd': "delta", 'e': "echo",
		'f': "foxtrot", 'g': "golf", 'h': "hotel", 'i': "india", 'j': "juliett",
		'k': "kilo", 'l': "lima", 'm': "mike", 'n': "november", 'o': "oscar",
		'p': "papa", 'q': "quebec", 'r': "romeo", 's': "sierra", 't': "tango",
		'u': "uniform", 'v': "victor", 'w': "whiskey", 'x': "x-ray", 'y': "yankee",
		'z': "zulu", '0': "zero", '1': "one", '2': "two", '3': "three",
		'4': "four", '5': "five", '6': "six", '7': "seven", '8': "eight", '9': "nine",
	}

	rc := []string{}
	for _, r := range s {
		if w, ok := words[r]; ok == true {
			rc = append(rc, w)
			continue
		}
		rc = append(rc, string(r))
	}

	return strings.Join(rc, " ")
}

func PhoneticArray(s string) string {
	words := [...]string{
		"alfa", "bravo", "charlie", "delta", "echo",
		"foxtrot", "golf", "hotel", "india", "juliett",
		"kilo", "lima", "mike", "november", "oscar",
		"papa", "quebec", "romeo", "sierra", "tango",
		"uniform", "victor", "whiskey", "x-ray", "yankee",
		"zulu", "zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
	}

	rc := []string{}
	for _, r := range s {
		if c := unicode.ToLower(r); c >= 'a' && c <= 'z' {
			pos := c - 'a'

			rc = append(rc, words[pos])
			continue
		}

		if r >= '0' && r <= '9' {
			pos := rune(len(words)) - ('9' - r) - 1
			rc = append(rc, words[pos])
			continue
		}

		rc = append(rc, string(r))
	}

	return strings.Join(rc, " ")
}

func PhoneticSlice(s string) string {
	words := []string{
		"alfa", "bravo", "charlie", "delta", "echo",
		"foxtrot", "golf", "hotel", "india", "juliett",
		"kilo", "lima", "mike", "november", "oscar",
		"papa", "quebec", "romeo", "sierra", "tango",
		"uniform", "victor", "whiskey", "x-ray", "yankee",
		"zulu", "zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
	}

	rc := []string{}
	for _, r := range s {
		if c := unicode.ToLower(r); c >= 'a' && c <= 'z' {
			pos := c - 'a'

			rc = append(rc, words[pos])
			continue
		}

		if r >= '0' && r <= '9' {
			pos := rune(len(words)) - ('9' - r) - 1
			rc = append(rc, words[pos])
			continue
		}

		rc = append(rc, string(r))
	}

	return strings.Join(rc, " ")
}

func main() {
	for _, s := range os.Args[1:] {
		fmt.Printf("%s\n", PhoneticStringMap(s))
	}
}
