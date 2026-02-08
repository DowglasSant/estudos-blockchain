package hash

import (
	"crypto/sha256"
	"fmt"
)

func BlockHash(index, nonce int, timestamp, data, previousHash string) string {
	record := fmt.Sprintf("%d%d%s%s%s", index, nonce, timestamp, data, previousHash)
	h := sha256.Sum256([]byte(record))
	return fmt.Sprintf("%x", h)
}
