package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

// DetectMD5HashWithPrefix detects the first MD5 hash (computed by appending growing numbers to the secret)
// that starts with the given prefix and returns the appended number as well as the hex representation of the hash.
func DetectMD5HashWithPrefix(secret string, prefix string, startIndex int) (int, string) {
	for i := startIndex; ; i++ {
		sum := md5.Sum([]byte(fmt.Sprintf("%s%d", secret, i)))
		hash := hex.EncodeToString(sum[:])
		if strings.HasPrefix(hash, prefix) {
			return i, hash
		}
	}
}
