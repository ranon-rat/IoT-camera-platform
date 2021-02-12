package main

// here i put some extra code for reduce the size of the code or something like that
import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
)

// manage the errors
func errorControl(err error, w http.ResponseWriter, message string, code int) {
	if err != nil {
		log.Println(err)
		http.Error(w, message, code)

	}

}

// encrypt the data
func encryptData(data string) *string {
	sum := sha256.Sum256([]byte(data)) // we encript the data
	v := hex.EncodeToString(sum[:])
	return &v
}
