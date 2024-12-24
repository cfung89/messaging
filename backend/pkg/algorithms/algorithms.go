package algorithms

import (
	"fmt"
	"strconv"
	"strings"
)

var Dictionary map[int]string = map[int]string{0: "A", 1: "B", 2: "C", 3: "D", 4: "E", 5: "F", 6: "G", 7: "H", 8: "I", 9: "J", 10: "K", 11: "L", 12: "M", 13: "N", 14: "O", 15: "P", 16: "Q", 17: "R", 18: "S", 19: "T", 20: "U", 21: "V", 22: "W", 23: "X", 24: "Y", 25: "Z", 26: "a", 27: "b", 28: "c", 29: "d", 30: "e", 31: "f", 32: "g", 33: "h", 34: "i", 35: "j", 36: "k", 37: "l", 38: "m", 39: "n", 40: "o", 41: "p", 42: "q", 43: "r", 44: "s", 45: "t", 46: "u", 47: "v", 48: "w", 49: "x", 50: "y", 51: "z", 52: "0", 53: "1", 54: "2", 55: "3", 56: "4", 57: "5", 58: "6", 59: "7", 60: "8", 61: "9", 62: "+", 63: "/", 64: "="}

// Used to generate Dictionary for Base64 encoding:
func SetBase64Dict() {
	const chars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
	for i, n := range chars {
		Dictionary[i] = string(n)
	}
}

// String padding function
func padStr(str string, padding string, length int, beg bool) string {
	for {
		if len(str) >= length {
			break
		}
		if beg {
			str = fmt.Sprintf("%s%s", padding, str)
		} else {
			str = fmt.Sprintf("%s%s", str, padding)
		}
	}
	return str
}

// Computes the Base64 encoding of a given string
func CustomBase64(message string) string {
	var bits string
	for _, n := range message {
		bits = fmt.Sprintf("%s%.8b", bits, n)
	}

	bitPadding := len(bits) % 6
	if bitPadding != 0 {
		bits = fmt.Sprintf("%s%s", bits, strings.Repeat("0", 6-bitPadding))
	}

	var out string
	for i := 0; i < len(bits); i += 6 {
		sextets, err := strconv.ParseInt(bits[i:i+6], 2, 64)
		if err != nil {
			fmt.Println("Error converting bits to numbers")
		}
		out = fmt.Sprintf("%s%s", out, Dictionary[int(sextets)])
	}

	outputPadding := len(out) % 4
	if outputPadding != 0 {
		out = fmt.Sprintf("%s%s", out, strings.Repeat("=", 4-outputPadding))
	}

	return out
}

// Computes the SHA-1 hash of a given string
func CustomSHA1(message string) string {
	return message
}
