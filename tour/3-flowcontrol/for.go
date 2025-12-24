package flowcontrol

import "fmt"

// LoopAscii de ascii do 33 ao 122
func LoopAscii() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d = %s\n", i, string(rune(i)))
	}
}
