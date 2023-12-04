package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	secret := strings.TrimSpace(io.ReadAll())
	for i := 1; ; i++ {
		sum := md5.Sum([]byte(fmt.Sprintf("%s%d", secret, i)))
		hash := hex.EncodeToString(sum[:])
		if hash[0] == '0' && hash[1] == '0' && hash[2] == '0' && hash[3] == '0' && hash[4] == '0' {
			fmt.Println(i)
			break
		}
	}
}
