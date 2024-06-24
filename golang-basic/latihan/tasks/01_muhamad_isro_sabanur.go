package tasks

func Palindrome(name string) bool {

	for i :=0; i < len(name)/2; i ++ {
		if name[i] != name[len(name) - i - 1] {
			return false
		}
			
	}
	
	return true
	
}