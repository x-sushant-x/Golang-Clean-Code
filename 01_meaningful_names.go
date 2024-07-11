package main

import (
	"fmt"
	"time"
)

// Constants should be at top of the file always.
const (
	MAX_SLEEP_TIME_IN_SECONDS = 10
)

func main() {
	// Constants should be named in all caps
	const APPLICATION_VERSION = 1.0

	// This single character variable does not specify it's intentions.
	// An extra comment is used to tell its intentions instead.
	d := 10 // elapsed time in days
	fmt.Println(d)

	// Variables should be meaningful and must clear their intentions by their name.
	// Example
	elapsedTimeInDays := 10
	fileAgeInDays := 2
	fmt.Println(elapsedTimeInDays)
	fmt.Println(fileAgeInDays)

	// Another example of non intentful names
	theList := getThem()
	f := 0
	g := 0

	for _, b := range theList {
		if b%2 == 0 {
			f++
			continue
		}
		g++
	}

	// Reading above code for the 1st time will not make any sense due to bad variable names.
	// But if you read it again and deeply we will understand that it is a code to count odd and even number.
	// This will ultimately take more time and more human efforts.

	// Making this code easy to understand
	// Now we have same numbers of variables, functions, operators and statements.
	// But it is way more easy to read and understand this code
	allNumbers := getAllNumbers()
	evenNumbers := 0
	oddNumbers := 0

	for _, number := range allNumbers {
		if number%2 == 0 {
			evenNumbers++
		} else {
			oddNumbers++
		}
	}

	// Avoid Disinformation
	// Disinformation in clean code refers to the use of misleading or unclear naming conventions, comments, or structure.
	// This can confuse readers of the code

	// Example
	// Below code send 2 integer variables to process function. But it's unclear what process function does.
	// We can't find out what process funtion does without reading the internal code.
	x := 10
	y := 20
	process(x, y)

	// Solution: - What we can do is name the function something which gives a clear understanding of it without reading whole code.
	addTwoNumbers(x, y)

	// Make Meaningful Differences
	// Don't rename of modify your variables just to satisfy compiler when you have to refer similar things.

	// Example
	// In below code I don't understand what this `p` means.
	// What is the significance of these numbers at the end?
	// Nothing! they are just added to satisfy compiler.
	p1 := Person{Name: "Sushant"}
	p2 := Person{Name: "John"}
	fmt.Println(p1)
	fmt.Println(p2)

	// Instead we should make meaningful difference like
	sushant := Person{Name: "Sushant"}
	john := Person{Name: "John"}
	fmt.Println(sushant)
	fmt.Println(john)

	// Another Example
	_ = Product{}
	_ = ProductInfo{}
	_ = ProductData{}
	// Above we have 3 different structs. Can you tell me whose object to use to find product price?
	// No? This is meaningless distinction.

	// Use Pronounceable names
	var genymdhms int64
	/* In my terminology above variable name define
	gen -> generation
	y -> year
	m -> month
	d -> day
	h -> hour
	m -> minute
	s -> second
	*/
	// Can we pronounce this? Probably no. So what we can do is
	var generationTimestamp int64
	// This look a way more better. You can also pronounce this in meetings as well.
	fmt.Println(genymdhms)
	fmt.Println(generationTimestamp)

	// Write searchable variables name.
	// Below code is not searchable. What will you search to find this code? 10? No this is shit.
	time.Sleep(10)
	// Instead what we can do it
	time.Sleep(MAX_SLEEP_TIME_IN_SECONDS)

	// Making Classes & Structs clean
	// Classes names should be noun for example: -
	// Customer, User and etc.

	// Method names :- They should be verbs.
	// Example
	steve := Customer{}
	steve.getPaymentHistroy()
	// Now our struct is noun and getPaymentHistory method is verb.
}

func getThem() []int {
	return []int{1, 2, 3, 4, 5, 6, 7}
}

func getAllNumbers() []int {
	return []int{1, 2, 3, 4, 5, 6, 7}
}

func process(x, y int) int {
	return x + y
}

func addTwoNumbers(x, y int) int {
	return x + y
}

type Person struct {
	Name string
}

type Product struct{}
type ProductInfo struct{}
type ProductData struct{}

type Customer struct{}

func (c Customer) getPaymentHistroy() {}
