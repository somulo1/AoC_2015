package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

// hashMD5 takes a string input and returns its MD5 hash in hexadecimal format.
func hashMD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	secretKey := "ckczppom"
	prefix := "000000" // We need hashes starting with at least five zeroes and six zeros.

	num := 1
	for {
		// Create the input string by appending the number to the secret key.
		input := secretKey + strconv.Itoa(num)
		// Compute the MD5 hash of the input string.
		hash := hashMD5(input)
		// Check if the hash starts with the required prefix.
		if len(hash) >= len(prefix) && hash[:len(prefix)] == prefix {
			// Print the smallest number that meets the criteria.
			fmt.Printf("The lowest number is: %d\n", num)
			break
		}
		// Increment the number and try again.
		num++
	}
}
