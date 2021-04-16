package stuff

// here i put some extra code for reduce the size of the code or something like that
import (
	"crypto/sha256"
	"encoding/hex"
)

// encrypt the data
func EncryptData(data string) *string {
	sum := sha256.Sum256([]byte(data)) // we encript the data
	v := hex.EncodeToString(sum[:])
	return &v
}
