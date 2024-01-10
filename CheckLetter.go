package hangman

import "unicode/utf8"

func CheckLetter(data *Data, choice string) (string, bool) {
	var i int
	var status bool = false

	numChars := utf8.RuneCountInString(choice)
	println(numChars)
	if numChars > 1 {
		if choice == data.Word {
			data.HiddenWord = data.Word
			status = true
			return data.HiddenWord, status
		} else {
			return data.HiddenWord, status
		}
	}
	for i < len(data.Word) {
		if string(data.Word[i]) == choice {
			data.HiddenWord = data.HiddenWord[:i] + choice + data.HiddenWord[i+1:]
			status = true
		}
		i++
	}
	return data.HiddenWord, status
}
