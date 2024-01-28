package main

import (
	"fmt"
)

type Cipher interface {
	encrypt() interface{}
}

type mystring string
type myarray []int
type mymap map[rune]int

func (s mystring) encrypt() interface{} {
	result := make([]int, len(s))
	for i, char := range s {
		result[i] = int(char) - 'A' + 1
	}
	return result
}

func (arr myarray) encrypt() interface{} {
	result := make([]int, len(arr))
	for i, num := range arr {
		result[i] = collatzfunction(num)
	}
	return result
}

func (m mymap) encrypt() interface{} {
	result := make([]int, len(m))
	i := 0
	for key, value := range m {
		result[i] = int(key) + value
		i++
	}
	return result
}

func collatzfunction(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func main() {
	var userinput string
	fmt.Print("Enter a string, map, or integer array: ")
	fmt.Scan(&userinput)

	switch userinput {
	case "string":
		var userinputstring string
		fmt.Print("Enter a string: ")
		fmt.Scan(&userinputstring)
		str := mystring(userinputstring)
		fmt.Println("Encrypted String:", str.encrypt())
	case "array":
		var arraysize int
		fmt.Print("Enter the size of the integer array: ")
		fmt.Scan(&arraysize)

		arr := make(myarray, arraysize)
		fmt.Print("Enter integers separated by spaces: ")
		for i := 0; i < arraysize; i++ {
			fmt.Scan(&arr[i])
		}
		fmt.Println("Encrypted Integer Array:", arr.encrypt())
	case "map":
		userMap := make(mymap)
		var mapSize int
		fmt.Print("Enter the size of the map: ")
		fmt.Scan(&mapSize)

		fmt.Println("Enter key-value pairs for the map:")
		for i := 0; i < mapSize; i++ {
			var userKey rune
			fmt.Print("Key: ")
			fmt.Scan(&userKey)

			var userValue int
			fmt.Print("Value: ")
			fmt.Scan(&userValue)

			userMap[userKey] = userValue
		}
		fmt.Println("Encrypted Map:", userMap.encrypt())
	default:
		fmt.Println("Invalid input. Please enter 'string', 'array', or 'map'.")
	}
}
