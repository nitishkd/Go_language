package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
)

func main() {
    // Open the file. 
	var array [] string
	var count int64
	var elem string = "Audrey X"
    f, _ := os.Open("C:\\Users\\Nitish\\Desktop\\CHESS\\casis\\results.txt")
    // Create a new Scanner for the file.
    scanner := bufio.NewScanner(f)
    // Loop over all lines in the file and print them.
    for scanner.Scan() {
	line := scanner.Text()
	array = append(array,line)
	//fmt.Println(line)
    }
	
	

	for _ ,element := range array{
	//fmt.Println(element)
	//fmt.Println(strings.Contains(element , elem))
	if (strings.Contains(element , elem)){
	fmt.Println("Found the string")
	fmt.Println(element)
	count = count + 1
	}
	
}	
fmt.Println(count)
}
