package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
)

func main() {
    // Open the file. 
	var arraycasis [] string
	var arraylme [] string
	var arraypatsy [] string
	var count1 int64
	var count2 int64
	var count3 int64
	var count int64
	var elem string
	var input string
	fmt.Println("Enter name")
	fmt.Scanln(&input)
	elem = input+" "
	
	
	//------------------------------------------------------------------------------------
    f, _ := os.Open("C:\\Users\\Nitish\\Desktop\\CHESS\\casis\\results.txt")
    // Create a new Scanner for the file.
    scanner := bufio.NewScanner(f)
    // Loop over all lines in the file and print them.
    for scanner.Scan() {
	line := scanner.Text()
	arraycasis = append(arraycasis,line)
	//fmt.Println(line)
    }
	//-----------------------------------------------------------------------------------
	g, _ := os.Open("C:\\Users\\Nitish\\Desktop\\CHESS\\lme\\results.txt")
    scanner2 := bufio.NewScanner(g)
    for scanner2.Scan() {
	line := scanner2.Text()
	arraylme = append(arraylme,line)
	//fmt.Println(line)
    }
	//-------------------------------------------------------------------------------------
	h, _ := os.Open("C:\\Users\\Nitish\\Desktop\\CHESS\\patsy\\results.txt")
    scanner3 := bufio.NewScanner(h)
	for scanner3.Scan() {
	line := scanner3.Text()
	arraypatsy = append(arraypatsy,line)
	//fmt.Println(line)
    }
	//-------------------------------------------------------------------------------------
	
	for _ ,element := range arraycasis{
	if (strings.Contains(element , elem)){
	fmt.Println(element)
	count1 = count1 + 1
	}
	}
	//---------------------------------------------------
	for _ ,element := range arraylme{
	if (strings.Contains(element , elem)){
	fmt.Println(element)
	count2 = count2 + 1
	}
	}
	//----------------------------------------------------
	for _ ,element := range arraypatsy{
	if (strings.Contains(element , elem)){
	fmt.Println(element)
	count3 = count3 + 1
	}
	}
	//-------------------------------------------------
	count = count1 + count2 + count3
fmt.Println(count)
if count == 1 {
	fmt.Println("SILVER")
} else if count == 2 {
fmt.Println("GOLD")
} else if count ==3 {
fmt.Println("PLATINUM")
} else {
fmt.Println("BRONZE")
}
}
