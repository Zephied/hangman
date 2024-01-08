package hangmanClassic

func CheckUsedLetters(data *Data, choice string) bool {
	var i int
	for _, letter := range data.LetterUsed {
		if choice == letter {
			return false
		} else {
			i++
		}
	}
	data.LetterUsed = append(data.LetterUsed, choice)
	return true
}
