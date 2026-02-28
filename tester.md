package main

import (
	"fmt"
	"net/http"
	"sort"
)

func test() {


	age := 35
	name := "Stik"

    // printer til lokal Terminal
	fmt.Println("Hello world from Go!")

	// print kun en linje
    fmt.Print("Hello Go")
	fmt.Print("Go World /n")

    // print linjer, flertall
    fmt.Println("My age is age", age, "and my name is", name)

	// format spessifier, format string. use: %v (% forran det du vil vise "v" står for var)
	fmt.Printf("my age is %v and my name is %v \n", age, name) // return var
	// Save formatted string
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("saved is:", str)


	// Array & slice
	// var ages [3]int = [3]int{20, 25, 30} // Denne linjen sier: ages skal holde et array med 3 ints. // Spesifiser igjen  "[3]int" og dette blir arrayet inni {} {20, 25, 30}
    // // Shothand
	// var values = [3]int{40, 46, 49} 
    // // Short short hand
	// names := [4]string{"t", "s", "n", "g"}
	// names[1] = "X"

	// fmt.Println(values, len(values))
	// fmt.Println(names, len(names))

	// // Slice (uses array under the hood)
	// var scor = []int{100, 50, 20}
    // scor[2] = 2
	// scor = append(scor, 85)

	// slice ranges
	// rangeOne := names[1:3]
	// rangeTne := names[1:]
	// rangeTRne := names[:3]
	// fmt.Println(scor)

ages := []int{45, 35, 75, 95, 25, 15}

sort.Ints(ages)
fmt.Print(ages)

index := sort.SearchInts(ages, 95)
fmt.Print(index)

names := []string{"Da", "Shi", "Jo"}
sort.Strings(names)
fmt.Print(names)

fmt.Print(sort.SearchStrings(names, "Jo"))
// Start serveren på port 8080 - skal alltid være sist i main()

    http.ListenAndServe(":8080", nil)
}

