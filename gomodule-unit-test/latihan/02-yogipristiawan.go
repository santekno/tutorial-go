package latihan

import (
	"errors"
	"regexp"
)

func validateInput(str string) (err error) {
	// validate length
	if len(str) > 1 {
		return errors.New("ERROR: input must be 1 character long")
	}

	// validate alphabet
	r, _ := regexp.Compile("^[a-zA-Z]$")
	if !r.MatchString(str) {
		return errors.New("ERROR: input must be an alphabet")
	}

	return nil
}

func checkVowel(str string) (isVowel bool) {
	const vowel string = "aiueoAIUEO"

	for _, val := range vowel {
		if string(val) == str {
			return true
		}
	}

	return
}
