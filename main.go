package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// import (
// 	"fmt"
// 	"math"
// 	"strings"
// )

var smomeNmae = "hello word"


// func sayGreetings(n string){
// 	fmt.Printf("Good morning %v\n", n)
// }

// func sayBy(n string){
// 	fmt.Printf("Goodbye %v", n)
// }


// func cycleNames(n []string, f func(string)){
// 	for _,v := range n{
// 		f(v)
// 	}
// }


// //return					//return type
// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }


// func getInitials(n string) (string, string){
// 	s :=  strings.ToUpper(n)
	
// 	names := strings.Split(s, " ")

// 	var initials []string 
// 	for _, v := range names{
// 		initials = append(initials, v[:1])
// 	}

// 	if(len(initials) > 1) {
// 		return initials[0], initials[1]
// 	}

// 	return initials[0], "_"
// }

var score = 99.9

func updateName(name *string){
	*name = "Code Bro"
}
func getInput(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt,": ")
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}
func createBill() bill {
 //reading user input
	reader := bufio.NewReader(os.Stdin)

	
	
	name, _ := getInput("Create a new bill name", reader)

	b := newBill(name)
	fmt.Println("Create the bill -", b.name)
	
	return b;
}

func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose options (a - add item, s - save bill, t - add tip)", reader)
	

	switch opt {
	case "a":
		name, _ := getInput("Item name", reader)
		price, _ := getInput("Item price", reader)
		//converting from string to float
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price my must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)

		fmt.Println("Item add - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Tip", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be number")
			promptOptions(b)
		}

		//update the tip
		b.updateTip(t)

		fmt.Println("Tip udapted - ", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the file - ", b.name)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}

func main(){

	myBill := createBill()
	promptOptions(myBill)
	//struct
	// mybill := newBill("broCode")
	// fmt.Print(mybill)

	// fmt.Println(mybill.format())
	// mybill.addItem("onion-soup", 10.0)
	// mybill.updateTip(1.0)

	// fmt.Println(mybill.format())

	//pointers
	// pointers()
	// //updated name
	// updateName(&name)
	// pointers()
	//map
	// menu := map[string]float64{
	// 	"soup": 3.44,
	// 	"pie": 3.24,
	// 	"salad": 5.55, 
	// }

	// fmt.Println(menu)
	// fmt.Println(menu["pie"])

	// //looping maps
	// for key, value := range menu{
	// 	fmt.Println(key, "-", value)
	// }

	// //int as key type
	// phonebook := map[int]string{
	// 	2342: "Ebenvezer",
	// 	2341: "Frimpong",
	// }

	// fmt.Print(phonebook)

	// //update item in map
	// phonebook[2342] = "NanaYaw"
	
	// fmt.Print(phonebook)



	//package scope
	// sayHello("BroCode")
	// showScore()
	// for _, value := range points{
	// 	fmt.Println(value)
	// }

	//functions
	// name := []string{"you", "me", "peach", "bowser"}
	// sayGreetings("Brocode") 
	// sayBy("Brocode")
	// cycleNames(name, sayGreetings) 
	// area1 := circleArea(15)
	// fmt.Println(area1)
	// fmt.Printf("%0.2f",area1)

	// fn1, sn := getInitials("Bro Code")
	// fmt.Println(sn, fn1)
	//Boolean
	// age := 45

	// if age < 40 {
	// 	fmt.Print("age is less than 40")
	// }else{
	// 	fmt.Print("age is not less than 40")
	// }


	// name := []string{"you", "me", "peach", "bowser"}

	// for index, value := range name{
	// 	if index == 1 {
	// 		fmt.Println("continuing at pos", index)
	// 		continue
	// 	}

	// 	if index > 2 {
	// 		fmt.Println("breaking at postion: ", index)
	// 		break
	// 	}

	// 	fmt.Printf("The index at pos %v is %v\n", index, value)
	// }

	//loop
	//x := 0
	// for x < 5 {
	// 	fmt.Println("value of x: ", x)
	// 	x++
	// }


	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of i: ",i)
	// }
	
	//name := []string{"you", "me", "peach", "bowser"}

	// for i := 0; i < len(name); i++ {
	// 	fmt.Print(name[i], " ")
	// }

	// for index, value := range name {
	// 	fmt.Printf("The value at index %v is %v\n", index, value)
	// }
	//if index is will not be used
	// for _, value := range name {
	// 	fmt.Printf("The value is %v\n", value)
	// 	value = "hello" // this will not change the original value of the slice
	// }
	//sort
	// age := []int{45, 34, 29, 50, 10, 28}
	// sort.Ints(age)
	// index := sort.SearchInts(age,30)
	// fmt.Print(index)

	// name := []string{"you", "me", "peach", "bowser"}

	// sort.Strings(name)
	// greetings := "hello there friends!"
	// fmt.Print(strings.Contains(greetings, "hello"))
	// fmt.Println(strings.ReplaceAll(greetings, "hello","hi"))
	// fmt.Println(strings.ToUpper(greetings))
	// fmt.Println(strings.Index(greetings, "ll"))
	// fmt.Println(strings.Split(greetings, " "))
	//array
	// //var age[3]int = [3]int{20, 25, 30}
	// var ages = [3]int{20, 25, 30}
	

	// name := [3]string{"Ebenezer", "Frimpong", "Your"}
	// fmt.Print(name, len(ages))

	// //slices (use arrays under hood)
	// var scores = []int{100, 50, 60}
	// scores = append(scores, 31)
	// fmt.Print(scores)

	// //slice ranges
	// rangeOne := name[0:1] // 1-2
	// rangeTwo := name[1:] //start from one and go to end
	// rangeThree := name[:2] //
	// fmt.Print(rangeOne)

	//fmt package
	// fmt.Print()
	// fmt.Println()
	
	// fmt.Printf("my age is %v and my name is %v\n", 20, "Ebenezer")
	// fmt.Printf("my age is %v and my name is %q\n", 20, "Ebenezer")
	// fmt.Printf("my age is %T and my name is %T\n", 20, "Ebenezer")
	// fmt.Printf("my age is %0.2f and my name is %T\n", 20.222, "Ebenezer")

	//splintf (save formated strings)
	// var str = fmt.Sprintf("my age is %0.2f and my name is %T\n", 20.222, "Ebenezer")
	// fmt.Println(str)

	//strings
	// var nameOne string = "BroCode"
	// var nameTwo = "Ebene"
	// var nameThree string
	// nameFour := "I am learning golang"

	// // fmt.Print(nameOne)

	// nameThree = nameOne
	// fmt.Print(nameThree, nameTwo, " ", nameFour)

	//int
	// var intOne int = 1
	// var ageTwo = 30
	// ageThree := 40

	// fmt.Print(intOne, ageTwo, ageThree)

	//bit and memory
	// var num int8 = 127
	// var numTwo int8 = -128
	// var numThree uint8 = 255 //255 bits

	// fmt.Print(num)

	//float
// 	var scoredOne float32 = 1.1
// 	var scoredTwo float64 = 232525.4
// 	scoreThree := 2323.3
}

//				Pass By Value
//GroupA(Non-Pointer )	| 		GroupB (Pointer wrapper Values)
// -----------------------------------------
//		Strings			|		Slices
//		Ints		    |		Maps
//		Floats			|		Fuctions
//		Booleans		|
//		Arrays			|
//		*Structs*		|