package main

import (
    "bufio"
    "fmt"
    "os"
	"unicode/utf8"
	"unicode"
	"strings"
)

	var arraycasis [] string
	var arraylme [] string
	var arraypatsy [] string
	var namesearch string
	var input string
	var choice int
	var choiceselect int
	
	

func main() {
    // Open the file. 
	
	//------------------------------------------------------------------------------------
    f, _ := os.Open("C:\\Users\\Nitish\\Desktop\\CHESS\\casis\\results.txt")
    // Create a new Scanner for the file.
    scanner := bufio.NewScanner(f)
    // Loop over all lines in the file and print them.
    for scanner.Scan() {
	line := scanner.Text()
	var cc int
	cc = utf8.RuneCountInString(line)
	if cc > 30{
	newline := makename(line)
	arraycasis = append(arraycasis,newline)
	}
	//fmt.Println(line)
    }
	//-----------------------------------------------------------------------------------
	g, _ := os.Open("C:\\Users\\Nitish\\Desktop\\CHESS\\lme\\results.txt")
    scanner2 := bufio.NewScanner(g)
    for scanner2.Scan() {
	line := scanner2.Text()
	var cc int
	cc = utf8.RuneCountInString(line)
	if cc > 30{
	newline := makename(line)
	arraylme = append(arraylme,newline)
	}
	//fmt.Println(line)
    }
	//-------------------------------------------------------------------------------------
	h, _ := os.Open("C:\\Users\\Nitish\\Desktop\\CHESS\\patsy\\results.txt")
    scanner3 := bufio.NewScanner(h)
	for scanner3.Scan() {
	line := scanner3.Text()
	var cc int
	cc = utf8.RuneCountInString(line)
	if cc > 30{
	newline := makename(line)
	arraypatsy = append(arraypatsy,newline)
	//fmt.Println(line)
	}
    }
	//-------------------------------------------------------------------------------------
	fmt.Println("Enter 1 for Searching by Name ")
	fmt.Println("Enter 2 for Searching by Section ")
	fmt.Scanln(&choiceselect)
	
	
	//------------------------------------
	if choiceselect == 1 {
	
	fmt.Println("Enter name") //Enter name for searching
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	input = strings.Title(input)
	namesearch = input + " "
	award2(namesearch)
	
	}
	//-------------------------------------------------------------------------------------
	//---------------------------------------------------------
	if choiceselect == 2 {
	
	
	fmt.Println("Enter 0 for k")
	fmt.Println("Enter 1 for 1st")
	fmt.Println("Enter 2 for 2nd")
	fmt.Println("Enter 3 for 3rd")
	fmt.Println("Enter 4 for 4th")
	fmt.Println("Enter 5 for 5th")
	fmt.Println("Enter 6 for U300")
	fmt.Println("Enter 7 for U600")
	fmt.Println("Enter 8 for monday")
	fmt.Println("Enter 9 for tuesday")
	fmt.Println("Enter 10 for wednesday")
	fmt.Println("Enter 11 for thursday")
	fmt.Println("Enter 12 for friday")
	
	fmt.Scanln(&choice)
	var namelist [] string
	var filelocation string
	
	switch choice{
	default:
		fmt.Println("Invalid case")
	case 0:
		filelocation = "C:\\Users\\Nitish\\Desktop\\CHESS\\sections\\k.txt"
	case 1:
		filelocation = "C:\\Users\\Nitish\\Desktop\\CHESS\\sections\\1st.txt"
	case 2:
		filelocation = "C:\\Users\\Nitish\\Desktop\\CHESS\\sections\\2nd.txt"
	case 3:
		filelocation = "C:\\Users\\Nitish\\Desktop\\CHESS\\sections\\3rd.txt"
	case 4:
		filelocation = "C:\\Users\\Nitish\\Desktop\\CHESS\\sections\\4th.txt"
	case 5:
		filelocation = "C:\\Users\\Nitish\\Desktop\\CHESS\\sections\\5th.txt"
	case 6:
		filelocation = "C:\\Users\\Nitish\\Desktop\\CHESS\\sections\\u300.txt"
	case 7:
		filelocation = "C:\\Users\\Nitish\\Desktop\\CHESS\\sections\\u600.txt"
	case 8:
		filelocation = "C:\\Users\\Nitish\\Desktop\\CHESS\\sections\\mon.txt"
	case 9:
		filelocation = "C:\\Users\\Nitish\\Desktop\\CHESS\\sections\\tue.txt"
	case 10:
		filelocation = "C:\\Users\\Nitish\\Desktop\\CHESS\\sections\\wed.txt"
	case 11:
		filelocation = "C:\\Users\\Nitish\\Desktop\\CHESS\\sections\\thu.txt"
	case 12:
		filelocation = "C:\\Users\\Nitish\\Desktop\\CHESS\\sections\\fri.txt"
	
	}
	i, _ := os.Open(filelocation)
    scanner4 := bufio.NewScanner(i)
	for scanner4.Scan() {
	line := scanner4.Text()
	namelist = append(namelist,line)
	}
	for _ ,stu := range namelist {
		var student string = camel(stu)
		fmt.Println(student + " : ")
		award(student)
	}
	}
	if choiceselect > 2 || choiceselect < 1 {
	fmt.Println("Invalid Choice")
	}
}

func award( elem string ){
	var count1 int64
	var count2 int64
	var count3 int64
	var count int64
	for _ ,element := range arraycasis{
	if (strings.EqualFold(element , elem)){
	fmt.Println(element)
	count1 = count1 + 1
	}
	}
	//---------------------------------------------------
	for _ ,element := range arraylme{
	if (strings.EqualFold(element , elem)){
	fmt.Println(element)
	count2 = count2 + 1
	}
	}
	//----------------------------------------------------
	for _ ,element := range arraypatsy{
	if (strings.EqualFold(element , elem)){
	fmt.Println(element)
	count3 = count3 + 1
	}
	}
	//-------------------------------------------------
	count = count1 + count2 + count3
	fmt.Println(count)
	//fmt.Println(elem + " : ")
	if count == 1 {
		fmt.Println("SILVER")
	} else if count == 2 {
		fmt.Println("GOLD")
	} else if count ==0 {
		fmt.Println("BRONZE")
	} else {
		fmt.Println("PLATINUM")
}
	fmt.Println("-------------------------------------------------")
}

//-------------------------------------------------------------

func camel( w string) string{
var result string
word := []rune(w)
for i := 1 ; i<len(word) ; i++{
	if unicode.IsUpper(word[i]){
	result = w[:i] + w[i:]
	break
	
	}

}	
return result
}

func makename(line string) string {
	l1 := strings.Split(line,"(")
	var t1,t3,t4 string
	t1 = l1[0]
	t2:=strings.Split(t1,", ")
	t3 = t2[1]
	t4 = t2[0]
	t5:= strings.Split(t4," ")
	var namefinal string
	namefinal = t3 + t5[1]
	return namefinal

}

func award2( elem string ){
	var count1 int64
	var count2 int64
	var count3 int64
	var count int64
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
	//fmt.Println(elem + " : ")
	if count == 1 {
		fmt.Println("SILVER")
	} else if count == 2 {
		fmt.Println("GOLD")
	} else if count ==0 {
		fmt.Println("BRONZE")
	} else {
		fmt.Println("PLATINUM")
}
	fmt.Println("-------------------------------------------------")
}
