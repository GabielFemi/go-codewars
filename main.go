package main

import "fmt"

func main() {
	stringEndsWith := stringEndsWith("abc", "d")
	fmt.Println(stringEndsWith)
}

func stringEndsWith(s1 string, s2 string) bool {
	hashMap1 := make(map[int]string)
	hashMap2 := make(map[int]string)
	for i := 0; i < len(s1); i++ {
		hashMap1[i] = string(s1[i])
	}

	for i := 0; i < len(s2); i++ {
		hashMap2[i] = string(s2[i])
	}
	if hashMap1[len(s1)-1] == hashMap2[len(s2)-1] {
		return true
	} else {
		return false
	}


}