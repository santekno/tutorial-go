package helper

func IsPolyndrome(word string) bool {
	if len(word) > 0 {
		result := true
		limit := len(word)/2
		for i := 0; i < limit; i++ {
			if (word[i] != word[len(word)-i-1]) {
				result = false
				break
			}
		}
		return result
	} else {
		return false
	}
	
}