package main
 
import (
 
  "fmt"
 
)
 
func vokalAtauKonsonan(char rune) string  {
 
  switch(char) {
 
    case 'a', 'e', 'i', 'o', 'u':
 
       return "Vokal"
 
    default:
 
      return "Konsonan"
 
  }
}
 
func main() {
 
  x := vokalAtauKonsonan('h')
 
  fmt.Println(x)
 
}