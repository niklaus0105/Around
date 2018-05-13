// iterate map
package main

import (
"fmt"
)


func main() {

	m := map[string]bool{"key1":true, "key2":false}
	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}

	arr := [5]int{1,2,3,4,5}
	for v := range arr {
		fmt.Println(v)
	}

	sli := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range sli {
		fmt.Printf("2^%d = %d\n", i, v)
	}


	// if conn, err := os.Open("/tmp/file"); err != nil {
	// 	panic(err)
	// }
	// conn.Read()

	// array := []string{}
	// x,y,z := array[0], array[1], array[2]
	// fmt.Printf("%s %s %s\n", x, y, z)

	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


