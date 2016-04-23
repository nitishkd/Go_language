// You can edit this code!
// Click here and start typing.
package main

import (
"fmt"
"strings"
)

func main() {

	s := []string {"acd ki dash", "b"}
	s1 := "ki"
	for _ , element := range s{
	var elem string = element
	fmt.Println(strings.ContainsAny(elem, s1))	
	}
	fmt.Println(strings.ContainsAny("team", "e"))
	
}