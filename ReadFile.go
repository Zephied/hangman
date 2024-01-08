package github.com/Zephied/hangman

import "os"

func ReadFile(filename string) []byte {
	var file []byte
	var err error
	file, err = os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}
	return file
}
