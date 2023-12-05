package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/crypto"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	secret := strings.TrimSpace(io.ReadAll())
	fmt.Println(crypto.DetectMD5HashWithPrefix(secret, "000000", 1))
}
