package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"math"
	"os"
	"strconv"
)

func generateOTP(length int) (string, error) {

	// Generate 32 random bytes for key and message
	key := make([]byte, 32)
	message := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	if _, err := rand.Read(message); err != nil {
		return "", err
	}

	// Create HMAC-SHA1 hash
	h := hmac.New(sha1.New, key)
	h.Write(message)
	hs := h.Sum(nil)

	// Extract offset from the last nibble of hs
	offset := hs[len(hs)-1] & 0x0F

	// Extract 31 bits starting from the offset
	if int(offset)+4 > len(hs) {
		return "", fmt.Errorf("invalid offset")
	}
	p := ((int(hs[offset]) & 0x7F) << 24) |
		(int(hs[offset+1]) << 16) |
		(int(hs[offset+2]) << 8) |
		int(hs[offset+3])

	// Generate OTP by taking modulus with 10^length
	divisor := int(math.Pow(10, float64(length)))
	hotp := p % divisor

	// Format OTP with leading zeros
	otp := fmt.Sprintf("%0*d", length, hotp)
	return otp, nil
}

func extractLengthFromArgs(args []string) int {
	if len(args) < 2 {
		return 6
	}

	length, err := strconv.Atoi(args[1])

	// Default length is 6 if not specified or out of bounds
	if err != nil || (length < 1 || length > 8) {
		return 6
	}

	return length
}

func main() {
	length := extractLengthFromArgs(os.Args)

	otp, err := generateOTP(length)
	if err != nil {
		fmt.Println("Error generating OTP:", err)
		return
	}
	fmt.Println("Generated OTP:", otp)
}
