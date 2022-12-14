package main

import (
	"fmt"
	"strconv"
)

func main() {
	hm := NewMap(5)
	hm.Insert("Ajay", "10")
	hm.Insert("Yadav", "11")
	hm.Insert("Kakodia", "12")
	hm.Insert("Arjun", "13")
	hm.Insert("Sandeep", "14")
	hm.Insert("Anjali", "15")

	fmt.Println(hm.Find("Kakodia"))
	fmt.Println(hm.Find("Anjali"))
	fmt.Println(hm.Find("Kaka"))
	hm.Insert("Kakodia", "Changed!!")
	fmt.Println(hm.Find("Kakodia"))
	fmt.Println(hm.Delete("Kakodia"))
	fmt.Println(hm.Find("Kakodia"))

	for i := 0; i <= 20; i++ {
		hm.Insert("key"+strconv.Itoa(i), strconv.Itoa(i))
	}
	// Key Find in Map
	for i := 0; i <= 20; i++ {
		val, _ := hm.Find("key" + strconv.Itoa(i))
		fmt.Println("key"+strconv.Itoa(i)+": ", val)
	}

}
