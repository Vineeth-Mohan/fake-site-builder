package Helper

import (
	"math/rand"
	"os"
	"time"
)

// For Creating Random String
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//CreateRandomString :- Function to create random string
func CreateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//CreateRandomInteger :- Function to find a random int between 0 and max
func CreateRandomInteger(max int) int {
	return rand.Intn(max)
}

//CreateFolderIfNotExist :- Create folder if it does not exist
func CreateFolderIfNotExist(completePath string) error {
	if _, err := os.Stat(completePath); os.IsNotExist(err) {
		error := os.Mkdir(completePath, os.ModePerm)
		return error
	}
	return nil
}
