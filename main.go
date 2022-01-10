package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// 01 Print Hello World!
	fmt.Println("Hello World!")
	
	// 02 Print my name
	name:="Siannas Chris"
	fmt.Println(name)
	
	// 03 Print my name and age
	fmt.Println("My name is",name)
	age:=63
	fmt.Println("My age is", age)

	//04 Print the numbers for 1 to 10
	for counter:=1;counter<=10;counter++{
		fmt.Print(counter," ")
	}

	fmt.Println()
	
	//05 Raise to power
	fmt.Println(math.Pow(2,10))

	//06 Print the numbers from 1 to 1000
	for i:=1;i<=1000;i++{
		fmt.Println(i)
	}

	//07 Generate a random number from 0 to 10
	// seed the generator
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(9)+1)

	//08 Print the current date
	dt:=time.Now()
	fmt.Println(dt.String())
	fmt.Println(dt.Format("01-02-2006"))
}


