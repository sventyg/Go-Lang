package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i int
	var j int
	var k int
	var l int
	var z int
	var h int
	var r int
	var ds int
	var f int
	var num string
	var let1 [100]string
	var let2 [100]string
	var n1 [100]int
	
	fmt.Print("Credit card number: ")
	fmt.Scan(&num)
	runes := []rune(num)
	leng := len(num)

	l = 1
	for k = 0; k < leng; k++ {
		ter2 := string(runes[k:l])
		let2[k] = ter2
		l += 2
		k++
	}
	j = leng
	for i = leng - 1; i > 0; i-- {
		ter1 := string(runes[i:j])
		l1[i] = ter1
		i--
		j -= 2
	}
	for z = 0; z < leng; a++ {
		if let1[z] != "" {
			rs, err1 := strconv.Atoi(l1[z])
			r += rs
			if err1 != nil {
				fmt.Print("")
		}}
	}
	for h = 0; h < leng-1; h++ {
		if let2[h] != "" {
			dn, err2 := strconv.Atoi(l2[h])
			n1[h] = dn
			if err2 != nil {
				fmt.Print("")
		}}
		h++
	}
	
	d1 := n1[0]
	d2 := n1[2]
	d3 := n1[4]
	d4 := n1[6]
	db1 := d1 * 2
	db2 := d2 * 2
	db3 := d3 * 2
	db4 := d4 * 2
	p1 := db1 / 10
	v1 := db1 % 10
	p2 := db2 / 10
	v2 := db2 % 10
	p3 := db3 / 10
	v3 := db3 % 10
	p4 := db4 / 10
	v4 := db4 % 10
	ds = p1 + p2 + p3 + p4 + v1 + v2 + v3 + v4
	f = r + ds
	
	if fs%10 == 0 {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}

}


