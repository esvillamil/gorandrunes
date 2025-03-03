// Ejercicio de rand runes
package main

import (
"fmt"

"github.com/esvillamil/gorandrunes/randrunes"
)

// main is the entry point of the program.
//
// It prints the result of calling the RandRunes function from the randrunes package
// with a length of 17 and a string of uppercase and lowercase letters.
//
// It also prints the result of calling the RandRunesR function from the randrunes package
// with a length of 32 and a string of "sin (θ)= cos(π/2 -θ)".
//
// There is a commented out line of code that assigns a string of uppercase and lowercase
// letters to the runeInput variable.
//
// No return type.
func main() {
fmt.Println(randrunes.RandRunes(17, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"))
runeInput := [] rune("sin (θ)= cos(π/2 -θ)")
fmt.Println(string(randrunes.RandRunesR(32, runeInput)))
}