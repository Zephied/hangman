package hangman

type Data struct {
	Word       string
	HiddenWord string
	Life       int
	LetterUsed []string
	Success    bool
	S          string
	Lost       bool
	Win        bool
	Eye        string
}
