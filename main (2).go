package main

import "fmt"

func main() {

	var letter string
	
	fmt.Println("Enter letter grade: ")
	fmt.Scanf("%s",&letter)
	
	runes := []rune(letter)
	sub1 := string(runes[0:1])
	sub2 := string(runes[1:2])
	
	if sub2 == "+" {
		add := 0.3
	}else{
	if sub2 == "-" {
		add := -0.3
	}}else{
	num := 0
	}
	
	if sub1 == "A" {
		num := 4
	}else{
	if sub1 == "B" {
		num := 3
	}else{
	if sub1 == "C" {
		num := 2
	}else{
	if sub1 == "D" {
		num := 1
	}else{
	if sub1 == "F" {
		num := 0	
		add := 0
	}
	
	fmt.Println("Number grade: ")
	fmt.Print(num + add)

}
