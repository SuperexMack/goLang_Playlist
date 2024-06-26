package main

import (
	"firstProgram/NewFilo"
	"fmt"
)

func main() {
	fmt.Println("Hello world!!")
	New.Mycall("This is my call ,that's why it is named as same")
	var myName string = "Mack walker"
	fmt.Println((myName))
	const pi = 3.14
	fmt.Println(pi);
	// pi = 89
	// fmt.Println(pi) we cannot perform this operation because pi is a constant
    
	mySkills := "MERN Developer"
	fmt.Println(mySkills)

	// so Basically line 18 is one of the good method to declare a variable because most of the time when we
	// are getting data then we don't have idea that what type of data it is so for that we just use this
	// method

	// Capitalizing the data

	Person := "Kamabuta"
	Relation := "Friend";
	fmt.Println(Person)
	fmt.Println(Relation)

	// So the basic thing is that when we need to import a Variable to the othere files then we need to
	// Capitalized the variable if we will use small letter to Declare the first letter of the Variable then
	// that variable cannnot be used in other files


}
