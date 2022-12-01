package main

import "fmt"

func main() {
	hm := NewMap(10)
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
}
